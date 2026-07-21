package verify

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/smtp"
	"purple-school-go/homework/3-validation-api/pkg/req"
	"purple-school-go/homework/3-validation-api/pkg/res"

	"github.com/jordan-wright/email"
)

type Handler interface {
	Verify() http.HandlerFunc
	Send() http.HandlerFunc
}

type VerifyHandler struct {
	store string
}

func NewVerifyHandler(r *http.ServeMux) Handler {
	handler := &VerifyHandler{}

	r.HandleFunc("GET /verify/{hash}", handler.Verify())
	r.HandleFunc("POST /send", handler.Send())

	return handler
}

func (h *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if (h.store == hash) {
			h.store = ""
			res.Json(w, true, 200)
		} else {
			res.Json(w, false, 404)
		}
	}
}

func (h *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}

		hash := sha256.Sum256([]byte(body.Address))
		newHash := hex.EncodeToString(hash[:])
		
		e := email.NewEmail()
		e.From = "Purple School <" + body.Email + ">"
		e.To = []string{body.Address}
		e.Subject = "Hello!"
		e.Text = []byte("http://localhost:8081/verify/" + newHash)
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", body.Email, body.Password, "smtp.gmail.com"))
		
		h.store = newHash
	}
}