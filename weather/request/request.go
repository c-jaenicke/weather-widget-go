package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Make a request to the given URL and return a []byte with the response body
// Logs fatal if response body is empty
// Prints error if status code is not 200
// Parameters: url string
// Returns: responseData []byte
func Request(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while making request: "+err.Error())
		os.Exit(1)
	}

	if response.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Response was not 200 OK, got: "+response.Status)
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Response body was empty. ")
		os.Exit(1)
	}

	return responseData
}
