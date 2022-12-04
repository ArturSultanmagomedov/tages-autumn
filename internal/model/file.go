package model

import "time"

type FileModel struct {
	Id           string    `db:"id"`
	FileName     string    `db:"file_name"` //TODO: вместо имени использовать хэш
	FilePath     string    `db:"file_path"`
	CreationTime time.Time `db:"creation_time"`
	UpdateTime   time.Time `db:"update_time"`
}

// GetFields чтобы передавать в sql.Scan() все поля структуры FileModel
func (r *FileModel) GetFields() []any {
	return []any{&r.Id, &r.FileName, &r.FilePath, &r.CreationTime, &r.UpdateTime}
}
