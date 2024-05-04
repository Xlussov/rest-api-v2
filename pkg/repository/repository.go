package repository

import (
	"github.com/Xlussov/rest-api-v2"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restApi.User) (string, error) 
	GetUser(username, password string) (restApi.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}