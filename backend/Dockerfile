FROM python:3.11-slim

WORKDIR /app

# Install git
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

# Install torch CPU
RUN pip install torch==2.11.0 --index-url https://download.pytorch.org/whl/test/cpu

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

ENV PYTHONPATH=/app

CMD uvicorn src.main:app --host 0.0.0.0 --port ${PORT:-8000}