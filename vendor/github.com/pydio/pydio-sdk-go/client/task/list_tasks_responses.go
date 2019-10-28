// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/*ListTasksOK handles this case with default header values.

Successful Response
*/
type ListTasksOK struct {
	Payload []*models.Task
}

func (o *ListTasksOK) Error() string {
	return fmt.Sprintf("[GET /tasks][%d] listTasksOK  %+v", 200, o.Payload)
}

func (o *ListTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
