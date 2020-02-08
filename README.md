# 2020 Passport Index Dataset
## Travel visa requirements for 199 countries, in .csv
Last updated on February 07, 2020

![Passport](passport.png)

### About datasets

There are 6 datasets with identical visa-free data. Three datasets are matrix and three are long (tidy) format. Each comes in 3 versions: with country codes as specified in ISO-2 (two-letter codes), ISO-3 (three-letter codes), and full country names from no particular standard.

In distance matrices (files with `matrix` in the filename), the first column represents a passport (=from), each remaining column represents a destination (=to).

Files in tidy format (with `tidy` in filename) have three columns: passport (from), destination (to), and the code. Note that unlike in the matrix, these files do not have instances were `passport = destination`, and as a result there are no `-1` code values.

### Dataset values

For visa-free regimes, the `number of days (a positive integer) is specified whenever available`. When not available, `VF` code is used (for example, in the EU countries with the freedom of movement days are not limited).

| Value | Explanation |
|---|---|
|7-360| number of visa-free days|
|VF| visa-free travel (where number of days is not applicable or known, eg freedom of movement)|
|VOA| visa on arrival|
|ETA| eTA (electronic travel authority) required|
|VR| visa required|
|-1| where passport=destination, in matrix files only|

### Update data with Jupyter notebook
You should be able to run the Jupyter notebook to update datasets whenever you want. The notebook uses BeautifulSoup to fetch PassportIndex source HTML and extract the `<script>` tag with relevant visa data, then transforms the data using `Js2Py` for Python processing, and then generates 6 datasets.

### Source & License
Since "Passport Index is a free tool, built with publicly available information and with content contributed by fans and government agencies around the world", feel free to use the dataset under the MIT license.

Source: https://www.passportindex.org
