// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListThingGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A filter that limits the results to those with the specified name prefix.
	NamePrefixFilter *string `location:"querystring" locationName:"namePrefixFilter" min:"1" type:"string"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// A filter that limits the results to those with the specified parent group.
	ParentGroup *string `location:"querystring" locationName:"parentGroup" min:"1" type:"string"`

	// If true, return child groups as well.
	Recursive *bool `location:"querystring" locationName:"recursive" type:"boolean"`
}

// String returns the string representation
func (s ListThingGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThingGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThingGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NamePrefixFilter != nil && len(*s.NamePrefixFilter) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefixFilter", 1))
	}
	if s.ParentGroup != nil && len(*s.ParentGroup) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ParentGroup", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingGroupsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NamePrefixFilter != nil {
		v := *s.NamePrefixFilter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "namePrefixFilter", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ParentGroup != nil {
		v := *s.ParentGroup

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "parentGroup", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Recursive != nil {
		v := *s.Recursive

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "recursive", protocol.BoolValue(v), metadata)
	}
	return nil
}

type ListThingGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The thing groups.
	ThingGroups []GroupNameAndArn `locationName:"thingGroups" type:"list"`
}

// String returns the string representation
func (s ListThingGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingGroupsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingGroups != nil {
		v := s.ThingGroups

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "thingGroups", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListThingGroups = "ListThingGroups"

// ListThingGroupsRequest returns a request value for making API operation for
// AWS IoT.
//
// List the thing groups in your account.
//
//    // Example sending a request using ListThingGroupsRequest.
//    req := client.ListThingGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListThingGroupsRequest(input *ListThingGroupsInput) ListThingGroupsRequest {
	op := &aws.Operation{
		Name:       opListThingGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/thing-groups",
	}

	if input == nil {
		input = &ListThingGroupsInput{}
	}

	req := c.newRequest(op, input, &ListThingGroupsOutput{})
	return ListThingGroupsRequest{Request: req, Input: input, Copy: c.ListThingGroupsRequest}
}

// ListThingGroupsRequest is the request type for the
// ListThingGroups API operation.
type ListThingGroupsRequest struct {
	*aws.Request
	Input *ListThingGroupsInput
	Copy  func(*ListThingGroupsInput) ListThingGroupsRequest
}

// Send marshals and sends the ListThingGroups API request.
func (r ListThingGroupsRequest) Send(ctx context.Context) (*ListThingGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThingGroupsResponse{
		ListThingGroupsOutput: r.Request.Data.(*ListThingGroupsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThingGroupsResponse is the response type for the
// ListThingGroups API operation.
type ListThingGroupsResponse struct {
	*ListThingGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThingGroups request.
func (r *ListThingGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}