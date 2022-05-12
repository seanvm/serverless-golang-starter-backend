package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	mockRequest := events.APIGatewayProxyRequest{
		Path:           "user",
		HTTPMethod:     "GET",
		PathParameters: map[string]string{"id": "1"},
	}

	res, e := userHandler(ctx, mockRequest)
	if e != nil {
		t.Error("TestGetUser:" + e.Error())
	}

	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
