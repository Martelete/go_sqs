package main

import (
	"fmt"
	"log"
)

func main() {
	region := "" // Set your AWS region

	// Edit this for your queue details
	queueURL := "" // Add your SQS URL
	message := ""  // Add your message to send to the queue

	err := SendMessageToSQS(queueURL, message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Println("Message sent successfully!")

	// List all queues and their DLQs
	if err := ListQueuesAndDLQs(region); err != nil {
		log.Fatalf("Error listing queues and DLQs: %v", err)
	}

	// Example: Check source queues for a particular DLQ
	// Replace with an actual DLQ URL if needed
	dlqUrl := "" // e.g.: "https://sqs.eu-west-1.amazonaws.com/123456789012/my-dlq"
	if dlqUrl != "" {
		if err := ListSourceQueuesForDLQ(region, dlqUrl); err != nil {
			log.Fatalf("Error listing source queues for DLQ: %v", err)
		}
	} else {
		fmt.Println("\nNo DLQ URL provided for source queue check.")
	}
}
