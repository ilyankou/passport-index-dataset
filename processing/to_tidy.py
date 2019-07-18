import pandas as pd

input = '../passport-index-matrix-iso3.csv'
output = '../passport-index-tidy-iso3.csv'

df = pd.read_csv(input, index_col=0)
df = df.stack()

df.index.names = ['Passport', 'Destination']

df.to_csv(output, index=True, header=['Value'])
