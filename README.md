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

Here is the code for the [main](https://github.com/Sgudkov/cardvalidator/blob/main/main.go#L40:L46) package:
```go
func main() {
	http.HandleFunc("/validate/", validateHandler)
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
