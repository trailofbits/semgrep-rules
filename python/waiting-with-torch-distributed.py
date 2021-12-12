import torch.distributed as dist

def bad(): 
  # ruleid: waiting-with-torch-distributed
  def run(rank, size):
      tensor = torch.zeros(1)
      req = None
      if rank == 0:
          tensor += 1
          # Send the tensor to process 1
          req = dist.isend(tensor=tensor, dst=1)
          print('Rank 0 started sending')
      else:
          # Receive tensor from process 0
          req = dist.irecv(tensor=tensor, src=0)
          print('Rank 1 started receiving')
      req.wait()
      print('Rank ', rank, ' has data ', tensor[0])

  req = dist.isend(tensor=tensor, dst=1)
  req = dist.irecv(tensor=tensor, src=0)
  return req
