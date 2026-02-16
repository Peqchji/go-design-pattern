package middleware

import (
	"net/http"
)

type StatusRecorder struct {
	http.ResponseWriter // Embedded: We inherit Write(), Header(), etc.
	StatusCode int
}

func NewStatusRecorder(writer http.ResponseWriter, code int) *StatusRecorder {
	return &StatusRecorder{
		ResponseWriter: writer,
		StatusCode: code,
	}
} 

func (r *StatusRecorder) WriteHeader(code int) {
	r.StatusCode = code
	r.ResponseWriter.WriteHeader(code) // Forward to the original writer
}

func (r *StatusRecorder) GetRecord() int {
	return r.StatusCode
}