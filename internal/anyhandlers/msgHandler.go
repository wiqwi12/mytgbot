package anyHandlers

import (
	"log/slog"
	db "telebotNew/internal/DB"
	"telebotNew/internal/tg"

	"gopkg.in/telebot.v3"
)

func StartHandler(bot *tg.Bot, c telebot.Context) error {

	exists, err := db.UserCheck(bot, c)
	if err != nil {
		slog.Error("usercheck error:", err)
	}
	if exists {
		return c.Send("Жду ссылку!")
	} else if !exists {
		if err := db.UserRegistration(bot, c); err != nil {
			slog.Error("user reg err:", err)
		}
		return c.Send("Ты отправляешь ссылки на статьи и решаешь, хочешь ли ты их прочитать в будущем или нет")
	}

	return nil
}

func LinkHandler(b *tg.Bot, c telebot.Context, link string) error {

	//msg := c.Callback().Message.Text

	article := ScrapHeader(link)

	btnOption1 := telebot.InlineButton{
		Unique: "SAVE",
		Text:   "Сохранить",
	}
	btnOption2 := telebot.InlineButton{

		Unique: "PASS",
		Text:   "Пропустить",
	}

	inlineKeyboard := [][]telebot.InlineButton{
		{btnOption1, btnOption2},
	}

	return c.Send(article, &telebot.ReplyMarkup{
		InlineKeyboard: inlineKeyboard,
	})
	return nil
}
