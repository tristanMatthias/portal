# Portal

## Requirements
- Go
- Python

## Installation
```
pip install transformers fastapi uvicorn
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## Local development

To start the python inference server:
```
python inference/server.py
```

(This will eventually be started via Go)

To run the app in development

```
wails dev
```

## List of transformers to try
https://huggingface.co/transformers/v4.4.2/
