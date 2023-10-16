// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/utils"
)

// NotFoundResponseDetail2 - Contains parameter or domain specific information related to the error and why it occurred.
type NotFoundResponseDetail2 struct {
}

var _ error = &NotFoundResponseDetail2{}

func (e *NotFoundResponseDetail2) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type NotFoundResponseDetail9Type string

const (
	NotFoundResponseDetail9TypeStr                     NotFoundResponseDetail9Type = "str"
	NotFoundResponseDetail9TypeNotFoundResponseDetail2 NotFoundResponseDetail9Type = "NotFoundResponse_detail_2"
)

type NotFoundResponseDetail9 struct {
	Str                     *string
	NotFoundResponseDetail2 *NotFoundResponseDetail2

	Type NotFoundResponseDetail9Type
}

func CreateNotFoundResponseDetail9Str(str string) NotFoundResponseDetail9 {
	typ := NotFoundResponseDetail9TypeStr

	return NotFoundResponseDetail9{
		Str:  &str,
		Type: typ,
	}
}

func CreateNotFoundResponseDetail9NotFoundResponseDetail2(notFoundResponseDetail2 NotFoundResponseDetail2) NotFoundResponseDetail9 {
	typ := NotFoundResponseDetail9TypeNotFoundResponseDetail2

	return NotFoundResponseDetail9{
		NotFoundResponseDetail2: &notFoundResponseDetail2,
		Type:                    typ,
	}
}

func (u *NotFoundResponseDetail9) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = NotFoundResponseDetail9TypeStr
		return nil
	}

	notFoundResponseDetail2 := new(NotFoundResponseDetail2)
	if err := utils.UnmarshalJSON(data, &notFoundResponseDetail2, "", true, true); err == nil {
		u.NotFoundResponseDetail2 = notFoundResponseDetail2
		u.Type = NotFoundResponseDetail9TypeNotFoundResponseDetail2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u NotFoundResponseDetail9) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.NotFoundResponseDetail2 != nil {
		return utils.MarshalJSON(u.NotFoundResponseDetail2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type NotFoundResponse9 struct {
	// Contains parameter or domain specific information related to the error and why it occurred.
	Detail *NotFoundResponseDetail9 `json:"detail,omitempty"`
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

var _ error = &NotFoundResponse9{}

func (e *NotFoundResponse9) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
