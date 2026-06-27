package postgres

import (
	"github.com/google/uuid"
	"github.com/nkchakradhari780/practice9/internal/modules"
)

func (pg *Postgres) CreateUserDB(user *modules.CreateUser) error {

	query := `INSERT INTO users 
				(id, name, email, password, role, address) 
			VALUES 
				($1, $2, $3, $4, $5, $6)`

	_, err := pg.Db.Exec(
		query,
		user.Id,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.Address)
	if err != nil {
		return err
	}

	return nil
}

func (pg *Postgres) GetUserByEmailDB(email string) (*modules.User, error) {
	query := `SELECT 
				id, name, email, password, Role, Address
			  FROM users
			  	WHERE email = $1`

	var user modules.User
	err := pg.Db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Address,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pg *Postgres) GetUserByIdDB(userId uuid.UUID) (*modules.GetUser, error) {
	query := `SELECT 
				id, name, email, role, address
			  FROM users
			  WHERE id = $1`

	var user modules.GetUser
	err := pg.Db.QueryRow(query, userId).Scan(
		&user.Id,
		&user.Name,
		&user.Email, 
		&user.Role, 
		&user.Address,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
