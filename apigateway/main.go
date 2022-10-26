package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(HandleRequest)
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

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request)
	log.Println("Body size d. \n", len(request.Body))
	log.Println(request.Body)
	return Response200("lo logramos ! ! !"), nil
}
