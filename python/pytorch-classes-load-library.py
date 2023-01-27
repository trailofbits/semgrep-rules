import torch

path = "lib.so"

# ok: pytorch-classes-load-library
torch.classes.load_library(path)

# ruleid: pytorch-classes-load-library
torch.classes.load_library(input())

def test1(p):
    # ruleid: pytorch-classes-load-library
    torch.classes.load_library(input() + p)

def test2(p):
    # ok: pytorch-classes-load-library
    load_library(p)

def test3(p):
    from torch.classes import load_library
    # ruleid: pytorch-classes-load-library
    load_library(p)

def test4(p):
    from torch.classes import load_library
    # ok: pytorch-classes-load-library
    load_library(path)
