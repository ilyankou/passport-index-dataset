# passport-index-dataset

![Passport](passport.png)

**Passport Index 2019: travel visa requirements for 199 countries, in .csv**

* Last updated on September 28, 2019
* In distance matrices (files containing `-matrix*.csv`), the first column represents a passport (=from), each remaining column represents a foreign country (=to). Files in tidy format (`-tidy*.csv`) have the first column as passport (from), the second column as destination (=to), and the third column as value.
* Values:
  * `3` = visa-free travel
  * `2` = eTA is required
  * `1` = visa can be obtained on arrival (which Passport Index considers visa-free)
  * `0` = visa is required
  * `-1` is for all instances where passport and destination are the same

Data in all 6 *.csv* datasets is the same. Files only differ in the way they present it (distance matrix or tidy), and country codes (ISO-2, ISO-3, and full country names from no particular standard).

Source: https://www.passportindex.org
