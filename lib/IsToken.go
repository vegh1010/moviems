// package moviems_lib is a moviems library which contains structs/functions related to moviems.
package moviems_lib

import (
	"net/http"
	"errors"
)

// ValidateToken will get access token header from http.Request and compare with another token.
// Token will be validated and a proper error will be return if any.
// It returns error if issue occurred.
func ValidateToken(req *http.Request, myToken string) (error) {
	header := "x-access-token"

	//get header from http.Request
	hToken := req.Header.Get(header)

	//validate token
	if hToken == "" {
		return errors.New("Header: " + header + "  Not Found")
	}

	//compare token
	if hToken != myToken {
		return errors.New("Invalid Token Provided")
	}

	return nil
}
