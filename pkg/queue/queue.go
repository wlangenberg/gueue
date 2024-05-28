package queue

import (
	"gueue/pkg/storage"
	"sync"
    "errors"
)

// The message queue type
type Queue struct {
    storage     *storage.Storage
    mu          sync.Mutex
    messages    []string
    writeCh     chan string
}

// NewQueue inits a new instance of the Queue type
func NewQueue(storage *storage.Storage) *Queue {
    q := &Queue{
        storage:    storage,
        messages:   make([]string, 0),
        writeCh:    make(chan string, 100), // buffered channel for write operations to database
    }

    go q.runWriter() // Start the writer goroutine
    return q
}

func (q *Queue) runWriter() {
    for message := range q.writeCh {
        q.storage.SaveMessage(message)
    }
}

// Enqueue message to queue in memory, and persist to db
func (q *Queue) Enqueue(message string) error {
    q.mu.Lock()
    defer q.mu.Unlock()

    // Add to in-memory queue
    q.messages = append(q.messages, message)

    // Send message to write channel
    q.writeCh <- message

    return nil
}

func (q *Queue) Dequeue() (string, error) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if len(q.messages) == 0 {
        return "", errors.New("Queue is empty!")
    }

    // Pop message from in-memory queue
    message := q.messages[0]
    q.messages = q.messages[1:]

    // Delete from storage
    err := q.storage.DeleteMessage(message)
    if err != nil {
        return "", err
    }
    
    return message, nil
}
