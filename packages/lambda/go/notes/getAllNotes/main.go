package main

import (
	"encoding/json"
	"sst-lambda-go-and-node/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetAllNotes() map[string]map[string]string {
	notes := db.Notes()
	return notes
}

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	notes := GetAllNotes()
	response, _ := json.Marshal(notes)

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
