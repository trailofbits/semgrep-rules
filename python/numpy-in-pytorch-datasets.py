import numpy as np
from torch.utils.data import Dataset
from tob.strangelib import Dataset as DatasetStrange

# ruleid: numpy-in-pytorch-datasets
class RandomDataset(Dataset):
    def __getitem__(self, index):
        return np.random.randint(0, 1000, 3)

    def __len__(self):
        return 1000
      
      
# ruleid: numpy-in-pytorch-datasets
class AnotherRandomDataset(Dataset):
    def __len__(self):
        return 1000
     
    def __getitem__(self, index):
        print("Hello World")
        x = np.random.randint(0, 1000, 3)
        return x 

# ruleid: numpy-in-pytorch-datasets
class AnotherRandomDatasetOther(Dataset):
    def __len__(self):
        return 1000
     
    def __getitem__(self, index):
        print("Hello World")
        x = numpy.random.randint(0, 1000, 3)
        return x

# ok: numpy-in-pytorch-datasets
class NotTorchDataset(DatasetStrange):
    def __len__(self):
        return 1000
     
    def __getitem__(self, index):
        print("Hello World")
        x = numpy.random.randint(0, 1000, 3)
        return x 

# ok: numpy-in-pytorch-datasets
class YetAnotherRandomDataset(Dataset):
    def __len__(self):
        return 1000
     
    def __getitem__(self, index):
        return index
