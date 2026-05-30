#!/bin/bash

echo "Checking Qdrant container..."

if docker ps -a --format '{{.Names}}' | grep -Eq '^qdrant-cortex$'; then
    echo "Starting existing qdrant-cortex..."
    docker start qdrant-cortex
else
    echo "Creating qdrant-cortex..."
    docker run -d \
        --name qdrant-cortex \
        -p 7000:6333 \
        -p 7001:6334 \
        qdrant/qdrant
fi


echo "Checking Redis container..."

if docker ps -a --format '{{.Names}}' | grep -Eq '^redis-cortex$'; then
    echo "Starting existing redis-cortex..."
    docker start redis-cortex
else
    echo "Creating redis-cortex..."
    docker run -d \
        --name redis-cortex \
        -p 7700:6379 \
        redis
fi


echo "Waiting for services..."
sleep 5

echo "Starting backend..."
cd backend
source venv/bin/activate 
uvicorn src.main:app --reload &

echo "Starting worker..."
python run_worker.py &

echo "Starting frontend..."
cd ../frontend
npm run dev &