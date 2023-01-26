import tensorflow

path = "somelib.so"

# ok: tensorflow-load-library
tensorflow.load_library(path)

# ruleid: tensorflow-load-library
tensorflow.load_library(input())

# ok: tensorflow-load-library
tensorflow.load_op_library(path)

# ruleid: tensorflow-load-library
tensorflow.load_op_library(input())

def test(p):
    # ruleid: tensorflow-load-library
    return tensorflow.load_library(p)

def test2(p):
    # ruleid: tensorflow-load-library
    return tensorflow.load_library(p + ".so")

def test3(p):
    # ruleid: tensorflow-load-library
    return tensorflow.load_op_library(p)

def test4(p):
    # ruleid: tensorflow-load-library
    return tensorflow.load_op_library(p + ".so")
