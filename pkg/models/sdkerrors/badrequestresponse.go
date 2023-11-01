// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"errors"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
)

// BadRequestResponseDetail2 - Contains parameter or domain specific information related to the error and why it occurred.
type BadRequestResponseDetail2 struct {
}

type BadRequestResponseDetailType string

const (
	BadRequestResponseDetailTypeStr                       BadRequestResponseDetailType = "str"
	BadRequestResponseDetailTypeBadRequestResponseDetail2 BadRequestResponseDetailType = "BadRequestResponse_detail_2"
)

type BadRequestResponseDetail struct {
	Str                       *string
	BadRequestResponseDetail2 *BadRequestResponseDetail2

	Type BadRequestResponseDetailType
}

func CreateBadRequestResponseDetailStr(str string) BadRequestResponseDetail {
	typ := BadRequestResponseDetailTypeStr

	return BadRequestResponseDetail{
		Str:  &str,
		Type: typ,
	}
}

func CreateBadRequestResponseDetailBadRequestResponseDetail2(badRequestResponseDetail2 BadRequestResponseDetail2) BadRequestResponseDetail {
	typ := BadRequestResponseDetailTypeBadRequestResponseDetail2

	return BadRequestResponseDetail{
		BadRequestResponseDetail2: &badRequestResponseDetail2,
		Type:                      typ,
	}
}

func (u *BadRequestResponseDetail) UnmarshalJSON(data []byte) error {

	badRequestResponseDetail2 := BadRequestResponseDetail2{}
	if err := utils.UnmarshalJSON(data, &badRequestResponseDetail2, "", true, true); err == nil {
		u.BadRequestResponseDetail2 = &badRequestResponseDetail2
		u.Type = BadRequestResponseDetailTypeBadRequestResponseDetail2
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = BadRequestResponseDetailTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u BadRequestResponseDetail) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.BadRequestResponseDetail2 != nil {
		return utils.MarshalJSON(u.BadRequestResponseDetail2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type BadRequestResponse struct {
	// Contains parameter or domain specific information related to the error and why it occurred.
	Detail *BadRequestResponseDetail `json:"detail,omitempty"`
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

var _ error = &BadRequestResponse{}

func (e *BadRequestResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
