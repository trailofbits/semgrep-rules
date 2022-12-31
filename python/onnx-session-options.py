import onnxruntime

path = ""
sess_options = onnxruntime.SessionOptions()
sess_options.register_custom_ops_library(path)
