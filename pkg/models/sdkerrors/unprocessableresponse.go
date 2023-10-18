// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type UnprocessableResponse struct {
	// Contains parameter or domain specific information related to the error and why it occurred.
	Detail *string `json:"detail,omitempty"`
	// Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Error_ *string `json:"error,omitempty"`
	// A human-readable message providing more details about the error.
	Message *string `json:"message,omitempty"`
	// Link to documentation of error type
	Ref *string `json:"ref,omitempty"`
	// HTTP status code
	StatusCode *float64 `json:"status_code,omitempty"`
	// The type of error returned
	TypeName *string `json:"type_name,omitempty"`
}

var _ error = &UnprocessableResponse{}

func (e *UnprocessableResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
