package moviems_lib

import (
	"net/http"
	"errors"
)

func ValidateToken(req *http.Request, myToken string) (error) {
	header := "x-access-token"

	hToken := req.Header.Get(header)
	if hToken == "" {
		return errors.New("Header: " + header + "  Not Found")
	}
	if hToken != myToken {
		return errors.New("Invalid Token Provided")
	}

	return nil
}
