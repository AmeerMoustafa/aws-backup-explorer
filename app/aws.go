package app

import (
	"aws-backup-explorer/utils"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func GetBackups() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(utils.EnvironmentVariables.AccessKey, utils.EnvironmentVariables.SecretKey, "")))

	if err != nil {
		log.Fatal("Could not authenticate, it is so over")
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = "us-east-1"

	})

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:    aws.String("barscloud.nightly.backups"),
		Prefix:    aws.String("DB21E/"),
		Delimiter: aws.String("/"),
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, prefix := range output.CommonPrefixes {
		folder_name := strings.TrimPrefix(*prefix.Prefix, "DB21E/")
		folder_name = strings.TrimSuffix(folder_name, "/")

		fmt.Printf("%s\n", folder_name)
	}

}
