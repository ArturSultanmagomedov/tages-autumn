package api

import (
	"context"
	"fmt"
	"go.uber.org/ratelimit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// структура FileHandler должна реализовывать интерфейс FileServiceServer
var _ FileServiceServer = (*FileHandler)(nil)

type LimitConfig struct {
	Read  int
	Write int
}

type FileHandler struct {
	UnimplementedFileServiceServer

	Uc          IFileUscaseHandler
	LoadLimiter ratelimit.Limiter
	ListLimiter ratelimit.Limiter
}

func (f *FileHandler) DownloadFile(ctx context.Context, fileName *FileName) (*File, error) {
	f.LoadLimiter.Take()

	bytes, err := f.Uc.DownloadFile(ctx, fileName.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something going wrong: %v", err)
	}
	return &File{Name: fileName.Name, B: bytes}, nil
}

func (f *FileHandler) UploadFile(ctx context.Context, file *File) (*Nothing, error) {
	f.LoadLimiter.Take()

	err := f.Uc.UploadFile(ctx, file.Name, file.B)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something going wrong: %v", err)
	}
	return nil, nil
}

func (f *FileHandler) GetFilesList(ctx context.Context, _ *Nothing) (*FilesList, error) {
	//f.ListLimiter.Take()
	fmt.Println("heheheheh")
	list, err := f.Uc.GetFilesList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something going wrong: %v", err)
	}

	return &FilesList{Names: list}, nil
}
