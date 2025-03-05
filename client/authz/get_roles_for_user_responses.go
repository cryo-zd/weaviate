//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package authz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/weaviate/weaviate/entities/models"
)

// GetRolesForUserReader is a Reader for the GetRolesForUser structure.
type GetRolesForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRolesForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRolesForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRolesForUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRolesForUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRolesForUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRolesForUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRolesForUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRolesForUserOK creates a GetRolesForUserOK with default headers values
func NewGetRolesForUserOK() *GetRolesForUserOK {
	return &GetRolesForUserOK{}
}

/*
GetRolesForUserOK describes a response with status code 200, with default header values.

Role assigned users
*/
type GetRolesForUserOK struct {
	Payload models.RolesListResponse
}

// IsSuccess returns true when this get roles for user o k response has a 2xx status code
func (o *GetRolesForUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get roles for user o k response has a 3xx status code
func (o *GetRolesForUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user o k response has a 4xx status code
func (o *GetRolesForUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles for user o k response has a 5xx status code
func (o *GetRolesForUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user o k response a status code equal to that given
func (o *GetRolesForUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get roles for user o k response
func (o *GetRolesForUserOK) Code() int {
	return 200
}

func (o *GetRolesForUserOK) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserOK  %+v", 200, o.Payload)
}

func (o *GetRolesForUserOK) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserOK  %+v", 200, o.Payload)
}

func (o *GetRolesForUserOK) GetPayload() models.RolesListResponse {
	return o.Payload
}

func (o *GetRolesForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserBadRequest creates a GetRolesForUserBadRequest with default headers values
func NewGetRolesForUserBadRequest() *GetRolesForUserBadRequest {
	return &GetRolesForUserBadRequest{}
}

/*
GetRolesForUserBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetRolesForUserBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user bad request response has a 2xx status code
func (o *GetRolesForUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user bad request response has a 3xx status code
func (o *GetRolesForUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user bad request response has a 4xx status code
func (o *GetRolesForUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user bad request response has a 5xx status code
func (o *GetRolesForUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user bad request response a status code equal to that given
func (o *GetRolesForUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get roles for user bad request response
func (o *GetRolesForUserBadRequest) Code() int {
	return 400
}

func (o *GetRolesForUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesForUserBadRequest) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserBadRequest  %+v", 400, o.Payload)
}

func (o *GetRolesForUserBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserUnauthorized creates a GetRolesForUserUnauthorized with default headers values
func NewGetRolesForUserUnauthorized() *GetRolesForUserUnauthorized {
	return &GetRolesForUserUnauthorized{}
}

/*
GetRolesForUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type GetRolesForUserUnauthorized struct{}

// IsSuccess returns true when this get roles for user unauthorized response has a 2xx status code
func (o *GetRolesForUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user unauthorized response has a 3xx status code
func (o *GetRolesForUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user unauthorized response has a 4xx status code
func (o *GetRolesForUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user unauthorized response has a 5xx status code
func (o *GetRolesForUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user unauthorized response a status code equal to that given
func (o *GetRolesForUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get roles for user unauthorized response
func (o *GetRolesForUserUnauthorized) Code() int {
	return 401
}

func (o *GetRolesForUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserUnauthorized ", 401)
}

func (o *GetRolesForUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserUnauthorized ", 401)
}

func (o *GetRolesForUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewGetRolesForUserForbidden creates a GetRolesForUserForbidden with default headers values
func NewGetRolesForUserForbidden() *GetRolesForUserForbidden {
	return &GetRolesForUserForbidden{}
}

/*
GetRolesForUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRolesForUserForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user forbidden response has a 2xx status code
func (o *GetRolesForUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user forbidden response has a 3xx status code
func (o *GetRolesForUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user forbidden response has a 4xx status code
func (o *GetRolesForUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user forbidden response has a 5xx status code
func (o *GetRolesForUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user forbidden response a status code equal to that given
func (o *GetRolesForUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get roles for user forbidden response
func (o *GetRolesForUserForbidden) Code() int {
	return 403
}

func (o *GetRolesForUserForbidden) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForUserForbidden) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserForbidden  %+v", 403, o.Payload)
}

func (o *GetRolesForUserForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRolesForUserNotFound creates a GetRolesForUserNotFound with default headers values
func NewGetRolesForUserNotFound() *GetRolesForUserNotFound {
	return &GetRolesForUserNotFound{}
}

/*
GetRolesForUserNotFound describes a response with status code 404, with default header values.

no role found for user
*/
type GetRolesForUserNotFound struct{}

// IsSuccess returns true when this get roles for user not found response has a 2xx status code
func (o *GetRolesForUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user not found response has a 3xx status code
func (o *GetRolesForUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user not found response has a 4xx status code
func (o *GetRolesForUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get roles for user not found response has a 5xx status code
func (o *GetRolesForUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get roles for user not found response a status code equal to that given
func (o *GetRolesForUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get roles for user not found response
func (o *GetRolesForUserNotFound) Code() int {
	return 404
}

func (o *GetRolesForUserNotFound) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserNotFound ", 404)
}

func (o *GetRolesForUserNotFound) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserNotFound ", 404)
}

func (o *GetRolesForUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewGetRolesForUserInternalServerError creates a GetRolesForUserInternalServerError with default headers values
func NewGetRolesForUserInternalServerError() *GetRolesForUserInternalServerError {
	return &GetRolesForUserInternalServerError{}
}

/*
GetRolesForUserInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type GetRolesForUserInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get roles for user internal server error response has a 2xx status code
func (o *GetRolesForUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get roles for user internal server error response has a 3xx status code
func (o *GetRolesForUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get roles for user internal server error response has a 4xx status code
func (o *GetRolesForUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get roles for user internal server error response has a 5xx status code
func (o *GetRolesForUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get roles for user internal server error response a status code equal to that given
func (o *GetRolesForUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get roles for user internal server error response
func (o *GetRolesForUserInternalServerError) Code() int {
	return 500
}

func (o *GetRolesForUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesForUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /authz/users/{id}/roles][%d] getRolesForUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRolesForUserInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRolesForUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetRolesForUserBody get roles for user body
swagger:model GetRolesForUserBody
*/
type GetRolesForUserBody struct {
	// user type
	// Required: true
	UserType *models.UserTypes `json:"user_type"`
}

// Validate validates this get roles for user body
func (o *GetRolesForUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUserType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRolesForUserBody) validateUserType(formats strfmt.Registry) error {
	if err := validate.Required("body"+"."+"user_type", "body", o.UserType); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"user_type", "body", o.UserType); err != nil {
		return err
	}

	if o.UserType != nil {
		if err := o.UserType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "user_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "user_type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get roles for user body based on the context it is used
func (o *GetRolesForUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUserType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRolesForUserBody) contextValidateUserType(ctx context.Context, formats strfmt.Registry) error {
	if o.UserType != nil {
		if err := o.UserType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "user_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "user_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRolesForUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRolesForUserBody) UnmarshalBinary(b []byte) error {
	var res GetRolesForUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
