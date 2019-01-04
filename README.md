# passport-index-dataset
**Passport Index 2019: travel visa requirements for 199 countries, in .csv**

* Last updated on January 04, 2019
* The first column represents a passport (=from), each remaining column represents a foreign country (=to)
* Values:
  * `3` = visa-free travel
  * `2` = eTA is required
  * `1` = visa can be obtained on arrival (which Passport Index considers visa-free)
  * `0` = visa is required
  * `-1` is for all instances where `row = column`

The 3 *.csv* datasets contain the same data and only differ in how they name the countires: ISO-Alpha-2, ISO-Alpha-3, and full country names from no particular standard.

Source: https://www.passportindex.org
