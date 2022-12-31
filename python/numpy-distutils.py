from numpy import f2py

sourcecode = ""
f2py.compile(sourcecode, modulename='add')