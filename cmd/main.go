package main


import (

	"log"
	"github.com/spf13/viper"
	"github.com/Beksultan15/go_project"
	"github.com/Beksultan15/go_project/pkg/handler"
	"github.com/Beksultan15/go_project/pkg/repository"
	"github.com/Beksultan15/go_project/pkg/service"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
    }

	db,err := repository.NewPostgresConfig(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("here is the error: %s", err.Error())

	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handler.NewHandler(services)

	srv := new(go_project.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil{
		log.Fatalf(err.Error())
	}

}

func initConfig()  error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}