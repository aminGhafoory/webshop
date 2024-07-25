package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aminGhafoory/webshop/controllers"
	"github.com/aminGhafoory/webshop/internal/database"
	"github.com/aminGhafoory/webshop/models"
	"github.com/aminGhafoory/webshop/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB models.PostgresConfig
}

func LoadConfig() (AppConfig, error) {
	godotenv.Load(".env")
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		return AppConfig{}, fmt.Errorf("no DB_URL provided")
	}
	c := AppConfig{
		DB: models.PostgresConfig{
			DBurl: DBURL,
		},
	}
	return c, nil
}

func main() {

	Config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error in parsing the .env file %v", err)
	}

	fmt.Println(Config.DB.DBurl)
	//DB SETUP

	Db, err := models.Open(Config.DB.DBurl)
	if err != nil {
		log.Fatalf("Error in Connecting to the DB : %v", err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatalf("Error in Connecting to the DB : %v", err)
	}

	DB := database.New(Db)

	//Controllers
	userC := controllers.Users{
		UserService: &models.UserService{
			DB: DB,
		},
	}

	r := chi.NewRouter()
	//middlwares
	r.Use(middleware.Logger)

	//handlers
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		c := views.Hello("amin")
		c.Render(context.Background(), w)
	})

	r.Get("/author", userC.TestHandler)

	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.FileServerFS(FS).ServeHTTP(w, r)

	})

	http.ListenAndServe(":3000", r)
}
