package verify

import "net/http"

type Handler interface {
	Verify() http.HandlerFunc
	Send() http.HandlerFunc
}

type VerifyHandler struct {}

func NewVerifyHandler(r *http.ServeMux) Handler {
	handler := &VerifyHandler{}

	r.HandleFunc("GET /verify", handler.Verify())
	r.HandleFunc("POST /send", handler.Send())

	return handler
}

func (h *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (h *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}