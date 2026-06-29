package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
)

func (pg *Postgres) CreateRefreshTokenDB(tokenId, userId uuid.UUID, tokenHash string, expiresAt time.Time) error {
	query := `INSERT INTO refresh_tokens 
				(id, user_id, token_hash, expires_at)
			  VALUES
			  	($1, $2, $3, $4)`

	_, err := pg.Db.Exec(
					query,
					tokenId,
					userId, 
					tokenHash,
					expiresAt,
					)
	
	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) GetRefreshTokenDB(userId uuid.UUID, tokenHash string) (*modules.RefreshToken, error) {
	query := `SELECT
				id, revoked, expires_at
			  FROM
			    refresh_tokens
			  WHERE 
			  	token_hash = $1 
			  AND 
			  	user_id = $2`
	
	var refreshToken modules.RefreshToken
	err := pg.Db.QueryRow(query, tokenHash, userId).Scan(&refreshToken.Id, &refreshToken.Revoked, &refreshToken.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}