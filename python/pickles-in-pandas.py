import pandas as pd 

touch = 1 

# ruleid: pickles-in-pandas
invalid = pd.read_pickle(touch) 

# ruleid: pickles-in-pandas
also_invalid = pd.read_pickle(touch) 

# ruleid: pickles-in-pandas
yet_another = touch.to_pickle
result = yet_another()

# ruleid: pickles-in-pandas
def fake_function(first=null, second=$touch.to_pickle):
    result = second()
    return result 
