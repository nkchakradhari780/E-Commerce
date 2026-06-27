package postgres

import (
	"time"

	"github.com/google/uuid"
)

func (pg *Postgres) CreateRefreshToken(userId uuid.UUID, tokenHash string, expiresAt time.Time) error {
	query := `INSERT INTO refresh_tokens 
				(user_id, token_hash, expires_at)
			  VALUES
			  	($1, $2, $3)`

	_, err := pg.Db.Exec(
					query,
					userId, 
					tokenHash,
					expiresAt,
					)
	
	if err != nil {
		return err
	}

	return nil
}