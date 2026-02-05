package payload

import (
	"encoding/json"
	"fmt"
)

type ClientError struct {
	ClientError ErrorStruct `json:"error"`
}

func (e ClientError) Error() string {
	return fmt.Sprintf("%+v", e.ClientError)
}

type ErrorStruct struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ErrorStruct) Error() string {
	return e.Message
}

func ParseErrorAsJson[T error](data []byte, errorJson *T) error {
	err := json.Unmarshal(data, errorJson)
	if err != nil {
		// return errormodel.ClientErrorDefaultCode(
		// 	fmt.Sprintf("error response: %s -- unable to unmarhal json: %v", data, err),
		// )
	}
	return nil
}
