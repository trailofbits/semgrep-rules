import pandas as pd 

original_df = pd.DataFrame({"foo": range(5), "bar": range(5, 10)})  

# ok: pickles-in-pandas
pd.to_pickle(original_df, "./dummy.pkl")

# ok: pickles-in-pandas
original_df.to_pickle("./dummy.pkl")

# ok: pickles-in-pandas
unpickled_df = pd.read_pickle("./dummy.pkl")  

# ok: pickles-in-pandas
def invalid_func(another = original_df.to_pickle("./dummy.pkl")):
  return another

# ok: pickles-in-pandas
uncsved_df = pd.read_csv("./dummy.pkl")  


def test(data, file):
  # ruleid: pickles-in-pandas
  pd.to_pickle(data, file)

  # ruleid: pickles-in-pandas
  data.to_pickle(file)

  # ruleid: pickles-in-pandas
  unpickled_df = pd.read_pickle(file)  

  # ruleid: pickles-in-pandas
  def invalid_func(another = original_df.to_pickle(file)):
    return another

  # ok: pickles-in-pandas
  uncsved_df = pd.read_csv(file)  
