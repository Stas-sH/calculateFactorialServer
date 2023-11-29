package handlers

import (
	"encoding/json"
	"net/http"

	"Stas-sH/test1.1/internal/data"
	"Stas-sH/test1.1/pkg/factorial"

	"github.com/julienschmidt/httprouter"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var inputData data.InputData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&inputData); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if inputData.A < 0 || inputData.B < 0 {
		http.Error(w, `{"error":"Incorrect input"}`, http.StatusBadRequest)
		return
	}

	resultChA := make(chan int)
	resultChB := make(chan int)
	go factorial.CalculateFactorial(inputData.A, resultChA)
	go factorial.CalculateFactorial(inputData.B, resultChB)

	resultA, resultB := <-resultChA, <-resultChB

	outData := data.OutputData{
		ResultA: resultA,
		ResultB: resultB,
	}

	resultJSON, err := json.Marshal(outData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}
