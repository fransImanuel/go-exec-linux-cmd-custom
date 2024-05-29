package env

import (
	"fmt"
	"go-exec-linux-cmd-custom/dto"
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	env := make(chan string, 1)
	//fmt.Println(os.Getenv("GO_ENV"))

	// if os.Getenv("GO_ENV") != "production" {
	// 	godotenv.Load(filepath.Join(".env"))
	// 	env <- os.Getenv(key)
	// } else {
	// }
	env <- os.Getenv(key)

	return <-env
}

func GetSMTPConfig() *dto.SMTPConfig {
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpEmail := "opsol.metacrm@gmail.com"
	smtpPassword := "ondrvqwqimgsvjmz"
	smtpName := "metaforce auto backup Logs"

	config := &dto.SMTPConfig{
		Host:     smtpHost,
		Port:     smtpPort,
		Email:    smtpEmail,
		Password: smtpPassword,
		Name:     smtpName,
	}

	// fmt.Printf("%+v", config)
	// panic(1)
	return config
}
