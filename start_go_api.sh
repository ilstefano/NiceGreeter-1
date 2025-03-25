#!/bin/bash

# Stop any existing Go API server
if [ -f go_api.pid ]; then
  PID=$(cat go_api.pid)
  if ps -p $PID > /dev/null; then
    echo "Stopping existing Go API server (PID: $PID)"
    kill $PID
  fi
  rm go_api.pid
fi

# Compile the Go API server
echo "Compiling Go API server..."
go build -o go_api_server cmd/api/main.go

if [ $? -ne 0 ]; then
  echo "Compilation failed!"
  exit 1
fi

# Start the Go API server in background
echo "Starting Go API server on port 8080..."
nohup ./go_api_server > go_api.log 2>&1 &
echo $! > go_api.pid

echo "Go API server started with PID: $(cat go_api.pid)"
echo "To check logs: cat go_api.log"
echo "To stop server: kill $(cat go_api.pid)"