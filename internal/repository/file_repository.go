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
	query := `insert into files (file_name, file_path) values ($1, $2) 
              on conflict (file_name) do update
              set file_path = excluded.file_path, update_time = current_timestamp;`
	if _, err := r.db.Exec(query, fileName, filePath); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (r *FileRepository) GetPath(fileName string) (string, error) {
	path := ""
	if err := r.db.QueryRow("select file_path from files where file_name = $1", fileName).
		Scan(&path); err != nil {
		return "", fmt.Errorf("%w", err)
	}
	return path, nil
}

func (r *FileRepository) GetAll() ([]model.FileModel, error) {
	var files []model.FileModel
	rows, err := r.db.Query("select id, file_name, file_path, creation_time, update_time from files")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	for rows.Next() {
		m := model.FileModel{}
		if err := rows.Scan(m.GetFields()...); err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		files = append(files, m)
	}
	return files, err
}
