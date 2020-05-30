package main

import (
	"context"
	"fmt"
	"github.com/apex/gateway"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
)

func responseStr(w http.ResponseWriter, code int, key string, data string) string {
	res := fmt.Sprintf(`{"data":{"%v": %v}}`, key, data)
	w.WriteHeader(code)
	fmt.Fprintln(w, res)
	return res
}

func Route1(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	responseStr(w, 200, "res", "Response 1")
}

func Route2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	responseStr(w, 201, "res", "Response 2")
}

func handler(lambdaRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("got request", lambdaRequest)
	request, err := gateway.NewRequest(context.Background(), lambdaRequest)
	if err != nil {
		fmt.Println("newRequest error", err)
	}

	router := httprouter.New()
	router.GET("/api/v1/res1", Route1)
	router.GET("/api/v1/res2", Route2)

	res := httptest.NewRecorder()
	router.ServeHTTP(res, request)

	return events.APIGatewayProxyResponse{
		Body:       res.Body.String(),
		StatusCode: res.Code,
	}, nil
}

func main() {
	lambda.Start(handler)
}
