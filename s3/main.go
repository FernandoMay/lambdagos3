package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-lambda-go/lambda"
)

var s3session *s3.S3

const (
	REGION      = "us-east-2"
	BUCKET_NAME = "lambdagos3"
)

type InputEvent struct {
	Entidad  int    `json:"entidad"`
	Distrito int    `json:"distrito"`
	Username string `json:"username"`
	Paquete  string `json:"paquete"`
}

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func response(stringResponse string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{Body: stringResponse, StatusCode: statusCode}
}

func Response200(stringResponse string) events.APIGatewayProxyResponse {
	return response(stringResponse, 200)
}
func Response500(stringResponse string) events.APIGatewayProxyResponse {
	return response(stringResponse, 500)
}

func main() {
	lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var input InputEvent
	var err error
	// var image []byte
	err = GetInput([]byte(request.Body), &input)
	// if err != nil {
	// 	image := GetData(input.Username)
	// }
	if err == nil {
		err = PutS3(input)
	}
	if err == nil {
		return Response200("Done!"), nil
	}

	return Response500(err.Error()), nil
}

func GetInput(body []byte, data interface{}) error {
	return json.Unmarshal(body, &data)
}

func PutS3(data InputEvent) error {
	_, err := s3session.PutObject(&s3.PutObjectInput{
		Body:   json.Marshal({"distrito": data.Distrito,}),
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(data.Username),
	})
	return err
}

func GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
	return nil, err
}
