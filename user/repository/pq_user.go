package repository

import (
	"database/sql"
)

type postgreUserRepo struct {
	DB *sql.DB
}

//func NewPostgreUserRepository(db *sql.DB) domain.UserRepository{
//	return &postgreUserRepo{
//		DB: db,
//	}
//}

