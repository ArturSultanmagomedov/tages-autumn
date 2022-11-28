package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"tages-autumn/internal/api"
	"tages-autumn/internal/repository"
	"tages-autumn/internal/usecase"
	"time"
)

func Init() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error initializing config: %w", err)
	}
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading env variables: %w", err)
	}

	return nil
}

func InitLogger() error {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	level, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	if viper.GetString("log.output") != "" {
		currentTime := viper.GetString("log.output") + time.Now().In(time.UTC).Format("20060102_150405") + ".txt"
		f, err := os.OpenFile(currentTime, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0406)
		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
		logrus.SetOutput(f)
	}

	return nil
}

func GetPostgresConfig() repository.Config {
	return repository.Config{
		Username: viper.GetString("db.username"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func GetAddress() string {
	return fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
}

func GetLimitConfig() api.LimitConfig {
	return api.LimitConfig{
		Read:  viper.GetInt("limit.read"),
		Write: viper.GetInt("limit.write"),
	}
}

func GetFilesConfig() usecase.FilesConfig {
	return usecase.FilesConfig{
		Path: viper.GetString("files.path"),
	}
}
