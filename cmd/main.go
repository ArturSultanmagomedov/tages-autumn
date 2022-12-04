package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
	"os"
	"os/signal"
	"syscall"
	"tages-autumn/configs"
	"tages-autumn/internal/api"
	"tages-autumn/internal/repository"
	"tages-autumn/internal/usecase"
)

func main() {
	if err := run(); err != nil {
		logrus.Fatalf("%v", err)
	}
}

func run() error {
	if err := configs.Init(); err != nil {
		return err
	}
	if err := configs.InitLogger(); err != nil {
		return err
	}

	postgres, err := repository.NewPostgresDB(configs.GetPostgresConfig())
	if err != nil {
		return fmt.Errorf("failed to initialize db: %w", err)
	}
	defer postgres.Close()

	repo := repository.NewFileRepository(postgres)

	fh := &api.FileHandler{
		Uc:          usecase.NewFileUsecase(usecase.Deps{Config: configs.GetFilesConfig(), Repo: repo}),
		LoadLimiter: ratelimit.New(configs.GetLimitConfig().Write, ratelimit.Per(1)),
		ListLimiter: ratelimit.New(configs.GetLimitConfig().Read, ratelimit.Per(1)),
	}

	srv := api.NewServer(api.Deps{FileHandler: fh})

	go func() {
		if err := srv.ListenAndServe(configs.GetAddress()); err != nil {
			logrus.Fatalf("filed to init server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	defer func() {
		srv.Stop()
	}()

	return nil
}
