package anyHandlers

import (
	"log/slog"
	db "telebotNew/internal/DB"
	"telebotNew/internal/tg"

	"gopkg.in/telebot.v3"
)

func CallbackHanlder(b *tg.Bot, c telebot.Context) error {

	article := ScrapHeader(c.Callback().Message.Text)
	data := c.Callback().Data
	slog.Info("Button pushed:", data)

	switch data {
	case "SAVE":

		err := db.SaveArticle(b, c, article)
		if err != nil {
			return c.Send("Ошибка при сохранении статьи. Пожалуйста, попробуйте снова.")
		}
		return c.Send("Статья сохранена!")
	case "PASS":
		return c.Send("Вы пропустили статью.")
	default:
		return c.Send("Неизвестная команда.")
	}
}
