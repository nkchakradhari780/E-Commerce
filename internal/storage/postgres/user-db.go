package postgres

import (
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