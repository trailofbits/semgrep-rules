import torch.nn as nn
import numpy as np

class MyModule(nn.Module):
    def __init__(self):
        self.dropout = nn.Dropout(0.5)

    def forward(self, x, y):
        x = self.dropout(x)
        # ruleid: numpy-in-pytorch-modules
        y = np.concatenate((x, y), axis=1)

        # ruleid: numpy-in-pytorch-modules
        np.ndarray.sort(y)

    def forward_correct(self, x, y):
        x = self.dropout(x)
        # ok: numpy-in-pytorch-modules
        y = nn.cat((x, y), 1)
