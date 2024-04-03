import tensorflow as tf

def id(x):
    return x

model_dir = id("model_dir")

# ok: pickles-in-tensorflow
m1 = tf.saved_model.load("model_dir")
# ruleid: pickles-in-tensorflow
m2 = tf.saved_model.load(model_dir)
