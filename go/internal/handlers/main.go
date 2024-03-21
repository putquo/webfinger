package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/putquo/webfinger/internal/jrd"
	"github.com/putquo/webfinger/internal/problemdetails"
)

func Webfinger(w http.ResponseWriter, r *http.Request) {
	withPrefix := r.URL.Query().Get("resource")
	if withPrefix == "" {
		problemdetails.New(problemdetails.BadRequest, http.StatusBadRequest, "The 'resource' query paramter is missing.").Write(w)
		return
	}

	parts := strings.Split(withPrefix, ":")
	if len(parts) != 2 || parts[0] != "attc" {
		problemdetails.New(problemdetails.BadRequest, http.StatusBadRequest, "The 'resource' query parameter is malformed.").Write(w)
		return
	}

	res, err := jrd.Resource(parts[1]).Jrd()
	if err != nil {
		if err.Error() == "resource not found" {
			problemdetails.New(problemdetails.NotFound, http.StatusNotFound, "The resource does not exist.").Write(w)
		} else {
			log.Printf("%v", err)
			problemdetails.New(problemdetails.InternalServerError, http.StatusInternalServerError, "An unexpected error occurred.")
		}
		return
	}

	w.Header().Set("Content-Type", "application/jrd+json")
	json.NewEncoder(w).Encode(res)
}
