// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimCableTerminationsCreateReader is a Reader for the DcimCableTerminationsCreate structure.
type DcimCableTerminationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimCableTerminationsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCableTerminationsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCableTerminationsCreateCreated creates a DcimCableTerminationsCreateCreated with default headers values
func NewDcimCableTerminationsCreateCreated() *DcimCableTerminationsCreateCreated {
	return &DcimCableTerminationsCreateCreated{}
}

/*
DcimCableTerminationsCreateCreated describes a response with status code 201, with default header values.

DcimCableTerminationsCreateCreated dcim cable terminations create created
*/
type DcimCableTerminationsCreateCreated struct {
	Payload *models.CableTermination
}

// IsSuccess returns true when this dcim cable terminations create created response has a 2xx status code
func (o *DcimCableTerminationsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations create created response has a 3xx status code
func (o *DcimCableTerminationsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations create created response has a 4xx status code
func (o *DcimCableTerminationsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations create created response has a 5xx status code
func (o *DcimCableTerminationsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations create created response a status code equal to that given
func (o *DcimCableTerminationsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimCableTerminationsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/cable-terminations/][%d] dcimCableTerminationsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimCableTerminationsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/cable-terminations/][%d] dcimCableTerminationsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimCableTerminationsCreateCreated) GetPayload() *models.CableTermination {
	return o.Payload
}

func (o *DcimCableTerminationsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CableTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCableTerminationsCreateDefault creates a DcimCableTerminationsCreateDefault with default headers values
func NewDcimCableTerminationsCreateDefault(code int) *DcimCableTerminationsCreateDefault {
	return &DcimCableTerminationsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimCableTerminationsCreateDefault describes a response with status code -1, with default header values.

DcimCableTerminationsCreateDefault dcim cable terminations create default
*/
type DcimCableTerminationsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cable terminations create default response
func (o *DcimCableTerminationsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cable terminations create default response has a 2xx status code
func (o *DcimCableTerminationsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cable terminations create default response has a 3xx status code
func (o *DcimCableTerminationsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cable terminations create default response has a 4xx status code
func (o *DcimCableTerminationsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cable terminations create default response has a 5xx status code
func (o *DcimCableTerminationsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cable terminations create default response a status code equal to that given
func (o *DcimCableTerminationsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCableTerminationsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/cable-terminations/][%d] dcim_cable-terminations_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/cable-terminations/][%d] dcim_cable-terminations_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}