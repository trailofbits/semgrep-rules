import onnxruntime

path = "some_library.dll"
sess_options = onnxruntime.SessionOptions()

# ok: onnx-session-options
sess_options.register_custom_ops_library(path)

# ok: onnx-session-options
register_custom_ops_library("not onnx method")

# ruleid: onnx-session-options
sess_options.register_custom_ops_library(input())

def test(arg):
    # ruleid: onnx-session-options
    return sess_options.register_custom_ops_library(arg)