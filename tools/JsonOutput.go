package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonOutput struct {
	APIOutput
	http.ResponseWriter
}

func NewJsonOutput(w http.ResponseWriter) (out JsonOutput) {
	out.ResponseWriter = w
	return
}

//standarize api output for microservice and some basic functionality for condition output
type APIOutput struct {
	Success    bool        `json:"success"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

//print result to response writer and set header
func (out *JsonOutput) Print() {
	out.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	out.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(out.ResponseWriter).Encode(out.APIOutput)
}

//print if statement is true with custom message
func (out *JsonOutput) PrintIf(isTrue bool, customMessage interface{}) bool {
	if isTrue {
		out.Success = false
		out.Message = customMessage
		out.Print()
	}
	return isTrue
}

func (out *JsonOutput) PrintErrorIf(err error, customMessage interface{}) (hasError bool) {
	if err != nil {
		hasError = true
		out.Message = customMessage
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
		out.Print()
	}
	return
}

func (out *JsonOutput) printErrorWithHttpStatusCode(err error, statusCode int) {
	if err != nil {
		// if message not specified use err.Error() as message
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
	}
	out.Message = fmt.Sprint(http.StatusText(statusCode)) + " " + fmt.Sprint(out.Message)
	http.Error(out.ResponseWriter, fmt.Sprint(out.Message), statusCode)
}

func (out *JsonOutput) PrintError400(err error) {
	out.printErrorWithHttpStatusCode(err, 400)
}

func (out *JsonOutput) PrintError500(err error) {
	out.printErrorWithHttpStatusCode(err, 500)
}
