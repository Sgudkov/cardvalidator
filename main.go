package main

import (
	"cardvalidator/cardvalidator"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CardInfo struct {
	Number string
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cardInfo := &CardInfo{}

	err = json.Unmarshal(body, cardInfo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = cardvalidator.ValidateCardNumber(cardInfo.Number)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "ok")
}

func main() {

	http.HandleFunc("/validate/", validateHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
