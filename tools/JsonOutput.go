//package tools provides tools required for moviems to create directories, create/write files, checking on environment
//variables and a standard json output for api handlers
package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonOutput inherent both APIOutput and http.ResponseWriter
type JsonOutput struct {
	StatusCode int
	APIOutput
	http.ResponseWriter
}

// Create new instance of JsonOutput
func NewJsonOutput(w http.ResponseWriter) (out JsonOutput) {
	out.ResponseWriter = w
	return
}

// APIOutput is a standarize api output for microservice's handlers
type APIOutput struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

// JsonOutput.Print() will set header type on http.ResponseWriter and encode APIOutput on it
func (out *JsonOutput) Print() {
	out.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	out.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(out.ResponseWriter).Encode(out.APIOutput)
}

// JsonOutput.PrintIf() verifies if statement is true. if it is, add message to output message, set status code and call Print()
// it returns bool value
func (out *JsonOutput) PrintIf(isTrue bool, customMessage interface{}, statusCode int) bool {
	if isTrue {
		out.Success = false
		out.Message = customMessage
		out.StatusCode = statusCode
		out.printErrorWithHttpStatusCode()
	}
	return isTrue
}

// JsonOutput.PrintErrorIf() verifies if err is not nil. if it's not, add custom message or error message, set status code
// to output message and call Print()
// it returns bool value
func (out *JsonOutput) PrintErrorIf(err error, customMessage interface{}, statusCode int) (hasError bool) {
	if err != nil {
		hasError = true
		out.Message = customMessage
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
		out.StatusCode = statusCode
		out.printErrorWithHttpStatusCode()
	}
	return
}

// JsonOutput.printErrorWithHttpStatusCode() will output error message with Status Code
func (out *JsonOutput) printErrorWithHttpStatusCode() {
	out.Message = fmt.Sprint(http.StatusText(out.StatusCode)) + " " + fmt.Sprint(out.Message)
	http.Error(out.ResponseWriter, fmt.Sprint(out.Message), out.StatusCode)
}
