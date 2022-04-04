package handlers

// Handler is api handlers.
type Handler struct {
	greetHandler
}

// New returns new Handler.
func New() Handler {
	h := Handler{}

	// (Omitted)

	return h
}
