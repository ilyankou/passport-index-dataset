package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	apiURL = "https://www.passportindex.org/core/visachecker.php"
	// apiToken is a some salt or hash, idk. may be changed in the future, but now it works
	apiToken = "1510b5a79a9413083d0342baf7086a5c"
	workers  = 6
	delay    = 50 * time.Millisecond
)

// visaInfo holds the normalized visa requirement for a country pair.
type visaInfo struct {
	Status string `json:"status"`
	Days   int    `json:"days,omitempty"`
}

// --- Status normalization ---------------------------------------------------

// statusSuffixes are stripped from API status values (applied after lowercase).
var statusSuffixes = []string{" (ease)", " (fast track)"}

// statusReplace maps normalized lowercase statuses to unified values.
var statusReplace = map[string]string{
	"exit-entry permit":        "visa required",
	"e-ticket":                 "visa free",
	"visa-free":                "visa free",
	"tourist registration":     "visa free",
	"digital arrival card":     "visa free",
	"arrival card":             "visa free",
	"evisa":                    "e-visa",
	"pre-enrollment":           "eta",
	"evisitors":                "eta",
	"tourist card":             "eta",
	"visa waiver registration": "eta",
	"not admitted":             "no admission",
	"trump ban":                "no admission",
}

// parseText splits "visa-free / 90 days" into status and days,
// normalizing the status to lowercase with unified naming.
func parseText(text string) (status string, days int) {
	parts := strings.SplitN(text, " / ", 2)
	status = strings.TrimSpace(parts[0])
	// "eVisa · visa on arrival" → take only the last part.
	if i := strings.LastIndex(status, " \u00b7 "); i >= 0 {
		status = status[i+len(" \u00b7 "):]
	}
	status = strings.ToLower(status)
	for _, suffix := range statusSuffixes {
		status = strings.TrimSuffix(status, suffix)
	}
	if r, ok := statusReplace[status]; ok {
		status = r
	}
	if len(parts) == 2 {
		daysStr := strings.TrimSuffix(strings.TrimSpace(parts[1]), " days")
		fmt.Sscanf(daysStr, "%d", &days)
	}
	return
}

// --- Sorting ----------------------------------------------------------------

func rowSortKey(iso2 string) string {
	if k, ok := sortKeyBase[iso2]; ok {
		return k
	}
	return iso2ToName[iso2]
}

func colSortKey(iso2 string) string {
	if k, ok := sortKeyCol[iso2]; ok {
		return k
	}
	return rowSortKey(iso2)
}

func sortCountries(keyFunc func(string) string) []string {
	sorted := make([]string, len(countryCodes))
	copy(sorted, countryCodes)
	sort.Slice(sorted, func(i, j int) bool {
		return keyFunc(sorted[i]) < keyFunc(sorted[j])
	})
	return sorted
}

// --- CSV output -------------------------------------------------------------

// visaValue returns the CSV display value for a visaInfo.
// For "visa free" entries returns the number of days; otherwise the status text.
func visaValue(vi visaInfo) string {
	if strings.HasPrefix(vi.Status, "visa free") && vi.Days > 0 {
		return strconv.Itoa(vi.Days)
	}
	return vi.Status
}

func writeMatrixCSV(filename string, data map[string]map[string]visaInfo, codeFunc func(string) string) error {
	rows := sortCountries(rowSortKey)
	cols := append(sortCountries(colSortKey)[1:], sortCountries(colSortKey)[0])

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.UseCRLF = false
	defer w.Flush()

	header := make([]string, len(cols)+1)
	header[0] = "Passport"
	for i, c := range cols {
		header[i+1] = codeFunc(c)
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, from := range rows {
		row := make([]string, len(cols)+1)
		row[0] = codeFunc(from)
		for i, to := range cols {
			if from == to {
				row[i+1] = "-1"
			} else if vi, ok := data[strings.ToUpper(from)][strings.ToUpper(to)]; ok {
				row[i+1] = visaValue(vi)
			}
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeTidyCSV(filename string, data map[string]map[string]visaInfo, codeFunc func(string) string) error {
	rows := sortCountries(rowSortKey)
	cols := append(sortCountries(colSortKey)[1:], sortCountries(colSortKey)[0])

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.UseCRLF = false
	defer w.Flush()

	if err := w.Write([]string{"Passport", "Destination", "Requirement"}); err != nil {
		return err
	}

	for _, from := range rows {
		fromUpper := strings.ToUpper(from)
		for _, to := range cols {
			val := ""
			if from == to {
				val = "-1"
			} else if vi, ok := data[fromUpper][strings.ToUpper(to)]; ok {
				val = visaValue(vi)
			}
			if err := w.Write([]string{codeFunc(from), codeFunc(to), val}); err != nil {
				return err
			}
		}
	}
	return nil
}

// --- API --------------------------------------------------------------------

type visaResponse struct {
	Text string `json:"text"`
	Col  string `json:"col"`
	Link int    `json:"link"`
	Dest string `json:"dest"`
	Pass string `json:"pass"`
}

func checkVisa(client *http.Client, from, to string) (*visaResponse, error) {
	form := url.Values{
		"s":  {from},
		"d":  {to},
		"cl": {apiToken},
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://www.passportindex.org/visa-checker/")
	req.Header.Set("Origin", "https://www.passportindex.org")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var vr visaResponse
	if err := json.Unmarshal(body, &vr); err != nil {
		return nil, fmt.Errorf("parsing JSON %q: %w", string(body), err)
	}

	return &vr, nil
}

// --- Main -------------------------------------------------------------------

type pair struct{ from, to string }

func main() {
	// Test mode: single request.
	if len(os.Args) >= 3 {
		client := &http.Client{Timeout: 10 * time.Second}
		vr, err := checkVisa(client, os.Args[1], os.Args[2])
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		status, days := parseText(vr.Text)
		fmt.Printf("From=%s To=%s Status=%q Days=%d\n", os.Args[1], os.Args[2], status, days)
		return
	}

	// Full scrape mode.
	total := int64(len(countryCodes) * (len(countryCodes) - 1))
	log.Printf("Starting full scrape: %d countries, %d pairs, %d workers",
		len(countryCodes), total, workers)

	bar := progressbar.NewOptions64(total,
		progressbar.OptionSetDescription("Scraping"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionSetItsString("pairs"),
		progressbar.OptionThrottle(200*time.Millisecond),
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "█",
			SaucerHead:    "█",
			SaucerPadding: "░",
			BarStart:      "",
			BarEnd:        "",
		}),
	)

	pairs := make(chan pair, 100)
	go func() {
		for _, from := range countryCodes {
			for _, to := range countryCodes {
				if from != to {
					pairs <- pair{from, to}
				}
			}
		}
		close(pairs)
	}()

	type result struct {
		from, to string
		status   string
		days     int
	}
	results := make(chan result, 100)
	var errCount atomic.Int64

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{Timeout: 15 * time.Second}
			for p := range pairs {
				vr, err := checkVisa(client, p.from, p.to)
				bar.Add(1)
				if err != nil {
					errCount.Add(1)
					bar.Describe(fmt.Sprintf("Scraping (errors: %d)", errCount.Load()))
					if errCount.Load() > 50 {
						log.Printf("Too many errors, worker stopping")
						return
					}
					time.Sleep(2 * time.Second)
					continue
				}
				status, days := parseText(vr.Text)
				results <- result{
					from:   strings.ToUpper(p.from),
					to:     strings.ToUpper(p.to),
					status: status,
					days:   days,
				}
				time.Sleep(delay)
			}
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	data := make(map[string]map[string]visaInfo)
	var written int
	for r := range results {
		if data[r.from] == nil {
			data[r.from] = make(map[string]visaInfo)
		}
		data[r.from][r.to] = visaInfo{Status: r.status, Days: r.days}
		written++
	}
	bar.Finish()
	fmt.Fprintln(os.Stderr)

	// Move existing output files to legacy/<date> before writing new ones.
	outputFiles := []string{
		"passport-index.json",
		"passport-index-matrix.csv",
		"passport-index-matrix-iso2.csv",
		"passport-index-matrix-iso3.csv",
		"passport-index-tidy.csv",
		"passport-index-tidy-iso2.csv",
		"passport-index-tidy-iso3.csv",
	}
	// Check if at least one output file exists.
	hasExisting := false
	for _, f := range outputFiles {
		if _, err := os.Stat(f); err == nil {
			hasExisting = true
			break
		}
	}
	if hasExisting {
		legacyDir := fmt.Sprintf("legacy/%s", time.Now().Format("2006-01-02"))
		if err := os.MkdirAll(legacyDir, 0o755); err != nil {
			log.Fatalf("Creating legacy dir: %v", err)
		}
		for _, f := range outputFiles {
			if _, err := os.Stat(f); err != nil {
				continue // file doesn't exist, skip
			}
			dst := fmt.Sprintf("%s/%s", legacyDir, f)
			if err := os.Rename(f, dst); err != nil {
				log.Fatalf("Moving %s to %s: %v", f, dst, err)
			}
			log.Printf("Moved %s → %s", f, dst)
		}
	}

	// Write JSON.
	jf, err := os.Create("passport-index.json")
	if err != nil {
		log.Fatalf("Creating JSON file: %v", err)
	}
	enc := json.NewEncoder(jf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		jf.Close()
		log.Fatalf("Writing JSON: %v", err)
	}
	jf.Close()
	log.Printf("Done! %d pairs written, %d errors", written, errCount.Load())

	// Write CSVs.
	toName := func(iso2 string) string { return iso2ToName[iso2] }
	toISO2 := func(iso2 string) string { return strings.ToUpper(iso2) }
	toISO3 := func(iso2 string) string { return iso2ToISO3[iso2] }

	csvFiles := []struct {
		file    string
		fn      func(string) string
		writeFn func(string, map[string]map[string]visaInfo, func(string) string) error
	}{
		{"passport-index-matrix.csv", toName, writeMatrixCSV},
		{"passport-index-matrix-iso2.csv", toISO2, writeMatrixCSV},
		{"passport-index-matrix-iso3.csv", toISO3, writeMatrixCSV},
		{"passport-index-tidy.csv", toName, writeTidyCSV},
		{"passport-index-tidy-iso2.csv", toISO2, writeTidyCSV},
		{"passport-index-tidy-iso3.csv", toISO3, writeTidyCSV},
	}
	for _, spec := range csvFiles {
		if err := spec.writeFn(spec.file, data, spec.fn); err != nil {
			log.Fatalf("Writing %s: %v", spec.file, err)
		}
		log.Printf("Written %s", spec.file)
	}
}
