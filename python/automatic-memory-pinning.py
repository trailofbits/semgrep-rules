import torch

loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=True)

loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=1)

loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper,
                    pin_memory=3)

# ruleid: automatic-memory-pinning
loader = torch.utils.data.DataLoader(dataset, batch_size=2, collate_fn=collate_wrapper)



