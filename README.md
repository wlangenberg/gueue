# Gueue: A Simple Message Queue in Go

Gueue is a simple, efficient message queue implementation in Go that uses SQLite for persistent storage. This project is designed to handle high traffic efficiently, making it suitable for use cases where reliable message delivery is critical.

## Features

- **In-memory queue with persistent storage**: Messages are stored in memory for fast access and persisted to SQLite for durability.
- **Concurrency**: Uses goroutines and channels to handle concurrent operations safely.
- **HTTP API**: Exposes endpoints to enqueue and dequeue messages.

## Getting Started

### Prerequisites

- Go 1.22.1 or later
- SQLite3

### Installation

1. **Clone the repository**

   ```sh
   git clone https://github.com/wlangenberg/gueue.git
   cd gueue
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```

### Usage

1. Run the server

   ```sh
   go run cmd/main.go
   ```

2. Enque messages
   ```sh
    curl -X POST -H "Content-Type: application/json" -d '{"message": "My first message"}' http://localhost:8080/enqueue
   
    curl -X POST -H "Content-Type: application/json" -d '{"message": "My second message"}' http://localhost:8080/enqueue
   ```

3. Dequeue messages
   ```sh
    curl http://localhost:8080/dequeue 
   ```
   Should respond with "My first message"  
   
   ```sh
    curl http://localhost:8080/dequeue 
   ```
   Should respond with "My last message"

## Project Structure
```
├── api
│   └── http
│       └── server.go     # HTTP server and handlers
├── cmd
│   └── main.go           # Entry point of the application
│   └── messages.db       # The database file for sqllite
├── docs                  # Documentation (Not implemented yet)
├── internal
│   └── config            # (Not implemented yet)
├── pkg
│   ├── queue
│   │   └── queue.go      # Queue implementation
│   └── storage
│       └── storage.go    # SQLite storage implementation
├── go.mod                # Go module file
└── go.sum                # Go dependencies file
```

## How It Works

### Queue Package

* queue.go: Implements the core message queue logic, including enqueuing and dequeuing messages. A goroutine handles database writes to ensure that only one write operation occurs at a time.

### Storage Package

* storage.go: Manages interactions with SQLite, including saving, deleting, and retrieving messages.

### HTTP API

* server.go: Sets up an HTTP server with endpoints to enqueue and dequeue messages.
   
