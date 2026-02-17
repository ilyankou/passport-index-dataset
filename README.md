# Passport Index Dataset

## Travel visa requirements for 199 countries, in .csv and .json

Last updated on 17 February 2026. Check out the `/legacy` folder for earlier versions.

![Passport](passport.png)

### About datasets

There are 7 datasets with identical visa requirements data. One JSON file and six CSV files - three in matrix format and
three in long (tidy) format. Each CSV comes in 3 versions: with country codes as specified in ISO-2 (two-letter codes),
ISO-3 (three-letter codes), and full country names from no particular standard.

* In distance matrices (files with `matrix` in the filename), the first column represents a passport (=from), each
  remaining column represents a destination (=to).
* Files in tidy format (with `tidy` in filename) have three columns: passport (from), destination (to), and the
  requirement.
* The JSON file (`passport-index.json`) contains a nested object keyed by ISO-2 codes:
  `{ "us": { "gb": { "status": "visa free", "days": 180 }, ... } }`.

### Dataset values

For visa-free regimes, the number of days (a positive integer) is specified whenever available in CSV files. When not
available, `visa free` code is used (for example, in the EU countries with the freedom of movement days are not
limited). The JSON file always contains both `status` and `days` (when available) fields.

| Value             | Explanation                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
|-------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `7`-`360`         | Number of visa-free days, where available (CSV only; JSON stores both status and days separately)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `visa free`       | Visa-free travel (where number of days is unknown or not applicable, such as freedom of movement), including tourist registration requirement for Seychelles, e-tickets (Dominican Republic), and (digital) arrival cards (Singapore, Malaysia)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `visa on arrival` | Destinations that grant visa on arrival, basically visa-free                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| `eta`             | Electronic travel authorisation, usually easier to get than e-visas. Includes [ESTA](https://en.wikipedia.org/wiki/Electronic_System_for_Travel_Authorization) (Electronic System for Travel Authorization for the USA), [eTA](https://www.canada.ca/en/immigration-refugees-citizenship/services/visit-canada/eta.html) (Electronic Travel Authorization for Canada), [eVisitors](https://immi.homeaffairs.gov.au/visas/getting-a-visa/visa-listing/evisitor-651) in Australia, [eTourist cards](https://www.surinametourism.sr/tourist-card/) for Suriname, pre-enrollment for Ivory Coast, visa waiver registration for Japan (passport must be registered for visa waiver online or at a Japanese diplomatic mission; the registration is valid for three years or until the passport expires, whichever comes first), and UK's [electronic travel authorisation (ETA)](https://www.gov.uk/guidance/apply-for-an-electronic-travel-authorisation-eta) |
| `e-visa`          | e-Visas                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `visa required`   | Obtaining a visa is required for travel. Includes [Cuba's tourist cards](https://en.wikipedia.org/wiki/Visa_policy_of_Cuba#Tourist_card) and China's [Exit-Entry permits](https://en.wikipedia.org/wiki/Mainland_Travel_Permit_for_Hong_Kong_and_Macao_Resident) for Macau and Hong Kong                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `no admission`    | Includes rare tricky situations, such as war conflicts or [Trump travel bans](https://en.wikipedia.org/wiki/Trump_travel_ban)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `-1`              | where passport=destination                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Legacy values

These values appeared in older versions of the dataset and are no longer present:

| Value                | Explanation                                                       |
|----------------------|-------------------------------------------------------------------|
| `Hayya Entry Permit` | Fan ID for the FIFA World Cup 2022 to enter Qatar in Nov-Dec 2022 |
| `covid ban`          | Travelling is banned for most people due to Covid-19 restrictions |

### Updating the dataset

The parser is a Go application packaged in Docker. You only need Docker and `make` to run it.

```bash
# Full update: build the parser, scrape all 199x198 country pairs, and write new files
make upgrade
```

This will:

1. Build the Docker image with the Go parser
2. Scrape visa requirements from passportindex.org for all country pairs
3. Move existing output files to `legacy/<YYYY-MM-DD>/`
4. Write new CSV and JSON files to the repository root

Previous versions of all datasets are stored in the `legacy/` folder.

### Use CLI instead

Check out this [visa-cli](https://github.com/rand-net/visa-cli) tool by [@rand-net](https://github.com/rand-net) to get
data for specific countries directly from your terminal.

### Source & License

Since "Passport Index is a free tool, built with publicly available information and with content contributed by fans and
government agencies around the world", feel free to use the dataset under the MIT license.

Source: https://www.passportindex.org
