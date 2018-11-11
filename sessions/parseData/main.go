package main

import (
  "github.com/a-rmz/private-api/lib/utils"
  "github.com/a-rmz/private-api/lib/db/models"
  "fmt"
  "context"
  "strings"
  "strconv"

	"github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-lambda-go/events"
)


// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)

type Request string
type Response events.APIGatewayProxyResponse

// Incoming text format is as follows:
//   00000000;FFFFFFFF;2-10-201810:00:05;2-10-201810:00:05;10
func parseData(requestID string, request Request) models.SessionFrame {
  tokens := strings.Split(fmt.Sprintf("%s", request), ";") 
  userId := tokens[0]
  deviceId := tokens[1]
  startTime := utils.ParseDate(tokens[2])
  endTime := utils.ParseDate(tokens[3])
  data, err := strconv.ParseFloat(tokens[4], 32)
  if err != nil {
    data = 0
  }

  session, err := models.GetSessionByRFIDAndDevice(userId, deviceId)
  if err == nil {
    frame := models.SessionFrame{
      SessionFrameId: requestID,
      SessionID: session.SessionID,
      SessionType: "data",
      StartTime: startTime,
      EndTime: endTime,
      Data: float32(data),
    }
    return frame
  }
  panic(err)
}

// 
// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request Request) (Response, error) {
  lc, _ := lambdacontext.FromContext(ctx)
  frame := parseData(lc.AwsRequestID, request)
  
  models.InsertSessionFrame(frame)

	resp := Response(utils.EmptyRes(200))

  return resp, nil
}

func main() {
	lambda.Start(Handler)
}
