package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ShiraazMoollatjie/goluhn"
)

type CardInfo struct {
	Number string
}

// validateHandler handles HTTP requests to validate a card number.
//
// It takes an http.ResponseWriter and an http.Request as parameters.
// It returns no values, but writes to the http.ResponseWriter.
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

	err = ValidateCardNumber(cardInfo.Number)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "ok")
}

// ValidateCardNumber validates a card number using the Luhn algorithm.
//
// It takes a card number as a string parameter.
// Returns an error if the card number is invalid, nil otherwise.
func ValidateCardNumber(cardNumber string) error {
	error := goluhn.Validate(cardNumber)
	return error
}

// main is the entry point of the program.
//
// It sets up the HTTP server to handle requests to the "/validate/" endpoint
// using the validateHandler function. Then, it prints a message to the console
// indicating that the server is starting, and starts the server on port 8080.
func main() {

	http.HandleFunc("/validate/", validateHandler)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
