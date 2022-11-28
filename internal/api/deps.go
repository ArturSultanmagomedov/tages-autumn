package api

import "context"

type IFileUscaseHandler interface {
	DownloadFile(ctx context.Context, fileName string) ([]byte, error)
	UploadFile(ctx context.Context, fileName string, buf []byte) error
	GetFilesList(ctx context.Context) ([]string, error)
}
