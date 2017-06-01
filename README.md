# passport-index-dataset
Passport Index 2017: visa requirements for 199 countries, in .csv

* Relevant as of June 1, 2017
* Each row represents a passport, each column represents a visiting country
* `1` means visa-free travel (this includes visa free, visa on arrival, and eTA)
* `0` means visa is required (implying one needs to apply for a visa beforehand)
* All rows where `row = column` are `0`, as you don't need a visa to visit your own country

Source: https://www.passportindex.org
