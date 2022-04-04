package main

import (
	"net/http"
	"os"

	"github.com/hyorimitsu/hello-skaffold/simple/handlers"
)

func main() {
	code := start()
	os.Exit(code)
}

func start() int {
	handler := handlers.New()
	handlers.SubscribeHandlers(handler)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		return 1
	}
	return 0
}
