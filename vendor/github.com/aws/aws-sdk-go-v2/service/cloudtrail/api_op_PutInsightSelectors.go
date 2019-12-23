// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutInsightSelectorsInput struct {
	_ struct{} `type:"structure"`

	// A JSON string that contains the insight types you want to log on a trail.
	// In this release, only ApiCallRateInsight is supported as an insight type.
	//
	// InsightSelectors is a required field
	InsightSelectors []InsightSelector `type:"list" required:"true"`

	// The name of the CloudTrail trail for which you want to change or add Insights
	// selectors.
	//
	// TrailName is a required field
	TrailName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutInsightSelectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutInsightSelectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutInsightSelectorsInput"}

	if s.InsightSelectors == nil {
		invalidParams.Add(aws.NewErrParamRequired("InsightSelectors"))
	}

	if s.TrailName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrailName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutInsightSelectorsOutput struct {
	_ struct{} `type:"structure"`

	// A JSON string that contains the insight types you want to log on a trail.
	// In this release, only ApiCallRateInsight is supported as an insight type.
	InsightSelectors []InsightSelector `type:"list"`

	// The Amazon Resource Name (ARN) of a trail for which you want to change or
	// add Insights selectors.
	TrailARN *string `type:"string"`
}

// String returns the string representation
func (s PutInsightSelectorsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutInsightSelectors = "PutInsightSelectors"

// PutInsightSelectorsRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Lets you enable Insights event logging by specifying the Insights selectors
// that you want to enable on an existing trail. You also use PutInsightSelectors
// to turn off Insights event logging, by passing an empty list of insight types.
// In this release, only ApiCallRateInsight is supported as an Insights selector.
//
//    // Example sending a request using PutInsightSelectorsRequest.
//    req := client.PutInsightSelectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/PutInsightSelectors
func (c *Client) PutInsightSelectorsRequest(input *PutInsightSelectorsInput) PutInsightSelectorsRequest {
	op := &aws.Operation{
		Name:       opPutInsightSelectors,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutInsightSelectorsInput{}
	}

	req := c.newRequest(op, input, &PutInsightSelectorsOutput{})
	return PutInsightSelectorsRequest{Request: req, Input: input, Copy: c.PutInsightSelectorsRequest}
}

// PutInsightSelectorsRequest is the request type for the
// PutInsightSelectors API operation.
type PutInsightSelectorsRequest struct {
	*aws.Request
	Input *PutInsightSelectorsInput
	Copy  func(*PutInsightSelectorsInput) PutInsightSelectorsRequest
}

// Send marshals and sends the PutInsightSelectors API request.
func (r PutInsightSelectorsRequest) Send(ctx context.Context) (*PutInsightSelectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutInsightSelectorsResponse{
		PutInsightSelectorsOutput: r.Request.Data.(*PutInsightSelectorsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutInsightSelectorsResponse is the response type for the
// PutInsightSelectors API operation.
type PutInsightSelectorsResponse struct {
	*PutInsightSelectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutInsightSelectors request.
func (r *PutInsightSelectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}