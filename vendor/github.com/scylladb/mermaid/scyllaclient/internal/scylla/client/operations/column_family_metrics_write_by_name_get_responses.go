// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ColumnFamilyMetricsWriteByNameGetReader is a Reader for the ColumnFamilyMetricsWriteByNameGet structure.
type ColumnFamilyMetricsWriteByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsWriteByNameGetOK creates a ColumnFamilyMetricsWriteByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteByNameGetOK() *ColumnFamilyMetricsWriteByNameGetOK {
	return &ColumnFamilyMetricsWriteByNameGetOK{}
}

/*ColumnFamilyMetricsWriteByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteByNameGetOK column family metrics write by name get o k
*/
type ColumnFamilyMetricsWriteByNameGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsWriteByNameGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/write/{name}][%d] columnFamilyMetricsWriteByNameGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsWriteByNameGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}