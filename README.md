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
