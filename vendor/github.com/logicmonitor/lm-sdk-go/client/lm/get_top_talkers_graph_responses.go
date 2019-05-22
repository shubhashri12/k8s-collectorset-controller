// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// GetTopTalkersGraphReader is a Reader for the GetTopTalkersGraph structure.
type GetTopTalkersGraphReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTopTalkersGraphReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTopTalkersGraphOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTopTalkersGraphDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTopTalkersGraphOK creates a GetTopTalkersGraphOK with default headers values
func NewGetTopTalkersGraphOK() *GetTopTalkersGraphOK {
	return &GetTopTalkersGraphOK{}
}

/*GetTopTalkersGraphOK handles this case with default header values.

successful operation
*/
type GetTopTalkersGraphOK struct {
	Payload *models.GraphPlot
}

func (o *GetTopTalkersGraphOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/topTalkersGraph][%d] getTopTalkersGraphOK  %+v", 200, o.Payload)
}

func (o *GetTopTalkersGraphOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GraphPlot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTopTalkersGraphDefault creates a GetTopTalkersGraphDefault with default headers values
func NewGetTopTalkersGraphDefault(code int) *GetTopTalkersGraphDefault {
	return &GetTopTalkersGraphDefault{
		_statusCode: code,
	}
}

/*GetTopTalkersGraphDefault handles this case with default header values.

Error
*/
type GetTopTalkersGraphDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get top talkers graph default response
func (o *GetTopTalkersGraphDefault) Code() int {
	return o._statusCode
}

func (o *GetTopTalkersGraphDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/topTalkersGraph][%d] getTopTalkersGraph default  %+v", o._statusCode, o.Payload)
}

func (o *GetTopTalkersGraphDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
