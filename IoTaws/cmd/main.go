package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	awsConfig := aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
		Endpoint:    aws.String(os.Getenv("IOT_ENDPOINT")),
	}

	sess, err := session.NewSession()

	if err != nil {
		fmt.Println("session error", err.Error())
	}

	svc := iotdataplane.New(sess, &awsConfig)

	params := &iotdataplane.PublishInput{
		Topic:   aws.String(os.Getenv("IOT_TOPICS")), // Required
		Payload: []byte("hello"),
		Qos:     aws.Int64(1),
	}

	if rsp, err := svc.Publish(params); err != nil {
		fmt.Println("iot error", err.Error())
	} else {
		fmt.Println(rsp)
	}
}
