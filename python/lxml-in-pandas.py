import pandas as pd
import pandas

touch = 1
touch2 = 2 
touch3 = 3

# ruleid: lxml-in-pandas
pd.read_html(touch)

# ruleid: lxml-in-pandas
pandas.read_html(touch, flavor="lxml")

kwargs = {"io": touch, "flavor":"lxml"}
# ruleid: lxml-in-pandas
pd.read_html(**kwargs)

# ok: lxml-in-pandas
pandas.read_html(touch, flavor="bs4")

# ok: lxml-in-pandas
pd.read_html(touch, flavor="html5lib")

kwargs2 = {"io": touch, "flavor":"html5lib"}
# ok: lxml-in-pandas
pd.read_html(**kwargs2)

def test(**kwargs):
    # ruleid: lxml-in-pandas
    pd.read_html(**kwargs)
