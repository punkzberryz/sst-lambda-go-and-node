package main

import (
	"encoding/json"
	"sst-lambda-go-and-node/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetNoteById(id string) map[string]string {
	notes := db.Notes()
	note := notes[id]
	if note == nil {
		return nil
	}
	return note
}

func handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	note := GetNoteById(request.PathParameters["id"])

	if note == nil {
		return events.APIGatewayProxyResponse{
			Body:       `{"error":true}`,
			StatusCode: 404,
		}, nil
	}

	response, _ := json.Marshal(note)

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
