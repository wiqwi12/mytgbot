package db

import (
	"fmt"
	"log/slog"
	"strconv"

	"telebotNew/internal/tg"

	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gopkg.in/telebot.v3"
)

func UserCheck(b *tg.Bot, c telebot.Context) (bool, error) {

	id := c.Sender().ID

	where := "chat_id = " + strconv.Itoa(int(id))

	query := squirrel.Select("COUNT(*)").From("articles").Where(where)

	finQuery, _, err := query.ToSql()
	if err != nil {
		slog.Error("TOSQL ERR IN USERCHECK")
		return false, err
	}

	var count int

	err = b.DbConnection.QueryRow(finQuery).Scan(&count)
	if err != nil {
		slog.Error("QUERY ERROR", err)
		return false, err
	}

	return count > 0, nil

}

func UserRegistration(b *tg.Bot, c telebot.Context) error {

	query := "INSERT INTO articles (chat_id) VALUES " + fmt.Sprintf("(%d)", c.Sender().ID)

	_, err := b.DbConnection.Exec(query)
	if err != nil {
		slog.Error("EXEC ERROR IN REGISTRATION", err)
		return err
	}

	return nil

}

func SaveArticle(b *tg.Bot, c telebot.Context, article string) error {

	id := c.Sender().ID
	link := c.Message().Text

	query := "INSERT INTO articles (chat_id, link, title) " + fmt.Sprintf("VALUES (%d, %s, %s)", id, link, article)

	_, err := b.DbConnection.Exec(query)
	if err != nil {
		slog.Error("SAVE ARTICLE EXEC ERROR", err)
		return err
	}
	return c.Send("Сохранил!")

}
