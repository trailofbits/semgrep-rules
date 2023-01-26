from numpy import f2py

sourcecode = "test.exe"

# ok: numpy-f2py-compile
f2py.compile(sourcecode, modulename='add')

# ok: numpy-f2py-compile
f2py.compile("test2.exe", modulename='sub')

# ok: numpy-f2py-compile
f2py.get_include(sourcecode)

# ruleid: numpy-f2py-compile
f2py.compile(input(), modulename='sub')

# ok: numpy-f2py-compile
f2py.compile(sourcecode, modulename=input())

def test(param):
    # ruleid: numpy-f2py-compile
    return f2py.compile(param, modulename='mul')
