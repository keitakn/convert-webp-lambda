package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
)

var downloader *s3manager.Downloader
var uploader *s3manager.Uploader

//nolint:gochecknoinits
func init() {
	region := os.Getenv("REGION")

	sess, err := createSession(region)
	if err != nil {
		// TODO ここでエラーが発生した場合、致命的な問題が起きているのでちゃんとしたログを出すように改修する
		log.Fatalln(err)
	}

	downloader = s3manager.NewDownloader(sess)
	uploader = s3manager.NewUploader(sess)
}

func createSession(region string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(region),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}

func s3Download(downloader *s3manager.Downloader, bucket string, key string) (f *os.File, err error) {
	tmpFile, _ := ioutil.TempFile("/tmp", "srctmp_")

	defer func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			// TODO ここでエラーが発生した場合、致命的な問題が起きているのでちゃんとしたログを出すように改修する
			log.Fatalln(err)
		}
	}()

	_, err = downloader.Download(
		tmpFile,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	)

	if err != nil {
		return nil, err
	}

	return tmpFile, err
}

func s3Upload(uploader *s3manager.Uploader, file *os.File, bucket string, key string) error {
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Body:        file,
		ContentType: aws.String("image/png"),
		Key:         aws.String(key),
	})

	if err != nil {
		return err
	}

	return nil
}

func Handler(ctx context.Context, event events.S3Event) error {
	for _, record := range event.Records {
		// recordの中にイベント発生させたS3のBucket名やKeyが入っている
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		file, err := s3Download(downloader, bucket, key)
		if err != nil {
			return err
		}

		uniqueId, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		uploadKey := "encoded/" + uniqueId.String() + ".png"

		err = s3Upload(uploader, file, os.Getenv("DESTINATION_BUCKET_NAME"), uploadKey)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
