package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Skarlso/miner/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Backup backs up a zipped world using AWS
func Backup(server string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	c := config.Config{}
	c.Unmarshal()
	bucket := c.Bucket
	filename := filepath.Join(config.Path(), archiveServer(server))
	file, err := os.Open(filename)

	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	log.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
