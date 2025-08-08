# Go SQS

# Purpose
This project provides simple Go tools and reusable functions to interact with Amazon Simple Queue Service (SQS):

* Send messages to an SQS queue.
* List all SQS queues in your AWS account/region.
* Display Dead Letter Queue (DLQ) configurations for each queue.
* List source queues that target a specified DLQ.

These tools are ideal for platform engineers or DevOps practitioners managing SQS resources and auditing DLQ usage.

# Prerequisites
* Go 1.16+ installed
* AWS SDK for Go v2 modules installed
* AWS credentials configured (via env vars, ~/.aws/credentials, or IAM roles)

# Usage
### 1. Sending a Message to SQS
Edit main.go and set your queue URL and the message you want to send:

```go
queueURL := "https://sqs.<region>.amazonaws.com/<account-id>/<queue-name>"
message := "Hello, this is an automated message from Go!"

err := SendMessageToSQS(queueURL, message)
if err != nil {
    log.Fatalf("Error sending message: %v", err)
}
```

Then run:

```bash
go run .
```

### 2. Listing All Queues and Their Dead Letter Queues (DLQs). 
* Edit the region as needed ("eu-west-1", "us-east-1", etc.):

```bash
region := "eu-west-1"
err := ListQueuesAndDLQs(region)
```
This will print each queue and the DLQ ARN if one is configured.

### 3. Listing Source Queues for a Given DLQ.
* If you know a DLQ queue URL, set:

```bash
dlqUrl := "https://sqs.eu-west-1.amazonaws.com/123456789012/my-dlq"
err := ListSourceQueuesForDLQ(region, dlqUrl)
```

The output will show which queues are configured to use your DLQ.

### 4. Customise in main.go.
* You may call any combination of these functions to suit your needs. If you leave dlqUrl empty, the DLQ source queue check will be skipped, run:
```bash
go run .
```
Typical output:

```bash
Message sent successfully!
SQS Queues found:
 - https://sqs.eu-west-1.amazonaws.com/446949660857/go_sqs
   -> Has Dead Letter Queue ARN: arn:aws:sqs:eu-west-1:446949660857:go_sqs_dlq
...
No DLQ URL provided for source queue check.
```

# Error Handling
If AWS configuration or credentials are missing, the program will print a descriptive error and exit.
If an SQS action fails (e.g., permission denied, queue does not exist), you’ll get a clear error with context.

# Recommendations
* Ensure your AWS credentials have necessary SQS permissions.
* Use environment variable AWS_REGION or configure region in code (config.WithRegion(region)).
* For production tools, consider accepting queue URL, message, region, etc., via CLI flags (with flag package).
