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

// ColumnFamilyMetricsRecentBloomFilterFalseRatioGetReader is a Reader for the ColumnFamilyMetricsRecentBloomFilterFalseRatioGet structure.
type ColumnFamilyMetricsRecentBloomFilterFalseRatioGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK creates a ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK with default headers values
func NewColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK() *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK {
	return &ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK{}
}

/*ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK handles this case with default header values.

ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK column family metrics recent bloom filter false ratio get o k
*/
type ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK struct {
	Payload interface{}
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK) Error() string {
	return fmt.Sprintf("[GET /column_family/metrics/recent_bloom_filter_false_ratio][%d] columnFamilyMetricsRecentBloomFilterFalseRatioGetOK  %+v", 200, o.Payload)
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ColumnFamilyMetricsRecentBloomFilterFalseRatioGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}