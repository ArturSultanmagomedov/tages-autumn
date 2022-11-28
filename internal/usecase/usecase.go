package usecase

import (
	"context"
	"fmt"
	"os"
	"tages-autumn/internal/api"
)

type FilesConfig struct {
	Path string
}

type Deps struct {
	Config FilesConfig
	repo   Repository
}

type FileUsecase struct {
	deps Deps
}

func NewFileUsecase(deps Deps) api.IFileUscaseHandler {
	return &FileUsecase{deps: deps}
}

func (f FileUsecase) DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
	fn := fileName //TODO: запрос к бд
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	var b []byte
	if _, err := file.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (f FileUsecase) UploadFile(ctx context.Context, fileName string, buf []byte) error {
	file, err := os.OpenFile(f.deps.Config.Path+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0406)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	} // TODO: сохранить в бд
	if _, err := file.Write(buf); err != nil {
		return err
	}
	return nil
}

func (f FileUsecase) GetFilesList(ctx context.Context) ([]string, error) {
	var list []string //TODO: вернуть из бд
	return list, nil
}
