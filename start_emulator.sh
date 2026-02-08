#!/bin/bash

# Start the Firestore emulator in the background
echo "Starting Firestore Emulator..."
# Use demo project to avoid auth requirement
firebase emulators:start --only firestore --project demo-social-circle &
EMULATOR_PID=$!

# Wait for emulator to start (dumb wait, can be improved with netcat loop)
sleep 5

echo "Firestore Emulator started (PID: $EMULATOR_PID)"
echo "To stop: kill $EMULATOR_PID"

# Export env vars for backend
export FIRESTORE_EMULATOR_HOST="localhost:8080"
export PROJECT_ID="demo-social-circle"

echo "Environment configured:"
echo "FIRESTORE_EMULATOR_HOST=$FIRESTORE_EMULATOR_HOST"
echo "PROJECT_ID=$PROJECT_ID"
