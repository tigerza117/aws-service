package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// SQSSendMessageAPI defines the interface for the GetQueueUrl and SendMessage functions.
// We use this interface to test the functions using a mocked service.
type SQSSendMessageAPI interface {
	GetQueueUrl(ctx context.Context,
		params *sqs.GetQueueUrlInput,
		optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)

	SendMessage(ctx context.Context,
		params *sqs.SendMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)

	ReceiveMessage(ctx context.Context,
		params *sqs.ReceiveMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
}

// GetQueueURL gets the URL of an Amazon SQS queue.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a GetQueueUrlOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to GetQueueUrl.
func GetQueueURL(c context.Context, api SQSSendMessageAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(c, input)
}

// SendMsg sends a message to an Amazon SQS queue.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a SendMessageOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to SendMessage.
func SendMsg(c context.Context, api SQSSendMessageAPI, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return api.SendMessage(c, input)
}

// GetMessages gets the most recent message from an Amazon SQS queue.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a ReceiveMessageOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to ReceiveMessage.
func GetMessages(c context.Context, api SQSSendMessageAPI, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(c, input)
}

func main() {
	queue := flag.String("q", "", "The name of the queue")
	flag.Parse()

	if *queue == "" {
		*queue = "transaction"
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),

		config.WithRegion("us-east-1"),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	println("Hello World")

	client := sqs.NewFromConfig(cfg)

	// Get URL of queue
	gQInput := &sqs.GetQueueUrlInput{
		QueueName: queue,
	}

	result, err := GetQueueURL(context.TODO(), client, gQInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := result.QueueUrl

	for i := 0; i < 100; i++ {
		gMInput := &sqs.ReceiveMessageInput{
			MessageAttributeNames: []string{
				string(types.QueueAttributeNameAll),
			},
			QueueUrl:            queueURL,
			MaxNumberOfMessages: 1,
			VisibilityTimeout:   int32(100),
		}

		msgResult, err := GetMessages(context.TODO(), client, gMInput)
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			fmt.Println(err)
			return
		}

		if msgResult.Messages != nil {
			fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
			fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
		} else {
			fmt.Println("No messages found")
		}
	}

	//for i := 0; i < 100; i++ {
	//	sMInput := &sqs.SendMessageInput{
	//		MessageBody:  aws.String("Information about the NY Times fiction bestseller for the week of 12/11/2016."),
	//		QueueUrl:     queueURL,
	//		DelaySeconds: 1,
	//		MessageAttributes: map[string]types.MessageAttributeValue{
	//			"Title": {
	//				DataType:    aws.String("String"),
	//				StringValue: aws.String("The Whistler"),
	//			},
	//			"Author": {
	//				DataType:    aws.String("String"),
	//				StringValue: aws.String("John Grisham"),
	//			},
	//			"WeeksOn": {
	//				DataType:    aws.String("Number"),
	//				StringValue: aws.String("6"),
	//			},
	//		},
	//	}
	//
	//	resp, err := SendMsg(context.TODO(), client, sMInput)
	//	if err != nil {
	//		fmt.Println("Got an error sending the message:")
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Println("Sent message with ID: " + *resp.MessageId)
	//}
}
