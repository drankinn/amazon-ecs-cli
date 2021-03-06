// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package firehoseiface provides an interface for the Amazon Kinesis Firehose.
package firehoseiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/firehose"
)

// FirehoseAPI is the interface type for firehose.Firehose.
type FirehoseAPI interface {
	CreateDeliveryStreamRequest(*firehose.CreateDeliveryStreamInput) (*request.Request, *firehose.CreateDeliveryStreamOutput)

	CreateDeliveryStream(*firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error)

	DeleteDeliveryStreamRequest(*firehose.DeleteDeliveryStreamInput) (*request.Request, *firehose.DeleteDeliveryStreamOutput)

	DeleteDeliveryStream(*firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error)

	DescribeDeliveryStreamRequest(*firehose.DescribeDeliveryStreamInput) (*request.Request, *firehose.DescribeDeliveryStreamOutput)

	DescribeDeliveryStream(*firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error)

	ListDeliveryStreamsRequest(*firehose.ListDeliveryStreamsInput) (*request.Request, *firehose.ListDeliveryStreamsOutput)

	ListDeliveryStreams(*firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error)

	PutRecordRequest(*firehose.PutRecordInput) (*request.Request, *firehose.PutRecordOutput)

	PutRecord(*firehose.PutRecordInput) (*firehose.PutRecordOutput, error)

	PutRecordBatchRequest(*firehose.PutRecordBatchInput) (*request.Request, *firehose.PutRecordBatchOutput)

	PutRecordBatch(*firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error)

	UpdateDestinationRequest(*firehose.UpdateDestinationInput) (*request.Request, *firehose.UpdateDestinationOutput)

	UpdateDestination(*firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error)
}
