# passport-index-dataset
Passport Index 2017: visa requirements for 199 countries, in .csv

* Relevant as of November 30, 2017
* The first column represents a passport (=from), each remaining column represents a foreign country (=to)
* `1` means visa-free travel (this includes visa free, visa on arrival, and eTA)
* `0` means visa is required
* All instances where `row = column` are `1`, as you don't need a visa to go to your own country

Source: https://www.passportindex.org
