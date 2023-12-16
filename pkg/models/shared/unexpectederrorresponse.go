// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/apideck-go/pkg/utils"
)

type Two struct {
}

type DetailType string

const (
	DetailTypeStr DetailType = "str"
	DetailTypeTwo DetailType = "2"
)

// Detail - Contains parameter or domain specific information related to the error and why it occurred.
type Detail struct {
	Str *string
	Two *Two

	Type DetailType
}

func CreateDetailStr(str string) Detail {
	typ := DetailTypeStr

	return Detail{
		Str:  &str,
		Type: typ,
	}
}

func CreateDetailTwo(two Two) Detail {
	typ := DetailTypeTwo

	return Detail{
		Two:  &two,
		Type: typ,
	}
}

func (u *Detail) UnmarshalJSON(data []byte) error {

	two := Two{}
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = &two
		u.Type = DetailTypeTwo
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = DetailTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Detail) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UnexpectedErrorResponse struct {
	// Contains parameter or domain specific information related to the error and why it occurred.
	Detail *Detail `json:"detail,omitempty"`
	// Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Error *string `json:"error,omitempty"`
	// A human-readable message providing more details about the error.
	Message *string `json:"message,omitempty"`
	// Link to documentation of error type
	Ref *string `json:"ref,omitempty"`
	// HTTP status code
	StatusCode *float64 `json:"status_code,omitempty"`
	// The type of error returned
	TypeName *string `json:"type_name,omitempty"`
}

func (o *UnexpectedErrorResponse) GetDetail() *Detail {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *UnexpectedErrorResponse) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *UnexpectedErrorResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UnexpectedErrorResponse) GetRef() *string {
	if o == nil {
		return nil
	}
	return o.Ref
}

func (o *UnexpectedErrorResponse) GetStatusCode() *float64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *UnexpectedErrorResponse) GetTypeName() *string {
	if o == nil {
		return nil
	}
	return o.TypeName
}
