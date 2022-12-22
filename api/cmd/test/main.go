package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s %s \n", sqsEvent.Records[0].MessageId, message.EventSource, message.Body, message.MessageAttributes)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
