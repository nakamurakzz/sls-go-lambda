package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleMain() error {
	fmt.Println("Hello World!")
	return nil
}
func main() {
	lambda.Start(HandleMain)
}
