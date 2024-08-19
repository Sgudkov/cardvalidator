# Card validator

The [main](https://github.com/Sgudkov/cardvalidator/blob/main/main.go#L40:L46) package is the entry point of the `cardvalidator` application. It is defined in the `main.go` file.

## Overview

The [main](https://github.com/Sgudkov/cardvalidator/blob/main/main.go#L40:L46) package sets up an HTTP server that listens on port 8080 and handles requests to the `/validate/` route.

## Functions

### main

The [main](https://github.com/Sgudkov/cardvalidator/blob/main/main.go#L40:L46) function is the entry point of the application. It sets up the HTTP server and starts listening for incoming requests.

### validateHandler

The `validateHandler` function is the handler for the `/validate/` route. It extracts the card number from the request URL, creates a [CardInfo](https://github.com/Sgudkov/cardvalidator/blob/main/main.go#L10:L12) struct, and performs some validation on the card number.

## Usage

To run the application, simply execute the [main](https://github.com/Sgudkov/cardvalidator/blob/main/main.go#L40:L46) function. The application will start listening on port 8080 and will handle incoming requests to the `/validate/` route.

## Example

Here's an example of how to use the application:

* Start the application by running `go run main.go`
* Open a web browser and navigate to `http://localhost:8080/validate/1234567890123456`
* The application will validate the card number and return a response indicating whether the card is valid or not.

## Code

**Card Validator API**
======================

Overview
--------

This is a simple API written in Go that validates card numbers using the Luhn algorithm. The API accepts HTTP requests to validate a card number and returns a response indicating whether the card number is valid or not.

**Endpoints**
------------

### `/validate/`

* **Method:** `POST`
* **Request Body:** `application/json`
* **Request Payload:** `{"Number": "card_number"}`
* **Response:**
	+ `200 OK` if the card number is valid
	+ `400 Bad Request` if the card number is invalid
	+ `500 Internal Server Error` if there's an error processing the request

**CardInfo Struct**
------------------

The `CardInfo` struct is used to represent the card number to be validated. It has a single field `Number` of type `string`.

```go
type CardInfo struct {
	Number string
}
```

**ValidateCardNumber Function**
------------------------------

The `ValidateCardNumber` function takes a card number as a string parameter and returns an error if the card number is invalid, nil otherwise. It uses the `goluhn` library to validate the card number using the Luhn algorithm.

```go
func ValidateCardNumber(cardNumber string) error {
	error := goluhn.Validate(cardNumber)
	return error
}
```

**validateHandler Function**
---------------------------

The `validateHandler` function handles HTTP requests to validate a card number. It takes an `http.ResponseWriter` and an `http.Request` as parameters, and returns no values. It writes to the `http.ResponseWriter` to send a response back to the client.

```go
func validateHandler(w http.ResponseWriter, r *http.Request) {
	// ...
}
```

**Usage**
---------

To use this API, send a `POST` request to the `/validate/` endpoint with a JSON payload containing the card number to be validated. For example:

```bash
curl -X POST \
  http://localhost:8080/validate/ \
  -H 'Content-Type: application/json' \
  -d '{"Number": "4111111111111111"}'
```

This should return a `200 OK` response if the card number is valid, or a `400 Bad Request` response if the card number is invalid.