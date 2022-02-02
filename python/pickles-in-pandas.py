import pandas as pd 

original_df = pd.DataFrame({"foo": range(5), "bar": range(5, 10)})  

# ruleid: pickles-in-pandas
pd.to_pickle(original_df, "./dummy.pkl")

# ruleid: pickles-in-pandas
original_df.to_pickle("./dummy.pkl")

# ruleid: pickles-in-pandas
unpickled_df = pd.read_pickle("./dummy.pkl")  

# ruleid: pickles-in-pandas
def invalid_func(another = original_df.to_pickle("./dummy.pkl")):
  return another
