package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Database struct {
		Name string
		Port string
		Pass string
		Host string
		User string
	}

	Runn struct {
		Host string
		Port string
	}
}

var app *App
var DB *gorm.DB

func GetConnection() *gorm.DB {
	conf := GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s passowrd=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.Database.Host,
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Name,
		conf.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	errOtel := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName(conf.Database.Name)))
	if errOtel != nil {
		log.Fatal(errOtel)
	}

	DB = db
	return db
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}
func initConfig() *App {
	conf := App{}
	err := godotenv.Load()
	if err != nil {
		conf.Database.Name = "dbexample"
		conf.Database.Pass = "12345"
		conf.Database.User = "root"
		conf.Database.Host = "localhost"
		conf.Database.Port = ""

		return &conf
	}

	conf.Database.Host = os.Getenv("POSTGRES_HOST")
	conf.Database.Name = os.Getenv("POSTGRES_NAME")
	conf.Database.Port = os.Getenv("POSTGRES_PORT")
	conf.Database.User = os.Getenv("POSTGRES_USER")
	conf.Database.Pass = os.Getenv("POSTGRES_PASS")

	conf.Runn.Host = os.Getenv("APP_HOST")
	conf.Runn.Port = os.Getenv("APP_PORT")

	return &conf
}
