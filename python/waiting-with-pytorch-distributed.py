import torch.distributed as dist

def bad(): 
  def run(rank, size):
      tensor = torch.zeros(1)
      req = None
      if rank == 0:
          tensor += 1
          # ok: waiting-with-pytorch-distributed
          req = dist.isend(tensor=tensor, dst=1)
          print('Rank 0 started sending')
      else:
          # ok: waiting-with-pytorch-distributed
          req = dist.irecv(tensor=tensor, src=0)
          print('Rank 1 started receiving')
      req.wait()
      print('Rank ', rank, ' has data ', tensor[0])

  # ruleid: waiting-with-pytorch-distributed
  req = dist.isend(tensor=tensor, dst=1)

  # ruleid: waiting-with-pytorch-distributed
  req = dist.irecv(tensor=tensor, src=0)
  return req
