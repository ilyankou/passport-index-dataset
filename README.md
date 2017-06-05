# passport-index-dataset
Passport Index 2017: visa requirements for 199 countries, in .csv

* Relevant as of June 1, 2017
* The first column represents a passport (=from), each remaining column represents a visiting country (=to)
* `1` means visa-free travel (this includes visa free, visa on arrival, and eTA)
* `0` means visa is required (implying one needs to apply for a visa beforehand)
* All instances where `row = column` are `1`, as you don't need a visa to visit your own country

Source: https://www.passportindex.org
