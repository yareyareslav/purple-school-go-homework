package randomApi

import (
	"fmt"
	"math/rand"
	"net/http"
)

func getRandomNumber() uint8 {
	return uint8(rand.Intn(6) + 1)
}

type RandomApiHttpHandler struct {}

func NewRandomApiHttpHandler(router *http.ServeMux) *RandomApiHttpHandler {
	handler := &RandomApiHttpHandler{}
	router.HandleFunc("/random", handler.GetRandomNumber)
	return handler
}

func (h *RandomApiHttpHandler) GetRandomNumber(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%d", getRandomNumber())))
}