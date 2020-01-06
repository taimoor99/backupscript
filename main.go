package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"time"
)

func main() {
	file, err := os.Open("/opt/backup/backup.tar")
	if err != nil {
		fmt.Println("Failed to open file", file, err)
		os.Exit(1)
	}
	defer file.Close()
	endpoint := "sfo2.digitaloceanspaces.com"
	region := "sfo2"
	creds := credentials.NewStaticCredentials("AwsAccessKey", "AwsSecretAccessKey", "")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
		Credentials: creds,
	}))
	uploader := s3manager.NewUploader(sess)

	filename := time.Now().Format("20060102150405")

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("bucket directory to upload"),
		Key:    aws.String(filename),
		Body:   file,
		ACL: aws.String("public-read"),
	})
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

