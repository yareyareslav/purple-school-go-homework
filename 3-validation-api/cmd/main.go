package cmd

import (
	"net/http"
	"purple-school-go/homework/3-validation-api/internal/verify"
)

func main() {
	router := http.NewServeMux()

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}

	verify.NewVerifyHandler(router)

	server.ListenAndServe()
}