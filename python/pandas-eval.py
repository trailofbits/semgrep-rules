import pandas as pd

def id(x):
    return x

expr = id("something")
colB = id("B")
subexpr = id("df.age * 2")

df1 = pd.DataFrame({'A': range(1, 6), 'B': range(10, 0, -2)})
# ok: pandas-eval
r11 = df1.eval('A + B')
# ruleid: pandas-eval
r12 = df1.eval(expr)
# ok: pandas-eval
r13 = df1.eval(f"A + B")
# ruleid: pandas-eval
r14 = df1.eval(f"A + {colB}")
# ok: pandas-eval
r15 = df1.eval(f"")

df2 = pd.DataFrame({"animal": ["dog", "pig"], "age": [10, 20]})
# ok: pandas-eval
pd.eval("double_age = df.age * 2", target=df2)
# ruleid: pandas-eval
pd.eval(expr, target=df2)
# ok: pandas-eval
pd.eval(f"double_age = df.age * 2", target=df2)
# ruleid: pandas-eval
pd.eval(f"double_age = {subexpr}", target=df2)

df3 = pd.DataFrame({
    'A': range(1, 6),
    'B': range(10, 0, -2)
})
# ok: pandas-eval
r31 = df3.query('A > B')
# ruleid: pandas-eval
r32 = df3.query(expr)
# ok: pandas-eval
r33 = df3.query(f'A > B')
# ruleid: pandas-eval
r34 = df3.query(f'A > {colB}')

class X:
    def query(self, x):
        pass

# ok: pandas-eval
X().query(expr)