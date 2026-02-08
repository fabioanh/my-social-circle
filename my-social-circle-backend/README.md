# My Social Circle Backend

A RESTful API built with Go to manage social circles, people, and memories.

## Architecture

- **`cmd/server`**: Entry point for the application. Initializes the Firestore client and starts the HTTP server.
- **`internal/api`**: Routing and HTTP handlers. Uses the standard library's `ServeMux` with Go 1.22+ pattern matching.
- **`internal/models`**: Data structures used across the application.
- **`internal/service`**: Business logic and data persistence (Firestore).

## API Endpoints

- **Groups**: `GET /groups`, `POST /groups`, `GET /groups/{id}`, `PUT /groups/{id}`.
- **People**: `POST /people`, `GET /groups/{id}/people`, `GET /people/{id}`, `PUT /people/{id}`.
- **Facts**: `POST /people/{id}/facts`, `DELETE /people/{id}/facts/{factId}`.
- **Health**: `GET /health`.

## Database Strategy

The backend uses **Cloud Firestore** in Datastore mode. 
- **Denormalization**: To optimize person lists, the `FirstFact` is denormalized and stored directly on the `Person` document.
- **Subcollections**: Facts are stored as a subcollection under each `Person` document to ensure scalability.
- **Zero-Value Lists**: All list endpoints return empty arrays `[]` instead of `null` for better frontend compatibility.
