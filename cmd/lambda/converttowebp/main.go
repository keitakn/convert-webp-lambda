package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var svc *s3.S3

//nolint:gochecknoinits
func init() {
	sess, err := session.NewSession()
	if err != nil {
		// TODO ここでエラーが発生した場合、致命的な問題が起きているのでちゃんとしたログを出すように改修する
		log.Fatalln(err)
	}

	svc = s3.New(sess)

	log.Println("⊂ﾟＵ┬───┬~")
	log.Println(svc)
	log.Println("⊂ﾟＵ┬───┬~")
}

func Handler(ctx context.Context, event events.S3Event) error {
	for _, record := range event.Records {
		// recordの中にイベント発生させたS3のBucket名やKeyが入っている
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		log.Println("🐱")
		log.Println(ctx)
		log.Println(bucket)
		log.Println(key)
		log.Println("🐱")
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
