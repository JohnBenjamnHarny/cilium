// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisassociateVpcCidrBlockCommon = "DisassociateVpcCidrBlock"

// DisassociateVpcCidrBlockCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateVpcCidrBlockCommon operation. The "output" return
// value will be populated with the DisassociateVpcCidrBlockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateVpcCidrBlockCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateVpcCidrBlockCommon Send returns without error.
//
// See DisassociateVpcCidrBlockCommon for more information on using the DisassociateVpcCidrBlockCommon
// API call, and error handling.
//
//    // Example sending a request using the DisassociateVpcCidrBlockCommonRequest method.
//    req, resp := client.DisassociateVpcCidrBlockCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DisassociateVpcCidrBlockCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisassociateVpcCidrBlockCommon,
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

// DisassociateVpcCidrBlockCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateVpcCidrBlockCommon for usage and error information.
func (c *VPC) DisassociateVpcCidrBlockCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisassociateVpcCidrBlockCommonRequest(input)
	return out, req.Send()
}

// DisassociateVpcCidrBlockCommonWithContext is the same as DisassociateVpcCidrBlockCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateVpcCidrBlockCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateVpcCidrBlockCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisassociateVpcCidrBlockCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisassociateVpcCidrBlock = "DisassociateVpcCidrBlock"

// DisassociateVpcCidrBlockRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateVpcCidrBlock operation. The "output" return
// value will be populated with the DisassociateVpcCidrBlockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateVpcCidrBlockCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateVpcCidrBlockCommon Send returns without error.
//
// See DisassociateVpcCidrBlock for more information on using the DisassociateVpcCidrBlock
// API call, and error handling.
//
//    // Example sending a request using the DisassociateVpcCidrBlockRequest method.
//    req, resp := client.DisassociateVpcCidrBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DisassociateVpcCidrBlockRequest(input *DisassociateVpcCidrBlockInput) (req *request.Request, output *DisassociateVpcCidrBlockOutput) {
	op := &request.Operation{
		Name:       opDisassociateVpcCidrBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateVpcCidrBlockInput{}
	}

	output = &DisassociateVpcCidrBlockOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisassociateVpcCidrBlock API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateVpcCidrBlock for usage and error information.
func (c *VPC) DisassociateVpcCidrBlock(input *DisassociateVpcCidrBlockInput) (*DisassociateVpcCidrBlockOutput, error) {
	req, out := c.DisassociateVpcCidrBlockRequest(input)
	return out, req.Send()
}

// DisassociateVpcCidrBlockWithContext is the same as DisassociateVpcCidrBlock with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateVpcCidrBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateVpcCidrBlockWithContext(ctx volcengine.Context, input *DisassociateVpcCidrBlockInput, opts ...request.Option) (*DisassociateVpcCidrBlockOutput, error) {
	req, out := c.DisassociateVpcCidrBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisassociateVpcCidrBlockInput struct {
	_ struct{} `type:"structure"`

	SecondaryCidrBlock *string `type:"string"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateVpcCidrBlockInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateVpcCidrBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateVpcCidrBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DisassociateVpcCidrBlockInput"}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSecondaryCidrBlock sets the SecondaryCidrBlock field's value.
func (s *DisassociateVpcCidrBlockInput) SetSecondaryCidrBlock(v string) *DisassociateVpcCidrBlockInput {
	s.SecondaryCidrBlock = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DisassociateVpcCidrBlockInput) SetVpcId(v string) *DisassociateVpcCidrBlockInput {
	s.VpcId = &v
	return s
}

type DisassociateVpcCidrBlockOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DisassociateVpcCidrBlockOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateVpcCidrBlockOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DisassociateVpcCidrBlockOutput) SetRequestId(v string) *DisassociateVpcCidrBlockOutput {
	s.RequestId = &v
	return s
}
