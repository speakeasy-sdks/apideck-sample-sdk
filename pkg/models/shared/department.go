// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/apideck-sample-sdk/pkg/utils"
	"time"
)

type Department struct {
	Code *string `json:"code,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The user who created the object.
	CreatedBy   *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// Department name
	Name *string `json:"name,omitempty"`
	// Parent ID
	ParentID *string `json:"parent_id,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

func (d Department) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Department) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Department) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Department) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Department) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Department) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Department) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Department) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Department) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *Department) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Department) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}
