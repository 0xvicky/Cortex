#!/bin/bash

echo "Stopping app processes..."
pkill -f uvicorn
pkill -f run_worker.py
pkill -f "npm run dev"

echo "Stopping docker containers..."
docker compose down