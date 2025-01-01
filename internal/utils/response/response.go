package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct{
	Status string `json:"status"`  //this line means that whenever this struct data is serialized into json this field will appear as status not Status
	Error string `json:"error"`
}

const (
	StatusOK="OK"
	StatusError="Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error{  //The actual data you want to send as a JSON object. It can be any type.
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)// Converts `data` to JSON and writes it to the response
}

func GeneralError(err error) Response{
	return Response{
		Status: StatusError,
		Error: err.Error(),
	}
}

func ValdationError(errs validator.ValidationErrors) Response{
	var errMsgs []string

	for _,err:=range errs{ //_ is the blank identifier, which discards the index returned by range.
		switch err.ActualTag(){
		case "required":
			errMsgs=append(errMsgs, fmt.Sprintf("field %s is required field",err.Field()))
		default:
			errMsgs=append(errMsgs,fmt.Sprintf("field %s is invalid",err.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error: strings.Join(errMsgs,", "),
	}
}