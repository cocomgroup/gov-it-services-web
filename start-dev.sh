#!/bin/bash

echo "🚀 Starting SecureGov Solutions Development Environment"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21+ first."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 18+ first."
    exit 1
fi

echo "✅ Prerequisites check passed"
echo ""

# Start backend
echo "📦 Starting Go backend on port 8080..."
cd backend
go mod download
go run main.go &
BACKEND_PID=$!
cd ..

# Wait for backend to start
sleep 2

# Start frontend
echo "📦 Installing frontend dependencies and starting dev server on port 5173..."
cd frontend
npm install
npm run dev &
FRONTEND_PID=$!
cd ..

echo ""
echo "✅ Development environment started!"
echo ""
echo "🌐 Frontend: http://localhost:5173"
echo "🔌 Backend API: http://localhost:8080"
echo ""
echo "Press Ctrl+C to stop all services"

# Wait for Ctrl+C
trap "kill $BACKEND_PID $FRONTEND_PID; exit" INT
wait