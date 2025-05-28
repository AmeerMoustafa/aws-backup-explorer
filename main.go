package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment file")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")))

	if err != nil {
		log.Fatal("Could not authenticate, it is so over")
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "us-east-1"

	})

	output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		log.Fatal(err)
	}

	for _, b := range output.Buckets {
		fmt.Printf(" %s created on %s \n", aws.ToString(b.Name), aws.ToTime(b.CreationDate))
	}

}
