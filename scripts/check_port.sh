#!/bin/bash

PORT=$1

if [ -z "$PORT" ]; then
  echo "❌ Usage: $0 <port>"
  exit 1
fi

echo "🔍 Checking if port $PORT is in use..."

# Detect OS type
OS="$(uname -s)"

if [[ "$OS" == MINGW* || "$OS" == CYGWIN* || "$OS" == MSYS* ]]; then
  # Windows using Git Bash
  PID=$(netstat -ano | grep ":$PORT" | grep LISTENING | awk '{print $5}' | head -n 1)
else
  # Linux/macOS
  PID=$(lsof -ti tcp:"$PORT")
fi

if [ -n "$PID" ]; then
  echo "⚠️  Port $PORT is in use by PID: $PID"
  
  echo "⛔ Killing process $PID automatically..."
  
  if [[ "$OS" == MINGW* || "$OS" == CYGWIN* || "$OS" == MSYS* ]]; then
    # For Windows via Git Bash
    cmd.exe /C "taskkill /PID $PID /F" > /dev/null 2>&1
  else
    # Unix-like
    kill -9 "$PID"
  fi

  echo "✅ Process $PID terminated. Port $PORT is now free."
else
  echo "✅ Port $PORT is free."
fi
