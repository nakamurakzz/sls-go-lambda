package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleMain(ctx context.Context) error {
	fmt.Println("Hello World")
	return nil
}
func main() {
	lambda.Start(HandleMain)
}
