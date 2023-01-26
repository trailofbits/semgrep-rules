import torch

path = "lib.so"

# ok: torch-classes-load-library
torch.classes.load_library(path)

# ruleid: torch-classes-load-library
torch.classes.load_library(input())

def test1(p):
    # ruleid: torch-classes-load-library
    torch.classes.load_library(input() + p)

def test2(p):
    # ok: torch-classes-load-library
    load_library(p)

def test3(p):
    from torch.classes import load_library
    # ruleid: torch-classes-load-library
    load_library(p)

def test4(p):
    from torch.classes import load_library
    # ok: torch-classes-load-library
    load_library(path)
