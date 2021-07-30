// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SQLOnSanNewIgroups The list of initiator groups to create.
//
// swagger:model sql_on_san_new_igroups
type SQLOnSanNewIgroups struct {

	// initiators
	Initiators []string `json:"initiators,omitempty"`

	// The name of the new initiator group.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name"`

	// The name of the host OS accessing the application. The default value is the host OS that is running the application.
	// Enum: [hyper_v vmware windows]
	OsType *string `json:"os_type,omitempty"`

	// The protocol of the new initiator group.
	// Enum: [fcp iscsi mixed]
	Protocol *string `json:"protocol,omitempty"`
}

// Validate validates this sql on san new igroups
func (m *SQLOnSanNewIgroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SQLOnSanNewIgroups) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

var sqlOnSanNewIgroupsTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hyper_v","vmware","windows"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanNewIgroupsTypeOsTypePropEnum = append(sqlOnSanNewIgroupsTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// sql_on_san_new_igroups
	// SQLOnSanNewIgroups
	// os_type
	// OsType
	// hyper_v
	// END RIPPY DEBUGGING
	// SQLOnSanNewIgroupsOsTypeHyperv captures enum value "hyper_v"
	SQLOnSanNewIgroupsOsTypeHyperv string = "hyper_v"

	// BEGIN RIPPY DEBUGGING
	// sql_on_san_new_igroups
	// SQLOnSanNewIgroups
	// os_type
	// OsType
	// vmware
	// END RIPPY DEBUGGING
	// SQLOnSanNewIgroupsOsTypeVmware captures enum value "vmware"
	SQLOnSanNewIgroupsOsTypeVmware string = "vmware"

	// BEGIN RIPPY DEBUGGING
	// sql_on_san_new_igroups
	// SQLOnSanNewIgroups
	// os_type
	// OsType
	// windows
	// END RIPPY DEBUGGING
	// SQLOnSanNewIgroupsOsTypeWindows captures enum value "windows"
	SQLOnSanNewIgroupsOsTypeWindows string = "windows"
)

// prop value enum
func (m *SQLOnSanNewIgroups) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanNewIgroupsTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanNewIgroups) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

var sqlOnSanNewIgroupsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","iscsi","mixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sqlOnSanNewIgroupsTypeProtocolPropEnum = append(sqlOnSanNewIgroupsTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN RIPPY DEBUGGING
	// sql_on_san_new_igroups
	// SQLOnSanNewIgroups
	// protocol
	// Protocol
	// fcp
	// END RIPPY DEBUGGING
	// SQLOnSanNewIgroupsProtocolFcp captures enum value "fcp"
	SQLOnSanNewIgroupsProtocolFcp string = "fcp"

	// BEGIN RIPPY DEBUGGING
	// sql_on_san_new_igroups
	// SQLOnSanNewIgroups
	// protocol
	// Protocol
	// iscsi
	// END RIPPY DEBUGGING
	// SQLOnSanNewIgroupsProtocolIscsi captures enum value "iscsi"
	SQLOnSanNewIgroupsProtocolIscsi string = "iscsi"

	// BEGIN RIPPY DEBUGGING
	// sql_on_san_new_igroups
	// SQLOnSanNewIgroups
	// protocol
	// Protocol
	// mixed
	// END RIPPY DEBUGGING
	// SQLOnSanNewIgroupsProtocolMixed captures enum value "mixed"
	SQLOnSanNewIgroupsProtocolMixed string = "mixed"
)

// prop value enum
func (m *SQLOnSanNewIgroups) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sqlOnSanNewIgroupsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SQLOnSanNewIgroups) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sql on san new igroups based on context it is used
func (m *SQLOnSanNewIgroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLOnSanNewIgroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLOnSanNewIgroups) UnmarshalBinary(b []byte) error {
	var res SQLOnSanNewIgroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HELLO RIPPY