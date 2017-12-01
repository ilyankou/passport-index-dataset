# passport-index-dataset
Passport Index 2017: travel visa requirements for 199 countries, in .csv

* Relevant as of November 30, 2017
* The first column represents a passport (=from), each remaining column represents a foreign country (=to)
* `3` means visa-free travel
* `2` means eTA is required (which is essentially visa free)
* `1` means visa can be obtained on arrival (which Passport Index considers visa-free)
* `0` means visa is required
* All instances where `row = column` are `-1`

Source: https://www.passportindex.org
