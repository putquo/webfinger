package problemdetails

import (
	"encoding/json"
	"net/http"
)

type ProblemDetails struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type Type int

const (
	BadRequest Type = iota
	NotFound
	InternalServerError
)

func statusType(status int) string {
	switch status {
	case 400:
		return "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	case 404:
		return "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	case 500:
		return "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	default:
		return "unknown"
	}
}

func New(status int, detail string) *ProblemDetails {
	return &ProblemDetails{
		Type:   statusType(status),
		Title:  http.StatusText(status),
		Status: status,
		Detail: detail,
	}
}

func (p ProblemDetails) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(p.Status)
	json.NewEncoder(w).Encode(p)
}
