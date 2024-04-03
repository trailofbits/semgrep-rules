import msgpack
import msgpack_numpy as m
import numpy as np

x = np.random.rand(5)
# ruleid: msgpack-numpy
x_enc = msgpack.packb(x, default=m.encode)
# ruleid: msgpack-numpy
x_rec = msgpack.unpackb(x_enc, object_hook=m.decode)

# ok: msgpack-numpy
x_enc2 = msgpack.packb(x)
# ok: msgpack-numpy
x_rec2 = msgpack.unpackb(x_enc2)

# ok: msgpack-numpy
x_enc3 = msgpack.load(x)
# ok: msgpack-numpy
x_rec3 = msgpack.loads(x_enc2)

m.patch()

# ruleid: msgpack-numpy
x_enc3 = msgpack.packb(x)
# ruleid: msgpack-numpy
x_rec3 = msgpack.unpackb(x_enc2)

# ruleid: msgpack-numpy
x_enc3 = msgpack.load(x)
# ruleid: msgpack-numpy
x_rec3 = msgpack.loads(x_enc2)
