package main

import (
	"aws-backup-explorer/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Could not load environment file")
	// }

	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")))

	// if err != nil {
	// 	log.Fatal("Could not authenticate, it is so over")
	// }

	// client := s3.NewFromConfig(cfg, func(o *s3.Options) {
	// 	o.Region = "us-east-1"

	// })

	// output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, b := range output.Buckets {
	// 	fmt.Printf(" %s created on %s \n", aws.ToString(b.Name), aws.ToTime(b.CreationDate))
	// }

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Could not load environment file: %s", err)
	}

	cmd.Execute()

	// layout := "20060102"

	// t, err := time.Parse(layout, "20250524")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t.Year())

}
