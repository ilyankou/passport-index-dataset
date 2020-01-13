# Passport Index Dataset
## Travel visa requirements for 199 countries, in .csv
Last updated on January 13, 2020

![Passport](passport.png)

### About datasets

All 6 datasets contain the same information, but represent data differently: two and three-letter codes (ISO2, ISO3), and full country names from no particular standard, as well as distance matrix and tidy formats.

In distance matrices (files with `matrix` in the filename), the first column represents a passport (=from), each remaining column represents a destination (=to).

Files in tidy format (with `tidy` in filename) have three columns: passport (from), destination (to), and the code. Note that unlike in the matrix, these files do not have instances were `passport = destination`, and as a result there are no `-1` code values.

### Codes
| Code | Explanation |
|---|---|
|3| visa-free travel|
|2| visa on arrival|
|1| eTA (electronic travel authority) required|
|0| visa required|
|-1| where passport=destination|

### Update data with Jupyter notebook
You should be able to run the Jupyter notebook to update datasets whenever you want. The notebook uses BeautifulSoup to fetch PassportIndex source HTML and extract the `<script>` tag with relevant visa data, then transforms the data using `Js2Py` for Python processing, and then generates 6 datasets.

### Source & License
Since "Passport Index is a free tool, built with publicly available information and with content contributed by fans and government agencies around the world", feel free to use the dataset under the MIT license.

Source: https://www.passportindex.org
