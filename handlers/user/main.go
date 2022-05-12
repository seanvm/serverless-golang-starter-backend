package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/seanvm/serverless-golang-starter-backend/app"
	"github.com/seanvm/serverless-golang-starter-backend/app/database"
	"github.com/seanvm/serverless-golang-starter-backend/app/factories"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var ResponseHeaders = make(map[string]string)
var db *database.Datastore
var sf factories.ServiceFactory

func init() {
	db = database.Db()

	sf = factories.ServiceFactory{
		Db: db,
	}
}

// userHandler is our lambda handler invoked by the `lambda.Start` function call
func userHandler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	ResponseHeaders["Access-Control-Allow-Origin"] = "*"
	statusCode := 200
	resp := Response{Body: "", StatusCode: statusCode, Headers: ResponseHeaders}

	var err error
	u := &app.User{}

	us := sf.BuildUserService()

	if request.HTTPMethod == "POST" {
		result, err := us.CreateUser(u)
		if err != nil {
			resp.StatusCode = 500
			resp.Body = err.Error()
			return resp, nil
		}

		bytes, _ := json.Marshal(result)
		respBody := string(bytes)
		resp.Body = respBody
		return resp, nil
	}

	if request.HTTPMethod == "GET" {
		id, ok := request.PathParameters["id"]
		if !ok {
			resp.StatusCode = 400
			resp.Body = errors.New("Missing required param: 'id'").Error()
			return resp, nil
		}

		u, err = us.GetUser(id)
		if err != nil {
			resp.StatusCode = 500
			resp.Body = err.Error()
			return resp, nil
		}

		bytes, _ := json.Marshal(u)
		respBody := string(bytes)
		resp.Body = respBody
		return resp, nil
	}

	resp.StatusCode = 400
	resp.Body = errors.New("Unhandled Request").Error()
	return resp, nil
}

func main() {
	lambda.Start(userHandler)
}
