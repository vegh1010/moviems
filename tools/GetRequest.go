package tools

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RequestResult struct {
	ResponseBody []byte
	StatusCode   int
	StatusString string
}

const APPLICATION_JSON = "application/json"

func GetRequest(urlPath, accessToken string) (result RequestResult, err error) {
	if accessToken == "" {
		err = errors.New("Access Token Required")
		return
	}

	var responseBody []byte
	var statusCode int
	var statusString string
	dataBuffer := new(bytes.Buffer)

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

	statusCode = newResponse.StatusCode
	statusString = newResponse.Status
	responseBody, err = ioutil.ReadAll(newResponse.Body)
	if err != nil {
		return
	}

	result.StatusCode = statusCode
	result.StatusString = statusString
	result.ResponseBody = responseBody

	return
}
