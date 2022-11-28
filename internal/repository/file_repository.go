package repository

import (
	"database/sql"
	"fmt"
	"tages-autumn/internal/model"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) SaveFile(fileName, filePath string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("filed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	return nil
}

func (r *FileRepository) GetPath(fileName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FileRepository) GetAll() ([]model.FileModel, error) {
	var files []model.FileModel
	//rows, err := r.db.Query("select id, file_name, file_path, creation_time, update_time from files")
	var err error
	if err != nil {
		return nil, fmt.Errorf("failed to get : %w", err)
	}
	//r.db.

	return files, err
}
