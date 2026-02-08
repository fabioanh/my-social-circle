# My Social Circle

A modern, full-stack application for tracking your "Social Circles" — friends, professional contacts, or interest groups. It helps you remember the first memories and facts about the people you meet.

## Project Structure

- **`/my-social-circle-backend`**: A high-performance Go API server using Google Cloud Firestore.
- **`/my-social-circle-frontend`**: A responsive, premium web interface built with SvelteKit and Tailwind CSS.
- **`/firebase.json`**: Configuration for the local Firestore emulator.
- **`start_emulator.sh`**: Helper script to start the local development environment.

## Tech Stack

- **Backend**: Go (Golang), net/http (standard library), Cloud Firestore.
- **Frontend**: SvelteKit, TypeScript, Tailwind CSS.
- **Infrastructure**: Designed for Google Cloud Run and Firestore.

## Getting Started

1. **Prerequisites**: Ensure you have Go 1.21+, Node.js, and the Firebase CLI installed.
2. **Start the Emulator**: Run `./start_emulator.sh` or `firebase emulators:start --only firestore --project demo-social-circle`.
3. **Run the Backend**:
   ```bash
   cd my-social-circle-backend
   export PROJECT_ID=demo-social-circle
   export FIRESTORE_EMULATOR_HOST=localhost:8080
   go run cmd/server/main.go
   ```
4. **Run the Frontend**:
   ```bash
   cd my-social-circle-frontend
   npm run dev
   ```

## Key Features

- **Circles**: Group people into circles (e.g., "Hiking", "Tech Network").
- **Facts**: Record facts and memories for each person.
- **First-Fact Focus**: Automatically tracks and displays the "First Fact" you recorded for someone to give you context on how you met.
- **Mobile Friendly**: Optimized for small screens with responsive layout components.
- **CRUD Operations**: Support for creating, renaming, and deleting data across circles, people, and facts.
