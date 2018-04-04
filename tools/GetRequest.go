//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RequestResult is used to format http request output
type RequestResult struct {
	ResponseBody []byte
	StatusCode   int
	StatusString string
}

const APPLICATION_JSON = "application/json"

// GetRequest will accept the url and access token to do a GET request on the specified url
// it returns RequestResult of the http request info and error if an issue occurred
func GetRequest(urlPath, accessToken string) (result RequestResult, err error) {
	//validate access token
	if accessToken == "" {
		err = errors.New("Access Token Required")
		return
	}

	dataBuffer := new(bytes.Buffer)

	//format request url
	requestUrl, err := url.ParseRequestURI(urlPath)
	if err != nil {
		err = errors.New("Invalid URL supplied. URL: " + urlPath + ". Error: " + err.Error())
		return
	}

	//new instance request
	newRequest, err := http.NewRequest("GET", requestUrl.String(), dataBuffer)
	if err != nil {
		return
	}

	// Set request content types
	newRequest.Header.Set("Content-Type", APPLICATION_JSON)

	//set request header x-access-token
	newRequest.Header.Set("x-access-token", accessToken)

	//create a new http client and execute http request
	newClient := &http.Client{}
	newResponse := new(http.Response)
	newResponse, err = newClient.Do(newRequest)
	if err != nil {
		err = errors.New("Failed to connect: " + err.Error())
		return
	}
	defer newResponse.Body.Close()

	//set values to RequestResult
	result.StatusCode = newResponse.StatusCode
	result.StatusString = newResponse.Status
	result.ResponseBody, err = ioutil.ReadAll(newResponse.Body)
	if err != nil {
		return
	}

	return
}
