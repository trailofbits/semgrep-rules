from tensorflow import keras
from keras.models import load_model

def id(x):
    return x

h5_file_path = id("model.h5")
keras_file_path = id("model.keras")

# ruleid: pickles-in-keras-deprecation
m1 = load_model("model.h5")

# ok: pickles-in-keras-deprecation
m2 = load_model("model.keras")

# ruleid: pickles-in-keras-deprecation
m3 = load_model(h5_file_path)

# ruleid: pickles-in-keras-deprecation
m4 = load_model(keras_file_path)
