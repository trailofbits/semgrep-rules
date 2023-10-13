from torch import nn, optim
import torch.nn.functional as F
import torch

PATH = "x"

# ok: pickles-in-pytorch
model = torch.load(PATH)

# ok: pickles-in-pytorch
torch.save(model, PATH)

# ok: pickles-in-pytorch
torch.save(model.state_dict(), PATH)

# ok: pickles-in-pytorch
model.load_state_dict(torch.load(PATH))


def test(arg):
    # ruleid: pickles-in-pytorch
    model = torch.load(arg)

    # ruleid: pickles-in-pytorch
    torch.save(model, arg)

    # ok: pickles-in-pytorch
    torch.save(model.state_dict(), arg)

    # ok: pickles-in-pytorch
    model.load_state_dict(torch.load(arg))
    
    state_dict = model.state_dict()
    # ok: pickles-in-pytorch
    torch.save(state_dict, arg)
