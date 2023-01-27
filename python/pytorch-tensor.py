import torch 

# ruleid: pytorch-tensor
y = torch.Tensor(x)

def foo(x): 
  # ruleid: pytorch-tensor
  return torch.Tensor(x)

import torch as t

# ruleid: pytorch-tensor
y = t.Tensor(x)

def foo(x): 
  # ruleid: pytorch-tensor
  return t.Tensor(x)


# ok: pytorch-tensor
y = torch.tensor([0, 1])