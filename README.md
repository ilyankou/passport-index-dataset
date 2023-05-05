# 2023 Passport Index Dataset
## Travel visa requirements for 199 countries, in .csv
Last updated on 05 May 2023. Check out the `/legacy` folder for earlier versions.

![Passport](passport.png)

### About datasets

There are 6 datasets with identical visa requirements data. Three datasets are matrix and three are long (tidy) format. Each comes in 3 versions: with country codes as specified in ISO-2 (two-letter codes), ISO-3 (three-letter codes), and full country names from no particular standard.

* In distance matrices (files with `matrix` in the filename), the first column represents a passport (=from), each remaining column represents a destination (=to).
* Files in tidy format (with `tidy` in filename) have three columns: passport (from), destination (to), and the requirement.

### Dataset values

For visa-free regimes, the number of days (a positive integer) is specified whenever available. When not available, `visa free` code is used (for example, in the EU countries with the freedom of movement days are not limited).

| Value | Explanation |
|---|---|
|`7`-`360`| Number of visa-free days, where available |
|`visa free`| Visa-free travel (where number of days is unknown or not applicable, such as freedom of movement), including tourist registration requirement for Seychelles|
|`visa on arrival`| Destinations that grant visa on arrival, basically visa-free |
|`e-visa`| Includes  [ESTA](https://esta.cbp.dhs.gov/) (Electronic System for Travel Authorization for the USA) and [eTA](https://www.canada.ca/en/immigration-refugees-citizenship/services/visit-canada/eta.html) (Electronic Travel Authorization for Canada), eVisas, [eVisitors](https://immi.homeaffairs.gov.au/visas/getting-a-visa/visa-listing/evisitor-651) in Australia, [eTourist cards](https://www.surinametourism.sr/tourist-card/) for Suriname, pre-enrollment for Ivory Coast, and UK's [electronic visa waivers](https://www.gov.uk/get-electronic-visa-waiver) |
|`visa required`| Obtaining a visa is required for travel. Includes Cuba's tourist cards |
|`covid ban`| Travelling is banned for most people. This is perhaps the most dynamic category right now, with varying exemptions|
|`no admission`| Includes rare tricky situations, such as war conflicts |
|`Hayya Entry Permit`| Fan ID for the FIFA World Cup 2022 to enter Qatar in Nov-Dec 2022 |
|`-1`| where passport=destination|

### Update data with Jupyter notebook
You should be able to run the Jupyter notebook to update datasets whenever you want.

### Use CLI instead
Check out this [visa-cli](https://github.com/rand-net/visa-cli) tool by [@rand-net](https://github.com/rand-net) to get data for specific countries directly from your terminal.

### Source & License
Since "Passport Index is a free tool, built with publicly available information and with content contributed by fans and government agencies around the world", feel free to use the dataset under the MIT license.

Source: https://www.passportindex.org
