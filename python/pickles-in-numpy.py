import numpy
import os
import pickle

class Test(object):
    def __init__(self):
        self.a = 1

    def __reduce__(self):
        return (os.system, ("python run.py",))


def test(param_path):
    payload = Test()

    with open(param_path, "wb") as f:
        pickle.dump(payload, f)

    # ruleid: pickles-in-numpy 
    y = numpy.load(param_path, allow_pickle = True)

    # ruleid: pickles-in-numpy 
    x = numpy.load(param_path, allow_pickle = 1)

    # ruleid: pickles-in-numpy 
    z = numpy.load(param_path, allow_pickle=3, encoding='latin1')

    # ruleid: pickles-in-numpy 
    z = numpy.load(allow_pickle='True', file=param_path, encoding='latin1')

    # ruleid: pickles-in-numpy 
    z = numpy.load(allow_pickle='False', file=param_path, encoding='latin1')

    # ok: pickles-in-numpy 
    q = numpy.load(param_path, allow_pickle = False)

    # ok: pickles-in-numpy 
    q = numpy.load(param_path, allow_pickle=None)

    # ok: pickles-in-numpy 
    q = numpy.load(param_path, allow_pickle=[])

    # ok: pickles-in-numpy 
    q = numpy.load(param_path, allow_pickle="")

    # ok: pickles-in-numpy 
    q = numpy.load(param_path, encoding="latin1")

    # ok: pickles-in-numpy
    q = numpy.load(param_path)


def test2():
    payload = Test()

    with open("this_is_a_file.pickle", "wb") as f:
        pickle.dump(payload, f)

    # ok: pickles-in-numpy 
    y = numpy.load("this_is_a_file.pickle", allow_pickle = True)

    # ok: pickles-in-numpy 
    x = numpy.load("this_is_a_file.pickle", allow_pickle = 1)

    # ok: pickles-in-numpy 
    z = numpy.load("this_is_a_file.pickle", allow_pickle=3, encoding='latin1')

    # ok: pickles-in-numpy 
    z = numpy.load(allow_pickle='True', file="completely_safe_for_sure", encoding='latin1')

    # ok: pickles-in-numpy 
    z = numpy.load(allow_pickle='False', file="completely_safe_for_sure", encoding='latin1')
