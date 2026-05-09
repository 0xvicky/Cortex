FROM python:3.11-slim AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

RUN pip install torch==2.4.0+cpu --index-url https://download.pytorch.org/whl/cpu --no-cache-dir

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Remove ALL GPU packages
RUN pip freeze | grep -iE "nvidia|cuda|triton" | xargs pip uninstall -y 2>/dev/null || true

FROM python:3.11-slim

WORKDIR /app

RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin

COPY . .

ENV PYTHONPATH=/app
CMD uvicorn src.main:app --host 0.0.0.0 --port ${PORT:-8000}