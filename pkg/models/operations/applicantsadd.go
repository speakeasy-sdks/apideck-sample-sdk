// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/utils"
	"net/http"
)

type ApplicantsAddSecurity struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=Authorization"`
}

func (o *ApplicantsAddSecurity) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type ApplicantsAddRequest struct {
	ApplicantInput shared.ApplicantInput `request:"mediaType=application/json"`
	// Include raw response. Mostly used for debugging purposes
	Raw *bool `default:"false" queryParam:"style=form,explode=true,name=raw"`
	// The ID of your Unify application
	XApideckAppID string `header:"style=simple,explode=false,name=x-apideck-app-id"`
	// ID of the consumer which you want to get or push data from
	XApideckConsumerID string `header:"style=simple,explode=false,name=x-apideck-consumer-id"`
	// Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.
	XApideckServiceID *string `header:"style=simple,explode=false,name=x-apideck-service-id"`
}

func (a ApplicantsAddRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApplicantsAddRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApplicantsAddRequest) GetApplicantInput() shared.ApplicantInput {
	if o == nil {
		return shared.ApplicantInput{}
	}
	return o.ApplicantInput
}

func (o *ApplicantsAddRequest) GetRaw() *bool {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *ApplicantsAddRequest) GetXApideckAppID() string {
	if o == nil {
		return ""
	}
	return o.XApideckAppID
}

func (o *ApplicantsAddRequest) GetXApideckConsumerID() string {
	if o == nil {
		return ""
	}
	return o.XApideckConsumerID
}

func (o *ApplicantsAddRequest) GetXApideckServiceID() *string {
	if o == nil {
		return nil
	}
	return o.XApideckServiceID
}

type ApplicantsAddResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Applicants
	CreateApplicantResponse *shared.CreateApplicantResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unexpected error
	UnexpectedErrorResponse *shared.UnexpectedErrorResponse
}

func (o *ApplicantsAddResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ApplicantsAddResponse) GetCreateApplicantResponse() *shared.CreateApplicantResponse {
	if o == nil {
		return nil
	}
	return o.CreateApplicantResponse
}

func (o *ApplicantsAddResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ApplicantsAddResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ApplicantsAddResponse) GetUnexpectedErrorResponse() *shared.UnexpectedErrorResponse {
	if o == nil {
		return nil
	}
	return o.UnexpectedErrorResponse
}
