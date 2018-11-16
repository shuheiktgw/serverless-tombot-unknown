package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/shuheiktgw/go-lambda-linebot/parser"
)

func handler(request events.SNSEvent) error {
	fmt.Printf("tombot unknown started, event: %s\n", request)

	event, err := parser.ParseSNSEvent(&request)
	if err != nil {
		return err
	}

	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		fmt.Printf("error occurred while initializing LINE client: %s\n", err)
		return err
	}

	message := linebot.NewTextMessage(`登録されていないコマンドです。利用可能なコマンドの一覧は"tmb help"でご確認ください!`)
	_, err = bot.ReplyMessage(event.ReplyToken, message).Do()
	if err != nil {
		fmt.Printf("error occurred while sending message: %s\n", err)
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
