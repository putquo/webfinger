package main

import (
	"fmt"
	"net/http"

	"github.com/putquo/webfinger/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Webfinger)
	fmt.Println("starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
