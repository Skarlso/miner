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
	"github.com/fatih/color"
)

// Backup backs up a zipped world using AWS
func Backup(server string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	c := config.Config{}
	c.Unmarshal()
	bucket := c.Bucket
	filename := archiveServer(server)
	file, err := os.Open(filename)

	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()
	uploader := s3manager.NewUploader(sess)
	cyan := color.New(color.FgCyan).SprintFunc()
	filebase := filepath.Base(filename)
	log.Println("Beginning to upload: ", cyan(filebase))
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filebase),
		Body:   file,
	})

	if err != nil {
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("%s uploaded %q to %q\n", green("Successfully"), cyan(filename), cyan(bucket))
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
