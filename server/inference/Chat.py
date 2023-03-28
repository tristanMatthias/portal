import os
from dataclasses import dataclass
from typing import List

from transformers import AutoModelForCausalLM, AutoTokenizer


@dataclass
class Chat:
  model_name: str

  def __init__(self, model_name: str):
    self.model_name = model_name
    # TODO: Generalize for windows
    self.model_path = os.path.expanduser(f"~/.portal/models/{self.model_name}")
    print(f"Loading model from {self.model_path}")
    self.tokenizer = AutoTokenizer.from_pretrained(self.model_path, local_files_only=True)
    self.model = AutoModelForCausalLM.from_pretrained(self.model_path, local_files_only=True)


  def chat(
    self,
    messages: List[str],
    max_length: int = 50,
    num_return_sequences: int = 1,
    no_repeat_ngram_size: int = 2

  ) -> str:
    prompt = messages[0]
    input_ids = self.tokenizer.encode(prompt, return_tensors="pt")


    # Generate a response using the model
    output = self.model.generate(
      input_ids,
      max_length=max_length,
      num_return_sequences=num_return_sequences,
      no_repeat_ngram_size=no_repeat_ngram_size,

    )

    # Decode the generated response
    return self.tokenizer.decode(output[0], skip_special_tokens=True)
