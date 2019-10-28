// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// ChangesReader is a Reader for the Changes structure.
type ChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangesOK creates a ChangesOK with default headers values
func NewChangesOK() *ChangesOK {
	return &ChangesOK{}
}

/*ChangesOK handles this case with default header values.

Successful Response
*/
type ChangesOK struct {
	Payload []*models.ChangesOKBodyItems
}

func (o *ChangesOK) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/changes/{sequenceId}][%d] changesOK  %+v", 200, o.Payload)
}

func (o *ChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
