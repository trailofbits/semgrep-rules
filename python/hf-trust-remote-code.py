import datasets
import transformers
from datasets import load_dataset
from transformers import (
    AutoConfig,
    AutoModel,
    AutoModelForCausalLM,
    AutoTokenizer,
    pipeline,
)

CHECKPOINT = "org/model"

# ruleid: hf-trust-remote-code
model = AutoModelForCausalLM.from_pretrained(CHECKPOINT, trust_remote_code=True)

# ruleid: hf-trust-remote-code
tokenizer = AutoTokenizer.from_pretrained(CHECKPOINT, trust_remote_code=True)

# ruleid: hf-trust-remote-code
config = AutoConfig.from_pretrained(CHECKPOINT, trust_remote_code=True)

# ruleid: hf-trust-remote-code
base = AutoModel.from_pretrained("org/custom", revision="main", trust_remote_code=True)

# ruleid: hf-trust-remote-code
pipe = pipeline("text-generation", model=CHECKPOINT, trust_remote_code=True)

# ruleid: hf-trust-remote-code
pipe2 = transformers.pipeline("text-generation", model=CHECKPOINT, trust_remote_code=True)

# ruleid: hf-trust-remote-code
ds = load_dataset("org/dataset", trust_remote_code=True)

# ruleid: hf-trust-remote-code
ds2 = datasets.load_dataset("org/dataset", split="train", trust_remote_code=True)


def configurable_load(trust: bool = True):
    # ruleid: hf-trust-remote-code
    return AutoModelForCausalLM.from_pretrained(CHECKPOINT, trust_remote_code=True)


# ok: hf-trust-remote-code
safe_model = AutoModelForCausalLM.from_pretrained(CHECKPOINT)

# ok: hf-trust-remote-code
safe_model_explicit = AutoModelForCausalLM.from_pretrained(CHECKPOINT, trust_remote_code=False)

# ok: hf-trust-remote-code
safe_tokenizer = AutoTokenizer.from_pretrained(CHECKPOINT, trust_remote_code=False)

# ok: hf-trust-remote-code
safe_config = AutoConfig.from_pretrained(CHECKPOINT)

# ok: hf-trust-remote-code
safe_pipe = pipeline("text-generation", model=CHECKPOINT)

# ok: hf-trust-remote-code
safe_pipe_explicit = transformers.pipeline("text-generation", model=CHECKPOINT, trust_remote_code=False)

# ok: hf-trust-remote-code
safe_ds = load_dataset("org/dataset", trust_remote_code=False)

# ok: hf-trust-remote-code
safe_ds_default = datasets.load_dataset("org/dataset", split="train")
