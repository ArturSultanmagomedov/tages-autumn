package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
	"os"
	"os/signal"
	"syscall"
	"tages-autumn/configs"
	"tages-autumn/internal/api"
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

	fh := &api.FileHandler{
		Uc:          usecase.NewFileUsecase(usecase.Deps{Config: configs.GetFilesConfig()}),
		LoadLimiter: ratelimit.New(configs.GetLimitConfig().Write),
		ListLimiter: ratelimit.New(configs.GetLimitConfig().Read),
	}

	srv := api.NewServer(api.Deps{FileHandler: fh})

	go func() {
		if err := srv.ListenAndServe(configs.GetAddress()); err != nil {
			logrus.Fatal(fmt.Errorf("filed to init server: %w", err))
		}
		//defer func() {
		//	if err := recover(); err != nil {}
		//}()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	defer func() {
		srv.Stop()
	}()

	return nil
}
