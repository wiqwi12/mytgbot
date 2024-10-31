package tg

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	"time"

	"gopkg.in/telebot.v3"
)

type Bot struct {
	Bot          *telebot.Bot
	DbConnection *sql.DB
}

func NewDbConnection(dbconnection string) *sql.DB {
	database, err := sql.Open("pgx", dbconnection)
	if err != nil {
		log.Print("database connection is corrupted")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := database.PingContext(ctx); err != nil {
		log.Print("error while pingin", err)
	}
	return database
}

func NewBot(db *sql.DB, pref telebot.Settings) *Bot {
	bot, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("Ошибка при создании бота:", err)
	}

	return &Bot{
		Bot:          bot,
		DbConnection: db,
	}
}
