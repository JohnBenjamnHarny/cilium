// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifySecurityGroupAttributesCommon = "ModifySecurityGroupAttributes"

// ModifySecurityGroupAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifySecurityGroupAttributesCommon operation. The "output" return
// value will be populated with the ModifySecurityGroupAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySecurityGroupAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySecurityGroupAttributesCommon Send returns without error.
//
// See ModifySecurityGroupAttributesCommon for more information on using the ModifySecurityGroupAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifySecurityGroupAttributesCommonRequest method.
//    req, resp := client.ModifySecurityGroupAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifySecurityGroupAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySecurityGroupAttributesCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySecurityGroupAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifySecurityGroupAttributesCommon for usage and error information.
func (c *VPC) ModifySecurityGroupAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupAttributesCommonWithContext is the same as ModifySecurityGroupAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroupAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifySecurityGroupAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySecurityGroupAttributes = "ModifySecurityGroupAttributes"

// ModifySecurityGroupAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifySecurityGroupAttributes operation. The "output" return
// value will be populated with the ModifySecurityGroupAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySecurityGroupAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySecurityGroupAttributesCommon Send returns without error.
//
// See ModifySecurityGroupAttributes for more information on using the ModifySecurityGroupAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifySecurityGroupAttributesRequest method.
//    req, resp := client.ModifySecurityGroupAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifySecurityGroupAttributesRequest(input *ModifySecurityGroupAttributesInput) (req *request.Request, output *ModifySecurityGroupAttributesOutput) {
	op := &request.Operation{
		Name:       opModifySecurityGroupAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySecurityGroupAttributesInput{}
	}

	output = &ModifySecurityGroupAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySecurityGroupAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifySecurityGroupAttributes for usage and error information.
func (c *VPC) ModifySecurityGroupAttributes(input *ModifySecurityGroupAttributesInput) (*ModifySecurityGroupAttributesOutput, error) {
	req, out := c.ModifySecurityGroupAttributesRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupAttributesWithContext is the same as ModifySecurityGroupAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroupAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifySecurityGroupAttributesWithContext(ctx volcengine.Context, input *ModifySecurityGroupAttributesInput, opts ...request.Option) (*ModifySecurityGroupAttributesOutput, error) {
	req, out := c.ModifySecurityGroupAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifySecurityGroupAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// SecurityGroupId is a required field
	SecurityGroupId *string `type:"string" required:"true"`

	SecurityGroupName *string `type:"string"`
}

// String returns the string representation
func (s ModifySecurityGroupAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySecurityGroupAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySecurityGroupAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifySecurityGroupAttributesInput"}
	if s.SecurityGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifySecurityGroupAttributesInput) SetDescription(v string) *ModifySecurityGroupAttributesInput {
	s.Description = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *ModifySecurityGroupAttributesInput) SetSecurityGroupId(v string) *ModifySecurityGroupAttributesInput {
	s.SecurityGroupId = &v
	return s
}

// SetSecurityGroupName sets the SecurityGroupName field's value.
func (s *ModifySecurityGroupAttributesInput) SetSecurityGroupName(v string) *ModifySecurityGroupAttributesInput {
	s.SecurityGroupName = &v
	return s
}

type ModifySecurityGroupAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifySecurityGroupAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySecurityGroupAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifySecurityGroupAttributesOutput) SetRequestId(v string) *ModifySecurityGroupAttributesOutput {
	s.RequestId = &v
	return s
}
