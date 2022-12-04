package usecase

import (
	"context"
	"fmt"
	"io/ioutil"
	cache2 "tages-autumn/internal/cache"
)

type FilesConfig struct {
	Path string
}

type Deps struct {
	Config FilesConfig
	Repo   Repository
}

type FileUsecase struct {
	deps Deps
}

func NewFileUsecase(deps Deps) *FileUsecase {
	return &FileUsecase{deps: deps}
}

func (f FileUsecase) DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
	path, err := f.deps.Repo.GetPath(fileName)
	if err != nil {
		return nil, fmt.Errorf("something wrong with database: %w", err)
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("something wrong with file: %w", err)
	}
	return file, nil
}

func (f FileUsecase) UploadFile(ctx context.Context, fileName string, buf []byte) error {
	err := ioutil.WriteFile(f.deps.Config.Path+fileName, buf, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	if err := f.deps.Repo.SaveFile(fileName, f.deps.Config.Path+fileName); err != nil {
		return err
	}
	return nil
}

func (f FileUsecase) GetFilesList(ctx context.Context) ([]string, error) {
	var list []string
	if len(cache2.GetFileListCache()) > 0 {
		return cache2.GetFileListCache(), nil
	} else {
		m, err := f.deps.Repo.GetAll()
		if err != nil {
			return nil, err
		}
		for i := range m {
			list = append(list, fmt.Sprintf("%s | %s | %s", m[i].FileName, m[i].CreationTime, m[i].UpdateTime))
		}
		cache2.UpdateFileListCache(list)
	}
	return list, nil
}
