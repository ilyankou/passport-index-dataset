import pandas as pd

input = 'passport-index-iso3-matrix.csv'
output = 'passport-index-iso3-tidy.csv'

df = pd.read_csv(input, index_col=0)
df = df.stack()

df.to_csv(output, header=True)
