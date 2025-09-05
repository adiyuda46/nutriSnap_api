package repository

import (
	"database/sql"
)

type Repository interface {
	GetEmailRepository(int) (string, error)
}

type DbRepository struct {
	db *sql.DB
}

// GetEmail implements Repository.
func (nr *DbRepository) GetEmailRepository(id int) (string, error) {
	query := `select "EMAIL" FROM public."USER_ACCESS" WHERE "USER_ID" = $1`

var	result string 
	err := nr.db.QueryRow(query,id).Scan(&result)
	if err != nil {
		return "", err
	}

	return result , nil 
}

func CreateRepository(db *sql.DB) Repository {
	return &DbRepository{db: db}
}
