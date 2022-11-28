package usecase

type FileRepo interface {
	SaveFile(fileName, filePath string) error
	GetPath(fileName string) (string, error)
}

type Repository struct {
	FileRepo
}
