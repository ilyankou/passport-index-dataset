# passport-index-dataset

![Passport Illustration](passport.png)

**Passport Index 2019: travel visa requirements for 199 countries, in .csv**

* Last updated on March 09, 2019
* In distance matrices (files named `*-matrix.csv`), the first column represents a passport (=from), each remaining column represents a foreign country (=to). Files in tidy format (those ending with `*-tidy.csv`) have it similar.
* Values:
  * `3` = visa-free travel
  * `2` = eTA is required
  * `1` = visa can be obtained on arrival (which Passport Index considers visa-free)
  * `0` = visa is required
  * `-1` is for all instances where passport and destination are the same

All 6 *.csv* datasets contain the same data, just differently formatted.

Source: https://www.passportindex.org
