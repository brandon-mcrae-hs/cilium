// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for ResetNetworkInterfaceAttribute.
type ResetNetworkInterfaceAttributeInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`

	// The source/destination checking attribute. Resets the value to true.
	SourceDestCheck *string `locationName:"sourceDestCheck" type:"string"`
}

// String returns the string representation
func (s ResetNetworkInterfaceAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetNetworkInterfaceAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetNetworkInterfaceAttributeInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetNetworkInterfaceAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResetNetworkInterfaceAttributeOutput) String() string {
	return awsutil.Prettify(s)
}

const opResetNetworkInterfaceAttribute = "ResetNetworkInterfaceAttribute"

// ResetNetworkInterfaceAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Resets a network interface attribute. You can specify only one attribute
// at a time.
//
//    // Example sending a request using ResetNetworkInterfaceAttributeRequest.
//    req := client.ResetNetworkInterfaceAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ResetNetworkInterfaceAttribute
func (c *Client) ResetNetworkInterfaceAttributeRequest(input *ResetNetworkInterfaceAttributeInput) ResetNetworkInterfaceAttributeRequest {
	op := &aws.Operation{
		Name:       opResetNetworkInterfaceAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetNetworkInterfaceAttributeInput{}
	}

	req := c.newRequest(op, input, &ResetNetworkInterfaceAttributeOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ResetNetworkInterfaceAttributeRequest{Request: req, Input: input, Copy: c.ResetNetworkInterfaceAttributeRequest}
}

// ResetNetworkInterfaceAttributeRequest is the request type for the
// ResetNetworkInterfaceAttribute API operation.
type ResetNetworkInterfaceAttributeRequest struct {
	*aws.Request
	Input *ResetNetworkInterfaceAttributeInput
	Copy  func(*ResetNetworkInterfaceAttributeInput) ResetNetworkInterfaceAttributeRequest
}

// Send marshals and sends the ResetNetworkInterfaceAttribute API request.
func (r ResetNetworkInterfaceAttributeRequest) Send(ctx context.Context) (*ResetNetworkInterfaceAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetNetworkInterfaceAttributeResponse{
		ResetNetworkInterfaceAttributeOutput: r.Request.Data.(*ResetNetworkInterfaceAttributeOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetNetworkInterfaceAttributeResponse is the response type for the
// ResetNetworkInterfaceAttribute API operation.
type ResetNetworkInterfaceAttributeResponse struct {
	*ResetNetworkInterfaceAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetNetworkInterfaceAttribute request.
func (r *ResetNetworkInterfaceAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
