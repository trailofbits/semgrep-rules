import torch

# ruleid: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=True)

# ruleid: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=1)

# ruleid: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=3)

loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper)

loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=False)

loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=0)

