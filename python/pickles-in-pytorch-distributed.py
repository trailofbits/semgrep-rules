import torch.distributed as dist

if dist.get_rank() == 0:
    objects = ["f", 1] 
else:
    objects = [None, None]
    
# ruleid: pickles-in-pytorch-distributed
dist.broadcast_object_list(objects, src=0)

# ruleid: pickles-in-pytorch-distributed
dist.all_gather_object(output, gather_objects[dist.get_rank()])

# ruleid: pickles-in-pytorch-distributed
dist.gather_object(
        gather_objects[dist.get_rank()],
        output if dist.get_rank() == 0 else None,
        dst=0
    )

# ruleid: pickles-in-pytorch-distributed
dist.scatter_object_list(output_list, objects, src=0)

# ok: pickles-in-pytorch-distributed
dist.scatter(output_list, objects, src=0)
