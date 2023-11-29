package main

import (
	"net/http"

	"Stas-sH/test1.1/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/calculate", handlers.CalculateHandler)

	http.ListenAndServe(":8989", router)
}
