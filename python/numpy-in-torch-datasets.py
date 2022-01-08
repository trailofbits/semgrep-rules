import numpy as np
from torch.utils.data import Dataset, DataLoader

# ruleid: numpy-in-torch-datasets
class RandomDataset(Dataset):
    def __getitem__(self, index):
        return np.random.randint(0, 1000, 3)

    def __len__(self):
        return 1000
      
      
# ruleid: numpy-in-torch-datasets
class AnotherRandomDataset(Dataset):
    def __len__(self):
        return 1000
     
    def __getitem__(self, index):
        print("Hello World")
        x = np.random.randint(0, 1000, 3)
        return x 
      
class YetAnotherRandomDataset(Dataset):
    def __len__(self):
        return 1000
     
    def __getitem__(self, index):
        return index
