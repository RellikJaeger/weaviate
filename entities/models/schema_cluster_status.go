// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchemaClusterStatus Indicates the health of the schema in a cluster.
//
// swagger:model SchemaClusterStatus
type SchemaClusterStatus struct {

	// Contains the sync check error if one occurred
	Error string `json:"error,omitempty"`

	// True if the cluster is in sync, false if there is an issue (see error).
	Healthy bool `json:"healthy"`

	// Hostname of the coordinating node, i.e. the one that received the cluster. This can be useful information if the error message contains phrases such as 'other nodes agree, but local does not', etc.
	Hostname string `json:"hostname,omitempty"`

	// The cluster check at startup can be ignored (to recover from an out-of-sync situation).
	IgnoreSchemaSync bool `json:"ignoreSchemaSync"`

	// Number of nodes that participated in the sync check
	NodeCount int64 `json:"nodeCount,omitempty"`
}

// Validate validates this schema cluster status
func (m *SchemaClusterStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schema cluster status based on context it is used
func (m *SchemaClusterStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchemaClusterStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemaClusterStatus) UnmarshalBinary(b []byte) error {
	var res SchemaClusterStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
