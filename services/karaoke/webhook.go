package karaoke

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	linewebhook "github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/ryo-kagawa/Vercel/domain"
	"github.com/ryo-kagawa/Vercel/services/karaoke/domain/model"
	"github.com/ryo-kagawa/Vercel/services/karaoke/infrastructure/database"
)

func (k Karaoke) Webhook(r *http.Request) (domain.HttpResponse, string, error) {
	environment, err := k.GetEnvironment()
	if err != nil {
		return domain.HttpResponse{}, "", err
	}
	if err := environment.Validate(); err != nil {
		return domain.HttpResponse{}, "", err
	}

	bot, err := messaging_api.NewMessagingApiAPI(
		environment.Line.LINE_CHANNEL_TOKEN,
	)
	if err != nil {
		return domain.HttpResponse{}, "", err
	}

	database, err := database.NewDatabase(environment.Database)
	if err != nil {
		return domain.HttpResponse{}, "", err
	}

	cb, err := linewebhook.ParseRequest(
		environment.Line.LINE_CHANNEL_SECRET,
		r,
	)
	if err != nil {
		return domain.HttpResponse{}, "", err
	}
	if len(cb.Events) == 0 {
		return domain.HttpResponse{}, "検証", nil
	}
	for _, event := range cb.Events {
		switch event := event.(type) {
		case linewebhook.MessageEvent:
			switch message := event.Message.(type) {
			case linewebhook.TextMessageContent:
				karaokeSongList, err := randomPickKaraokeSong(message.Text, database)
				if err != nil {
					return domain.HttpResponse{}, "", err
				}
				lineMessages := make([]messaging_api.MessageInterface, 0, len(karaokeSongList))
				for _, karaokeSong := range karaokeSongList {
					lineMessages = append(lineMessages, karaokeSong.GenerateLineTextMessage())
				}
				bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: event.ReplyToken,
						Messages:   lineMessages,
					},
				)
			}
		}
	}

	return domain.HttpResponse{
		Header: domain.HttpResponseHeader{
			HttpStatusCode: http.StatusOK,
			Contents:       []domain.HttpResponseHeaderContent{},
		},
		Body: "",
	}, "finish", nil
}

func randomPickKaraokeSong(text string, database database.Database) ([]model.KaraokeSong, error) {
	switch text {
	case "DAM":
		return database.Dam()
	case "JOYSOUND":
		return database.Joysound()
	default:
		return database.Ramdom()
	}
}
