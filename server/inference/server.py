from typing import List

# Create a fast API
from fastapi import FastAPI
from Chat import Chat
from pydantic import BaseModel

app = FastAPI()

class Message(BaseModel):
    messages: List[str]


# Model cache
model_cache = {}

# Create a chat endpoint that has a "model" parameter, and a messages body (list of strings)
@app.post("/chat/{model}")
def chat(model: str, message: Message):
    if model not in model_cache:
        chat = model_cache[model] = Chat(model)
    else:
        chat = model_cache[model]
    return chat.chat(message.messages)


# Run the server
if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="localhost", port=9997)
