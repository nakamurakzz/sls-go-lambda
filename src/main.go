package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Event: S3
func handleS3Event(content context.Context, s3Event events.S3Event ) error {
	log.Println(content, s3Event)
	for _, record := range s3Event.Records {
		s3Entity := record.S3
		bucket := s3Entity.Bucket.Name
		key := s3Entity.Object.Key

		fileContent, err := readFileFromS3(bucket, key)
		if err != nil {
			return err
		}

		fmt.Printf("File content: %s\n", fileContent)

		log.Println("Sleeping for 5 seconds...")
		time.Sleep(10 * time.Second)
		log.Println("Done")
	}

	return nil
}

func readFileFromS3(bucket, key string) (string, error) {
	s := session.Must(session.NewSession())
	s3Client := s3.New(s)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := s3Client.GetObject(input)
	if err != nil {
		return "", err
	}

	defer result.Body.Close()

	content, err := io.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	lambda.Start(handleS3Event)
}
