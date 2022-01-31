import pandas as pd  

touch = 1
touch2 = 2 
touch3 = 3

# ruleid: lxml-in-pandas
pandas.read_html(touch)

# ruleid: lxml-in-pandas
pandas.read_html(touch, flavor="lxml")

kwargs = {"io": touch, "flavor":"lxml"}

# ruleid: lxml-in-pandas
pandas.read_html(**kwargs)

pandas.read_html(touch, flavor="bs4")

pandas.read_html(touch, flavor="html5lib")
