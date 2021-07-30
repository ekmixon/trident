// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ExportPolicyDeleteReader is a Reader for the ExportPolicyDelete structure.
type ExportPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportPolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportPolicyDeleteOK creates a ExportPolicyDeleteOK with default headers values
func NewExportPolicyDeleteOK() *ExportPolicyDeleteOK {
	return &ExportPolicyDeleteOK{}
}

/* ExportPolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ExportPolicyDeleteOK struct {
}

func (o *ExportPolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{id}][%d] exportPolicyDeleteOK ", 200)
}

func (o *ExportPolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportPolicyDeleteDefault creates a ExportPolicyDeleteDefault with default headers values
func NewExportPolicyDeleteDefault(code int) *ExportPolicyDeleteDefault {
	return &ExportPolicyDeleteDefault{
		_statusCode: code,
	}
}

/* ExportPolicyDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1703944    | Failed to delete rule |
| 1703945    | Ruleset is in use by a volume.  It cannot be deleted until all volumes that refer to it are first deleted|
| 1703946    | Cannot determine if the ruleset is in use by a volume.  It cannot be deleted until all volumes that refer to it are first deleted|
| 1703947    | Cannot delete default ruleset.  This ruleset will be deleted when the owning Vserver is deleted|
| 1703952    | Invalid ruleset name provided. No spaces are allowed in a ruleset name|
| 1703953    | This ruleset is in use by a qtree export policy.  It cannot be deleted until all qtree policies that refer to it are first deleted|

*/
type ExportPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export policy delete default response
func (o *ExportPolicyDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExportPolicyDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nfs/export-policies/{id}][%d] export_policy_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ExportPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}