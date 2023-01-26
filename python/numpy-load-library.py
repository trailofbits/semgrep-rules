import numpy

path = "libname"
directory  = "loader_path/"

# ok: numpy-load-library
numpy.ctypeslib.load_library(path, directory)

# ok: numpy-load-library
numpy.ctypeslib.as_ctypes(2)

# ok: numpy-load-library
numpy.ctypeslib.load_library("lib", "./loader")

# ruleid: numpy-load-library
numpy.ctypeslib.load_library("lib", input())

# ruleid: numpy-load-library
numpy.ctypeslib.load_library(input(), "./loader")

def test(param):
    # ruleid: numpy-load-library
    return numpy.ctypeslib.load_library(param, "./loader")
