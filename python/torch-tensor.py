import torch 

# ruleid: torch-tensor
y = torch.Tensor(x)

def foo(x): 
  # ruleid: torch-tensor
  return torch.Tensor(x)


import torch as t

# ruleid: torch-tensor
y = t.Tensor(x)

def foo(x): 
  # ruleid: torch-tensor
  return t.Tensor(x)
