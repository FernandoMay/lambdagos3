package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Handler(request map[string]interface{}) (string, error) {
	if len(request) != 240 {
		log.Println("wrong number of items")
		return fmt.Sprintf("Unexpected length of arguments: %d", len(request)), errors.New("Error: wrong number of args")
	}

	jsonBod, err := json.Marshal(request)
	if err != nil {
		log.Println(err.Error())
		return "can't create json response", errors.New("Error: can't create json response")
	}

	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	u := request["username"]
	key := fmt.Sprintf("responses/%s", u)
	_, ierr := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("lambdagos3"),
		Key:    aws.String(key),
		Body:   bytes.NewReader(jsonBod),
	})

	if ierr != nil {
		log.Printf("There was an issue uploading to s3: %s", ierr.Error())
		return "Unable to save response", errors.New("Error: cant save response")
	}

	return string(jsonBod), nil
}

func main() {
	lambda.Start(Handler)
}
