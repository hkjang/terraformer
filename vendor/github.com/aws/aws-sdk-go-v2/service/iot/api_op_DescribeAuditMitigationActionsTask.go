// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeAuditMitigationActionsTaskInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the audit mitigation task.
	//
	// TaskId is a required field
	TaskId *string `location:"uri" locationName:"taskId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAuditMitigationActionsTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAuditMitigationActionsTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAuditMitigationActionsTaskInput"}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}
	if s.TaskId != nil && len(*s.TaskId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAuditMitigationActionsTaskInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "taskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeAuditMitigationActionsTaskOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the mitigation actions and their parameters that are applied as
	// part of this task.
	ActionsDefinition []MitigationAction `locationName:"actionsDefinition" type:"list"`

	// Specifies the mitigation actions that should be applied to specific audit
	// checks.
	AuditCheckToActionsMapping map[string][]string `locationName:"auditCheckToActionsMapping" type:"map"`

	// The date and time when the task was completed or canceled.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The date and time when the task was started.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`

	// Identifies the findings to which the mitigation actions are applied. This
	// can be by audit checks, by audit task, or a set of findings.
	Target *AuditMitigationActionsTaskTarget `locationName:"target" type:"structure"`

	// Aggregate counts of the results when the mitigation tasks were applied to
	// the findings for this audit mitigation actions task.
	TaskStatistics map[string]TaskStatisticsForAuditCheck `locationName:"taskStatistics" type:"map"`

	// The current status of the task.
	TaskStatus AuditMitigationActionsTaskStatus `locationName:"taskStatus" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeAuditMitigationActionsTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAuditMitigationActionsTaskOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ActionsDefinition != nil {
		v := s.ActionsDefinition

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "actionsDefinition", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AuditCheckToActionsMapping != nil {
		v := s.AuditCheckToActionsMapping

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "auditCheckToActionsMapping", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v2)})
			}
			ls1.End()
		}
		ms0.End()

	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Target != nil {
		v := s.Target

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "target", v, metadata)
	}
	if s.TaskStatistics != nil {
		v := s.TaskStatistics

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "taskStatistics", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if len(s.TaskStatus) > 0 {
		v := s.TaskStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "taskStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opDescribeAuditMitigationActionsTask = "DescribeAuditMitigationActionsTask"

// DescribeAuditMitigationActionsTaskRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about an audit mitigation task that is used to apply mitigation
// actions to a set of audit findings. Properties include the actions being
// applied, the audit checks to which they're being applied, the task status,
// and aggregated task statistics.
//
//    // Example sending a request using DescribeAuditMitigationActionsTaskRequest.
//    req := client.DescribeAuditMitigationActionsTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeAuditMitigationActionsTaskRequest(input *DescribeAuditMitigationActionsTaskInput) DescribeAuditMitigationActionsTaskRequest {
	op := &aws.Operation{
		Name:       opDescribeAuditMitigationActionsTask,
		HTTPMethod: "GET",
		HTTPPath:   "/audit/mitigationactions/tasks/{taskId}",
	}

	if input == nil {
		input = &DescribeAuditMitigationActionsTaskInput{}
	}

	req := c.newRequest(op, input, &DescribeAuditMitigationActionsTaskOutput{})
	return DescribeAuditMitigationActionsTaskRequest{Request: req, Input: input, Copy: c.DescribeAuditMitigationActionsTaskRequest}
}

// DescribeAuditMitigationActionsTaskRequest is the request type for the
// DescribeAuditMitigationActionsTask API operation.
type DescribeAuditMitigationActionsTaskRequest struct {
	*aws.Request
	Input *DescribeAuditMitigationActionsTaskInput
	Copy  func(*DescribeAuditMitigationActionsTaskInput) DescribeAuditMitigationActionsTaskRequest
}

// Send marshals and sends the DescribeAuditMitigationActionsTask API request.
func (r DescribeAuditMitigationActionsTaskRequest) Send(ctx context.Context) (*DescribeAuditMitigationActionsTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAuditMitigationActionsTaskResponse{
		DescribeAuditMitigationActionsTaskOutput: r.Request.Data.(*DescribeAuditMitigationActionsTaskOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAuditMitigationActionsTaskResponse is the response type for the
// DescribeAuditMitigationActionsTask API operation.
type DescribeAuditMitigationActionsTaskResponse struct {
	*DescribeAuditMitigationActionsTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAuditMitigationActionsTask request.
func (r *DescribeAuditMitigationActionsTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}