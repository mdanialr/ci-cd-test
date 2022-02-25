package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	res := struct {
		Message string `json:"message"`
	}{Message: "Hello From Earth!!"}
	js, err := json.Marshal(res)
	if err != nil {
		log.Println("ERROR:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
