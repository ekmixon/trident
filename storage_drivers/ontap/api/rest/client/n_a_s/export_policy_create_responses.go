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

// ExportPolicyCreateReader is a Reader for the ExportPolicyCreate structure.
type ExportPolicyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportPolicyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExportPolicyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExportPolicyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExportPolicyCreateCreated creates a ExportPolicyCreateCreated with default headers values
func NewExportPolicyCreateCreated() *ExportPolicyCreateCreated {
	return &ExportPolicyCreateCreated{}
}

/* ExportPolicyCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ExportPolicyCreateCreated struct {
	Payload *models.ExportPolicyResponse
}

func (o *ExportPolicyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/nfs/export-policies][%d] exportPolicyCreateCreated  %+v", 201, o.Payload)
}
func (o *ExportPolicyCreateCreated) GetPayload() *models.ExportPolicyResponse {
	return o.Payload
}

func (o *ExportPolicyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportPolicyCreateDefault creates a ExportPolicyCreateDefault with default headers values
func NewExportPolicyCreateDefault(code int) *ExportPolicyCreateDefault {
	return &ExportPolicyCreateDefault{
		_statusCode: code,
	}
}

/* ExportPolicyCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1703952    | Invalid ruleset name provided. No spaces allowed in a ruleset name|
| 1703954    | Export policy does not exist |
| 1704049    | Invalid clientmatch: clientmatch lists require an effective cluster version of Data ONTAP 9.0 or later. Upgrade all nodes to Data ONTAP 9.0 or above to use features that operate on lists of clientmatch strings in export-policy rules |
| 1704055    | Export policies are only supported for data Vservers |
| 3277000    | Upgrade all nodes to Data ONTAP 9.0.0 or above to use krb5p as a security flavor in export-policy rules |
| 3277083    | User ID is not valid. Enter a value for User ID from 0 to 4294967295 |

*/
type ExportPolicyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the export policy create default response
func (o *ExportPolicyCreateDefault) Code() int {
	return o._statusCode
}

func (o *ExportPolicyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/nfs/export-policies][%d] export_policy_create default  %+v", o._statusCode, o.Payload)
}
func (o *ExportPolicyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExportPolicyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}