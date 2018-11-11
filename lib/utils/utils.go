package utils

import (
  "time"
  "strconv"
	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

func EmptyRes(status int) (Response) {
  return Response{
    StatusCode: status,
  }
}

func ParseDate(date string) (string) {
  parsed, err := time.Parse("2-01-200615:04:05", date)
  if err != nil {
    return strconv.FormatInt(parsed.Unix(), 10)
  }
  return strconv.FormatInt(time.Now().Unix(), 10)
}
