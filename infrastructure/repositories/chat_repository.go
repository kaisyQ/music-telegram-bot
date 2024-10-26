package chat_repository

import (
	"context"
	"fmt"
	"kaisyq/tg/music/infrastructure/database"

	"github.com/jackc/pgx/v5"
)

type ChatRepository struct {
	Database *database.Database
}

func (repository ChatRepository) Insert(ctx context.Context, chatId uint32) error {
	sql := "insert into chats(chat_id) values (@chatId)"

	args := pgx.NamedArgs{"chatId": chatId}

	_, err := repository.Database.Pool.Exec(ctx, sql, args)

	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}
