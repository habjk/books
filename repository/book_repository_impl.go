package repository

import "database/sql"

type BookRepositoryImpl struct{
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository{
	return &BookRepositoryImpl(Db:Db)
}