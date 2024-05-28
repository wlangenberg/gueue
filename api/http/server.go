package http

import (
    "encoding/json"
    "net/http"
    "gueue/pkg/queue"
)


type Server struct {
    queue   *queue.Queue
}


func NewServer(queue *queue.Queue) *Server {
	return &Server{queue: queue}
}


func (s *Server) enqueueHandler(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Message string `json:"message"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := s.queue.Enqueue(request.Message); err != nil {
        http.Error(w, "Failed to enque message", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}


func (s *Server) dequeueHandler(w http.ResponseWriter, r *http.Request) {
	message, err := s.queue.Dequeue()
	if err != nil {
		http.Error(w, "Failed to dequeue message", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{Message: message}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func (s *Server) Start(addr string) {
    http.HandleFunc("/enqueue", s.enqueueHandler)
    http.HandleFunc("/dequeue", s.dequeueHandler)
    http.ListenAndServe(addr, nil)
}
