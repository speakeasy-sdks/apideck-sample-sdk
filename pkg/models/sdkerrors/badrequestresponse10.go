// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/utils"
)

// BadRequestResponseDetail2 - Contains parameter or domain specific information related to the error and why it occurred.
type BadRequestResponseDetail2 struct {
}

var _ error = &BadRequestResponseDetail2{}

func (e *BadRequestResponseDetail2) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

type BadRequestResponseDetail10Type string

const (
	BadRequestResponseDetail10TypeStr                       BadRequestResponseDetail10Type = "str"
	BadRequestResponseDetail10TypeBadRequestResponseDetail2 BadRequestResponseDetail10Type = "BadRequestResponse_detail_2"
)

type BadRequestResponseDetail10 struct {
	Str                       *string
	BadRequestResponseDetail2 *BadRequestResponseDetail2

	Type BadRequestResponseDetail10Type
}

func CreateBadRequestResponseDetail10Str(str string) BadRequestResponseDetail10 {
	typ := BadRequestResponseDetail10TypeStr

	return BadRequestResponseDetail10{
		Str:  &str,
		Type: typ,
	}
}

func CreateBadRequestResponseDetail10BadRequestResponseDetail2(badRequestResponseDetail2 BadRequestResponseDetail2) BadRequestResponseDetail10 {
	typ := BadRequestResponseDetail10TypeBadRequestResponseDetail2

	return BadRequestResponseDetail10{
		BadRequestResponseDetail2: &badRequestResponseDetail2,
		Type:                      typ,
	}
}

func (u *BadRequestResponseDetail10) UnmarshalJSON(data []byte) error {

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = BadRequestResponseDetail10TypeStr
		return nil
	}

	badRequestResponseDetail2 := new(BadRequestResponseDetail2)
	if err := utils.UnmarshalJSON(data, &badRequestResponseDetail2, "", true, true); err == nil {
		u.BadRequestResponseDetail2 = badRequestResponseDetail2
		u.Type = BadRequestResponseDetail10TypeBadRequestResponseDetail2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u BadRequestResponseDetail10) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.BadRequestResponseDetail2 != nil {
		return utils.MarshalJSON(u.BadRequestResponseDetail2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type BadRequestResponse10 struct {
	// Contains parameter or domain specific information related to the error and why it occurred.
	Detail *BadRequestResponseDetail10 `json:"detail,omitempty"`
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

var _ error = &BadRequestResponse10{}

func (e *BadRequestResponse10) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
