package repository

import (
	"fmt"

	"github.com/Xlussov/rest-api-v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user restApi.User) (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "0", err
	}

	var id string = uuid.String()

	query := fmt.Sprintf("INSERT INTO %s (id, name, username, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, id, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return "0", err
	}

	return id, nil
}



func (r *AuthPostgres) GetUser(username, password string) (restApi.User, error) {
	var user restApi.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}