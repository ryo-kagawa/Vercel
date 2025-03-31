package model

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

type KaraokeSong struct {
	ArtistName string
	SongName   string
	Lyrics     string
	DamId      string
	JoysoundId string
}

func (k KaraokeSong) GenerateLineTextMessage() messaging_api.TextMessage {
	text := ""
	text += fmt.Sprintf("アーティスト: %s\n", k.ArtistName)
	text += fmt.Sprintf("曲名: %s\n", k.SongName)
	if k.DamId != "" {
		text += fmt.Sprintf("DAM選曲番号: %s\n", k.DamId)
	}
	if k.JoysoundId != "" {
		text += fmt.Sprintf("JOYSOUND曲番号: %s\n", k.JoysoundId)
	}
	text += fmt.Sprintf("歌詞\n%s", k.Lyrics)

	return messaging_api.TextMessage{
		Text: text,
	}
}
