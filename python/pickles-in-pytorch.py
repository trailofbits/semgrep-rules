from torch import nn, optim
import torch.nn.functional as F
import torch

PATH = 0 

# ruleid: pickles-in-pytorch
model = torch.load(PATH)

# ruleid: pickles-in-pytorch
torch.save(model, PATH)

torch.save(model.state_dict(), PATH)

model.load_state_dict(torch.load(PATH))
