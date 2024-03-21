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

func (t Type) into() string {
	switch t {
	case BadRequest:
		return "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	case NotFound:
		return "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	case InternalServerError:
		return "https://datatracker.ietf.org/doc/html/rfc7231#section-6.5.1"
	default:
		return "unknown"
	}
}

func New(pType Type, status int, detail string) *ProblemDetails {
	return &ProblemDetails{
		Type:   pType.into(),
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
