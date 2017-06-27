package d_dns_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/muka/ddns/models"
)

// SaveRecordReader is a Reader for the SaveRecord structure.
type SaveRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SaveRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSaveRecordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSaveRecordOK creates a SaveRecordOK with default headers values
func NewSaveRecordOK() *SaveRecordOK {
	return &SaveRecordOK{}
}

/*SaveRecordOK handles this case with default header values.

SaveRecordOK save record o k
*/
type SaveRecordOK struct {
	Payload *models.APIRecord
}

func (o *SaveRecordOK) Error() string {
	return fmt.Sprintf("[POST /v1/record][%d] saveRecordOK  %+v", 200, o.Payload)
}

func (o *SaveRecordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
