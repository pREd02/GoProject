package main

import (
	"GoProject"
	"GoProject/internal/handler"
	repository2 "GoProject/internal/reposistory"
	"GoProject/internal/service"
	"GoProject/pkg/logger"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("error loading env variables: %s", err.Error())
	}
}

func main() {

	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			log.Println(err)
			return
		}
	}
	file, err := os.Create(fmt.Sprintf("./logs/[%s].txt", time.Now().Format("01-02-2006_15-04-05")))
	if err != nil {
		log.Println(err)
		return
	}
	writers := io.MultiWriter(os.Stdout, file)
	if err = initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	loggers := logger.GetLogger(viper.GetString("logger_level"), writers)
	db, err := repository2.NewMySQLDB(repository2.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		loggers.Errorf("failed to initialize db:%s", err.Error())
	}
	fmt.Println("Запуск веб-сервера на http://localhost:8000/")
	repos := repository2.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, &loggers)
	srv := new(GoProject.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(writers)); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())

	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
