package storage



import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

// The Storage struct that represent sqllite3 storage
type Storage struct {
    db *sql.DB
}


// NewStorage inits a new storage with an sqllite3 database
func NewStorage(dataSourceName string) (*Storage, error) {
    db, err := sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Ensure the table exists
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS messages (
            id INTEGER PRIMARY KEY,
            message TEXT NOT NULL
        )`)
    if err != nil {
        return nil, err
    }

    return &Storage{db: db}, nil
}


func (s *Storage) SaveMessage(message string) error {
    _, err := s.db.Exec("INSERT INTO messages (message) VALUES (?)", message)
    return err
}

func (s *Storage) DeleteMessage(message string) error {
    _, err := s.db.Exec("DELETE FROM messages WHERE message = ?", message)
    return err
}

func (s *Storage) RetrieveMessages(message string) ([]string, error) {
    rows, err := s.db.Query("SELECT message from messages")
    if err != nil {
        return nil, err
    }    
    defer rows.Close()

    var messages []string
    for rows.Next() {
        var message string
        if err := rows.Scan(&message); err != nil {
            return nil, err
        }
        
        messages = append(messages, message)
    }

    return messages, nil
}
