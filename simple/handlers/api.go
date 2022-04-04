package handlers

import (
	"fmt"
	"net/http"
)

// Api defined all api handlers.
type Api interface {
	GetGreet(w http.ResponseWriter, r *http.Request)
}

// ApiWrapper is an Api wrapper to do some processing before and after the actual processing of api.
type ApiWrapper struct {
	api Api
}

// GetGreet calls handler.
// If some processing is to be done before or after the actual processing, it should be described here.
func (aw *ApiWrapper) GetGreet(w http.ResponseWriter, r *http.Request) {
	aw.api.GetGreet(w, r)
}

// SubscribeHandlers subscribe all handlers
func SubscribeHandlers(api Api) {
	wrapper := ApiWrapper{
		api,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
	http.HandleFunc("/greet", wrapper.GetGreet)
}
