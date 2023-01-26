import joblib
import skops.io as sio

path = "test.joblib"

# ok: scikit-joblib-load
joblib.load(path)

# ruleid: scikit-joblib-load
joblib.load(input())

def test(param):
    param += ".joblib"

    # ruleid: scikit-joblib-load
    x = joblib.load(param)

    # ok: scikit-joblib-load
    unknown_types = sio.get_untrusted_types(param)
    clf = sio.loads(param, trusted=unknown_types)

    # ok: scikit-joblib-load
    clf = sio.loads(param, trusted=True)

    return x, clf
