# My Social Circle - Run Instructions

## Prerequisites
- Go 1.21+
- Node.js & npm
- Firebase CLI (`npm install -g firebase-tools`)
- Firestore Emulator (part of Firebase CLI)

## Quick Start (with Emulator)

1. **Start the Firestore Emulator**:
   In a separate terminal, run:
   ```bash
   firebase emulators:start --only firestore --project demo-social-circle
   ```
   *Note: The emulator will likely start on port **8080** by default.*

2. **Start the Backend**:
   In another terminal, navigate to the backend directory and set the environment variables:
   ```bash
   cd my-social-circle-backend
   export PROJECT_ID=demo-social-circle
   export FIRESTORE_EMULATOR_HOST=localhost:8080
   go run cmd/server/main.go
   ```
   The backend is configured to run on port **8081** to avoid conflict with the emulator.
   You should see: `Server listening on port 8081` and connection logs for the emulator.

3. **Start the Frontend**:
   In a third terminal:
   ```bash
   cd my-social-circle-frontend
   npm install
   npm run dev
   ```
   Open `http://localhost:5173`.

## Troubleshooting
- **Backend Port**: The backend runs on 8081 (default) to not clash with Firestore (8080).
- **Frontend Errors**: Ensure `src/lib/api.ts` points to `http://localhost:8081`.
- **"os imported and not used"**: This compilation error has been fixed.
