package usecase

import "tages-autumn/internal/model"

//go:generate mockgen -source=deps.go -destination=mocks/mock.go

type Repository interface {
	SaveFile(fileName, filePath string) error
	GetPath(fileName string) (string, error)
	GetAll() ([]model.FileModel, error)
}
