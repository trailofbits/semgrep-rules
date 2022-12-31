import torch.nn as nn
import numpy

class MyModule(nn.Module):
    def __init__(self):
        self.dropout = nn.Dropout(0.5)

    def forward(self, x):
        x = self.dropout(x)
        # ruleid: numpy-in-pytorch-modules
        y = numpy.concatenate((x, y), axis=1)
