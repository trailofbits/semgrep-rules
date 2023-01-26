import torch

# ok: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=True)

# ok: automatic-memory-pinning
torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=1)

# ok: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=3)

# ruleid: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper)

# ruleid: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, pin_memory=False, batch_size=2, collate_fn=collate_wrapper)

