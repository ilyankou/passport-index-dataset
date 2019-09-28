import pandas as pd

input = 'passport-index-matrix.csv'
output = 'passport-index-tidy.csv'

df = pd.read_csv(input, index_col=0)
df = df.stack()

df.index.names = ['Passport', 'Destination']

df.to_csv(output, index=True, header=['Value'])
