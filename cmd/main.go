package main

import (
	"bufio"
	"fmt"
	"go-service-echo/config"
	"go-service-echo/domain/users"
	"go-service-echo/infrastructure/database"
	"go-service-echo/util/bcrypt"
	"go-service-echo/util/faker"
	"go-service-echo/util/logger"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}

	config := config.New()

	database, err := database.New(config.Database)
	if err != nil {
		logger.Error(err)
	}

	go doBatch(database, 100000)

	// time.Sleep(time.Second / 2)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func doBatch(database *database.Database, l int) {
	pass, _ := bcrypt.HashPassword("password")

	for i := 0; i < l; i++ {
		data := &users.User{
			ID:       uuid.New().String(),
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: pass,
		}

		if err := database.SQL.Create(data).Error; err != nil {
			logger.Error(err)
		}
	}
}
