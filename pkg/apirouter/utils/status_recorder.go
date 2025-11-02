package utils

import "net/http"

type StatusRecorder struct {
	http.ResponseWriter
	Status  int
	Flusher http.Flusher
}

func (rec *StatusRecorder) WriteHeader(code int) {
	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *StatusRecorder) Flush() {
	rec.Flusher.Flush()
}
