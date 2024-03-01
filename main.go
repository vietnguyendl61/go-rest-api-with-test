package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go-rest-api-with-test/handlers"
	"go-rest-api-with-test/model"
	"go-rest-api-with-test/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	//db := InitDB()
	//if err != nil {
	//	log.Fatalln("Error when init db: " + err.Error())
	//}
	//MigrateDB(db)

	userHandler := handlers.NewJobHandler()
	router := mux.NewRouter()
	router.HandleFunc("/user/create", userHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		dataResponse := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "OK",
			Message: "",
		}

		utils.ResponseWithData(w, http.StatusOK, dataResponse)
	})

	log.Println("API is running in port: " + os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		log.Fatalln("Error: " + err.Error())
	}
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	_ = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	models := []interface{}{
		&model.User{},
	}

	for _, m := range models {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Println("Error when migrate: " + err.Error())
			return
		}
	}
}
