package repository

import (
	"database/sql"
	"tages-autumn/internal/usecase"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

func NewRepository(db *sql.DB) *usecase.Repository {
	return &usecase.Repository{
		FileRepo: NewFileRepository(db),
	}
}
