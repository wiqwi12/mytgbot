package main

import (
	"log/slog"
	"os"

	anyHandlers "telebotNew/internal/anyhandlers"

	"telebotNew/internal/interface/middleware"
	"telebotNew/internal/tg"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {

	if err := godotenv.Load(); err != nil {
		slog.Error("ENV LOAD ERR", err)
	}

	pref := telebot.Settings{
		Token: os.Getenv("BOT_TOKEN"),
	}

	db := tg.NewDbConnection(os.Getenv("DB_CONNECTION"))

	bot := tg.NewBot(db, pref)

	bot.Bot.Handle(telebot.OnText, func(c telebot.Context) error {
		return middleware.MsgHandler(bot, c)
	})
	bot.Bot.Handle(telebot.OnCallback, func(c telebot.Context) error {
		return anyHandlers.CallbackHanlder(bot, c)
	})

	bot.Bot.Start()

}
