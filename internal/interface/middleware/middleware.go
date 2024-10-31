package middleware

import (
	anyHandlers "telebotNew/internal/anyhandlers"
	"telebotNew/internal/tg"

	"gopkg.in/telebot.v3"
)

func MsgHandler(b *tg.Bot, c telebot.Context) error {
	msgText := c.Text()

	if msgText == "/start" {
		anyHandlers.StartHandler(b, c)
	} else if anyHandlers.IsLink(msgText) {
		anyHandlers.LinkHandler(b, c, msgText)
	}

	return nil

}
