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
)

// DcimRacksDeleteReader is a Reader for the DcimRacksDelete structure.
type DcimRacksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRacksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRacksDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksDeleteNoContent creates a DcimRacksDeleteNoContent with default headers values
func NewDcimRacksDeleteNoContent() *DcimRacksDeleteNoContent {
	return &DcimRacksDeleteNoContent{}
}

/*
DcimRacksDeleteNoContent describes a response with status code 204, with default header values.

DcimRacksDeleteNoContent dcim racks delete no content
*/
type DcimRacksDeleteNoContent struct {
}

// IsSuccess returns true when this dcim racks delete no content response has a 2xx status code
func (o *DcimRacksDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks delete no content response has a 3xx status code
func (o *DcimRacksDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks delete no content response has a 4xx status code
func (o *DcimRacksDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks delete no content response has a 5xx status code
func (o *DcimRacksDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks delete no content response a status code equal to that given
func (o *DcimRacksDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimRacksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteNoContent ", 204)
}

func (o *DcimRacksDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteNoContent ", 204)
}

func (o *DcimRacksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRacksDeleteDefault creates a DcimRacksDeleteDefault with default headers values
func NewDcimRacksDeleteDefault(code int) *DcimRacksDeleteDefault {
	return &DcimRacksDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRacksDeleteDefault describes a response with status code -1, with default header values.

DcimRacksDeleteDefault dcim racks delete default
*/
type DcimRacksDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks delete default response
func (o *DcimRacksDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim racks delete default response has a 2xx status code
func (o *DcimRacksDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks delete default response has a 3xx status code
func (o *DcimRacksDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks delete default response has a 4xx status code
func (o *DcimRacksDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks delete default response has a 5xx status code
func (o *DcimRacksDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks delete default response a status code equal to that given
func (o *DcimRacksDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRacksDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcim_racks_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcim_racks_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
