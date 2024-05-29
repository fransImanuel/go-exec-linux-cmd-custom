package env

import (
	"fmt"
	"go-exec-linux-cmd-custom/constant"
	"go-exec-linux-cmd-custom/dto"
	"os"
	"strconv"

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
	smtpHost := GodotEnv(constant.ENV_SMTP_HOST)
	smtpPort, err := strconv.Atoi(GodotEnv(constant.ENV_SMTP_PORT))
	if err != nil {
		fmt.Println("GetSMTPConfig() - error while parsing smtp port: ", err)
	}
	smtpEmail := GodotEnv(constant.ENV_SMTP_EMAIL)
	smtpPassword := GodotEnv(constant.ENV_SMTP_PASSWORD)
	smtpName := GodotEnv(constant.ENV_SMTP_NAME)

	fmt.Println(smtpHost, smtpPort, smtpEmail, smtpPassword, smtpName)

	if smtpHost == "" {
		smtpHost = "smtp.gmail.com"
	}

	if smtpPort < 1 {
		smtpPort = 587
	}

	if smtpEmail == "" {
		smtpEmail = "opsol.metacrm@gmail.com"
	}

	if smtpPassword == "" {
		smtpPassword = "ondrvqwqimgsvjmz"
	}

	if smtpName == "" {
		smtpName = "metaforce auto backup Logs"
	}

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
