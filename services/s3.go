package services

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Manager struct {
}

func (s *S3Manager) UploadFile(bucket string, filename string) {

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key_id"), os.Getenv("aws_secret_access_key"), ""),
	})

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		ContentType: aws.String("text/html"),
		Body:        file,
		ACL:         aws.String("public-read"),
	})
	if err != nil {
		// Print the error and exit.
		log.Printf("Unable to upload %q to %q, %v", filename, bucket, err)
		return
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
