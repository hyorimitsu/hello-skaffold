package handlers

import (
	"fmt"
	"net/http"
)

type greetHandler struct {
	// (Omitted)
}

// GetGreet get greet string.
func (h greetHandler) GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Skaffold!")
}
