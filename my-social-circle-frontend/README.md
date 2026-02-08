# My Social Circle Frontend

A sleek and responsive web application built with SvelteKit and Tailwind CSS.

## Features

- **Dashboard**: High-level view of all your circles with quick-create functionality.
- **Circle Details**: View all people in a specific circle, add new members, and rename circles.
- **Person Details**: Comprehensive view of a person's facts and memories, with the ability to edit names and manage (add/delete) specific records.
- **Responsive Design**: Custom-built with Tailwind CSS to switch between "Desktop Mode" (full layouts) and "Focus Mode" (condensed mobile view).
- **Error Resilience**: Robust fetch layer that handles network errors and empty API responses gracefully.

## Tech Stack

- **Framework**: SvelteKit 2.x
- **Language**: TypeScript
- **Styling**: Tailwind CSS
- **State Management**: Svelte's reactive stores and local state.

## Getting Started

1.  **Install Dependencies**:
    ```bash
    npm install
    ```
2.  **Run Locally**:
    ```bash
    npm run dev
    ```
    The app will start at `http://localhost:5173`.

## Architecture

- **`src/lib/api.ts`**: Centralized API client using the native `fetch` API.
- **`src/lib/types.ts`**: Shared TypeScript definitions ensuring type safety between backend and frontend.
- **`src/routes`**: File-based routing for Dashboard, Group Details, and Person snapshots.
