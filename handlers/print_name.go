package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PrintName(w http.ResponseWriter, req *http.Request) {
	param := strings.TrimPrefix(req.URL.Path, "/print/")

	if param == "" {
		param = "there"
	}

	res := struct {
		Message string `json:"message"`
	}{Message: fmt.Sprintf("Good day %s!", param)}

	js, err := json.Marshal(res)
	if err != nil {
		log.Println("ERROR:", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
