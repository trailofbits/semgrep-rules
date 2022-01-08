import torch.nn as nn
import numpy as np

class MyModule(nn.Module):
    def __init__(self):
        self.dropout = nn.Dropout(0.5)

    def forward(self, x):
        x = self.dropout(x)
        # ruleid: numpy-in-pytorch-modules
        y = np.concatenate((x, y, z), axis=1)
