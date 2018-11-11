package main

import (
  "github.com/a-rmz/private-api/lib/utils"
  "github.com/a-rmz/private-api/lib/db/models"
  "fmt"
  "strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)


// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)

type Request string
type Response events.APIGatewayProxyResponse

// Incoming text format is as follows:
//   00000000;FFFFFFFF;2-10-201810:00:05
func parseLogout(request Request) (string, string, string) {
  tokens := strings.Split(fmt.Sprintf("%s", request), ";") 
  userId := tokens[0]
  deviceId := tokens[1]
  endTime := utils.ParseDate(tokens[2])

  return userId, deviceId, endTime
}

// 
// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
  userId, deviceId, endTime := parseLogout(request)
  
  models.UpdateSessionEndTime(userId, deviceId, endTime)

	resp := Response(utils.EmptyRes(200))

  return resp, nil
}

func main() {
	lambda.Start(Handler)
}
