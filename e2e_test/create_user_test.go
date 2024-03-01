package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go-rest-api-with-test/model"
	"go-rest-api-with-test/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
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

func (s *EndToEndSuite) TestSucceedAPICreateUser(t *testing.T) {
	c := http.Client{}
	idString := "3fd36bb2-43df-44b4-a74e-2e9a5504738e"
	id, _ := uuid.Parse(idString)

	createUserRequest := model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
		UserName: "vietnguyen",
		Password: "password",
		Address:  "123",
		Email:    "vietnguyen@gmail.com",
	}

	tmp, _ := json.Marshal(&createUserRequest)
	requestBody := bytes.NewReader(tmp)
	response, _ := c.Post("http://localhost:8001/user/create", "application/json", requestBody)

	s.Equal(http.StatusCreated, response.StatusCode)

	db := InitDB()
	userRepo := repo.NewUserRepo(db)
	userRepo.GetOne(context.Background(), id.String())

	expectJson := `{"error":"UserName can not be empty"}`
	responseBody, _ := io.ReadAll(response.Body)
	s.JSONEq(expectJson, string(responseBody))
}

func (s *EndToEndSuite) TestMissingFieldAPICreateUser(t *testing.T) {
	c := http.Client{}
	idString := "3fd36bb2-43df-44b4-a74e-2e9a5504738e"
	id, _ := uuid.Parse(idString)

	// Missing UserName field
	createUserRequest := model.User{
		BaseModel: model.BaseModel{
			ID: id,
		},
		Password: "abc",
		Address:  "123",
		Email:    "vietnguyen@gmail.com",
	}

	tmp, _ := json.Marshal(&createUserRequest)

	requestBody := bytes.NewReader(tmp)

	response, _ := c.Post("http://localhost:8001/user/create", "application/json", requestBody)

	s.Equal(http.StatusBadRequest, response.StatusCode)

	expectJson := `{"error":"UserName can not be empty"}`
	responseBody, _ := io.ReadAll(response.Body)
	s.JSONEq(expectJson, string(responseBody))
}
