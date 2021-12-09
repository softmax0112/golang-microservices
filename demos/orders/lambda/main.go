package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Get the path parameter that was sent
	prod_num := request.PathParameters["productNumber"]

	//Generate message that want to be sent as body
	message := fmt.Sprintf(" { \"Message\" : \"%s - Reordered Successfully \" } ", prod_num)

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
