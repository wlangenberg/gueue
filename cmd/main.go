package main

import (
    "log"
    "gueue/pkg/queue"
    "gueue/pkg/storage"
    "gueue/api/http"
)


func main() {
    // Init storage
    storage, err := storage.NewStorage("messages.db")
    if err != nil {
        log.Fatalf("Failed to init storage: %v", err)
    }

    // Init queue
    q := queue.NewQueue(storage)

    server := http.NewServer(q)
    log.Println("Starting server on 8080")
    server.Start(":8080")
}
