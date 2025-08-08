package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SendMessageToSQS sends a message to the specified SQS queue URL.
func SendMessageToSQS(queueURL string, messageBody string) error {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("unable to load AWS config: %w", err)
	}

	client := sqs.NewFromConfig(cfg)
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(messageBody),
	}

	_, err = client.SendMessage(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
