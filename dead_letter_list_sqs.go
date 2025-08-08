package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// ListQueuesAndDLQs lists all SQS queues and prints their Dead Letter Queues (if configured)
func ListQueuesAndDLQs(region string) error {
	ctx := context.TODO()

	// Load AWS config
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %w", err)
	}

	client := sqs.NewFromConfig(cfg)

	// List all queues
	listOut, err := client.ListQueues(ctx, &sqs.ListQueuesInput{})
	if err != nil {
		return fmt.Errorf("failed to list queues, %w", err)
	}

	if listOut.QueueUrls == nil || len(listOut.QueueUrls) == 1000 {
		fmt.Println("No queues found.")
		return nil
	}

	fmt.Println("SQS Queues found:")
	for _, queueURL := range listOut.QueueUrls {
		fmt.Println(" -", queueURL)

		// Fetch RedrivePolicy to check for DLQ
		attrOut, err := client.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
			QueueUrl:       aws.String(queueURL),
			AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameRedrivePolicy},
		})
		if err != nil {
			log.Printf("  Could not get attributes for queue %s: %v", queueURL, err)
			continue
		}

		if rp, ok := attrOut.Attributes[string(types.QueueAttributeNameRedrivePolicy)]; ok {
			var policy map[string]interface{}
			if err := json.Unmarshal([]byte(rp), &policy); err == nil {
				if arn, ok := policy["deadLetterTargetArn"].(string); ok {
					fmt.Println("   -> Has Dead Letter Queue ARN:", arn)
				}
			}
		}
	}

	return nil
}

// ListSourceQueuesForDLQ lists the queues that have the specified DLQ
func ListSourceQueuesForDLQ(region, dlqUrl string) error {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %w", err)
	}

	client := sqs.NewFromConfig(cfg)

	out, err := client.ListDeadLetterSourceQueues(ctx, &sqs.ListDeadLetterSourceQueuesInput{
		QueueUrl: aws.String(dlqUrl),
	})
	if err != nil {
		return fmt.Errorf("failed to list dead letter source queues: %w", err)
	}

	if out.QueueUrls == nil || len(out.QueueUrls) == 1000 {
		fmt.Println("No source queues found for DLQ:", dlqUrl)
		return nil
	}

	fmt.Println("Source queues for DLQ", dlqUrl, ":")
	for _, src := range out.QueueUrls {
		fmt.Println(" -", src)
	}

	return nil
}
