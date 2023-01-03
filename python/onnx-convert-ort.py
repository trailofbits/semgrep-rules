import os

# ruleid: onnx-session-options
os.system("python -m onnxruntime.tools.convert_onnx_models_to_ort ./bin/onnx --custom_op_library ./bin/custom_op.dll")