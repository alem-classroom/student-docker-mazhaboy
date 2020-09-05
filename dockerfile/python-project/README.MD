# Simple Python API

### Requirements:
- python3

### Install dependencies:
```bash
pip install -r requirements.txt
```

### Run application on port _3000_:
```bash
uvicorn main:app --host 0.0.0.0 --port 3000
```

### To run on port from env variable:
```bash
uvicorn main:app --host 0.0.0.0 --port $PORT
```
