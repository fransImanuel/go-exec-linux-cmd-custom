package constant

import (
	"fmt"
	"os"

	"go-exec-linux-cmd-custom/schemas"

	"github.com/joho/godotenv"
)

var ConfigEnv schemas.SchemaEnvironment

const (
	ENV_DOMAIN_IMAGE  = "DOMAIN_IMAGE"
	ENV_GO_ENV        = "GO_ENV"
	ENV_GO_ENV_SDH    = "GO_ENV_SDH"
	ENV_REST_PORT     = "REST_PORT"
	ENV_REST_PORT_SDH = "REST_PORT_SDH"
	ENV_TIMEZONE      = "TIMEZONE"
	ENV_JWT_SECRET    = "JWT_SECRET"
	ENV_SWAGGER_HOST  = "SWAGGER_HOST"
	ENV_DOMAIN        = "DOMAIN"
	ENV_SDH_API_AUTH  = "SDH_API_AUTH"

	ENV_DB_HOST    = "DB_HOST"
	ENV_DB_PORT    = "DB_PORT"
	ENV_DB_USER    = "DB_USER"
	ENV_DB_NAME    = "DB_NAME"
	ENV_DB_PASS    = "DB_PASS"
	ENV_DB_SSLMODE = "DB_SSLMODE"

	ENV_REDIS_HOST     = "REDIS_HOST"
	ENV_REDIS_PORT     = "REDIS_PORT"
	ENV_REDIS_PASSWORD = "REDIS_PASSWORD"
	ENV_REDIS_DB       = "REDIS_DB"

	ENV_SMTP_HOST     = "SMTP_HOST"
	ENV_SMTP_PORT     = "SMTP_PORT"
	ENV_SMTP_EMAIL    = "SMTP_EMAIL"
	ENV_SMTP_PASSWORD = "SMTP_PASSWORD"
	ENV_SMTP_NAME     = "SMTP_NAME"

	ENV_ONE_SIGNAL_APP_ID  = "ENV_ONE_SIGNAL_APP_ID"
	ENV_ONE_SIGNAL_API_KEY = "ENV_ONE_SIGNAL_API_KEY"

	ENV_EXTERNAL_CREDENTIAL = "EXTERNAL_CREDENTIAL"
	ENV_MIDTRANS_HOST       = "MIDTRANS_HOST"
	ENV_MIDTRANS_API        = "MIDTRANS_API"
	ENV_MIDTRANS_SERVER_KEY = "MIDTRANS_SERVER_KEY"
	ENV_MIDTRANS_CALLBACK   = "MIDTRANS_CALLBACK"

	ENV_MINIO_HOST       = "MINIO_HOST"
	ENV_MINIO_ACCESS_KEY = "MINIO_ACCESS_KEY"
	ENV_MINIO_SECRET_KEY = "MINIO_SECRET_KEY"
	ENV_MINIO_LOCATION   = "MINIO_LOCATION"
	ENV_MINIO_SSL        = "MINIO_SSL"
	ENV_MINIO_DOMAIN     = "MINIO_DOMAIN"

	DEFAULT_SMTP_HOST     = "smtp.gmail.com"
	DEFAULT_SMTP_PORT     = 587
	DEFAULT_SMTP_EMAIL    = "sdh.notification@gmail.com"   //!
	DEFAULT_SMTP_PASSWORD = "Visionet*1!"                  //!
	DEFAULT_SMTP_NAME     = "San Diego Hills Notification" //!
)

func Environment(file_env string) (config schemas.SchemaEnvironment) {
	err := godotenv.Load(file_env)
	if err != nil {
		fmt.Println("Error loading .env file - ", err.Error())
	}

	// Read environment variables from docker-compose.yml
	config.DB_HOST = os.Getenv(ENV_DB_HOST)
	config.DB_PORT = os.Getenv(ENV_DB_PORT)
	config.DB_USER = os.Getenv(ENV_DB_USER)
	config.DB_NAME = os.Getenv(ENV_DB_NAME)
	config.DB_PASS = os.Getenv(ENV_DB_PASS)
	config.DB_SSLMODE = os.Getenv(ENV_DB_SSLMODE)

	config.DOMAIN = os.Getenv(ENV_DOMAIN)
	config.SMTP_NAME = os.Getenv(ENV_SMTP_NAME)
	config.SMTP_HOST = os.Getenv(ENV_SMTP_HOST)
	config.SMTP_PORT = os.Getenv(ENV_SMTP_PORT)
	config.SMTP_EMAIL = os.Getenv(ENV_SMTP_EMAIL)
	config.SMTP_PASSWORD = os.Getenv(ENV_SMTP_PASSWORD)

	config.ONE_SIGNAL_APP_ID = os.Getenv(ENV_ONE_SIGNAL_APP_ID)
	config.ONE_SIGNAL_API_KEY = os.Getenv(ENV_ONE_SIGNAL_API_KEY)

	config.NEW_RELIC_LICENCE = os.Getenv("NEW_RELIC_LICENCE_KEY")
	config.NEW_RELIC_APP_NAME = os.Getenv("NEW_RELIC_APP_NAME")

	config.MINIO_HOST = os.Getenv(ENV_MINIO_HOST)
	config.MINIO_ACCESS_KEY = os.Getenv(ENV_MINIO_ACCESS_KEY)
	config.MINIO_SECRET_KEY = os.Getenv(ENV_MINIO_SECRET_KEY)
	config.MINIO_LOCATION = os.Getenv(ENV_MINIO_LOCATION)
	config.MINIO_SSL = os.Getenv(ENV_MINIO_SSL)

	config.SDH_API_AUTH = os.Getenv(ENV_SDH_API_AUTH)
	config.EXTERNAL_CREDENTIAL = os.Getenv(ENV_EXTERNAL_CREDENTIAL)
	config.MIDTRANS_HOST = os.Getenv(ENV_MIDTRANS_HOST)
	config.MIDTRANS_API = os.Getenv(ENV_MIDTRANS_API)
	config.MIDTRANS_SERVER_KEY = os.Getenv(ENV_MIDTRANS_SERVER_KEY)
	config.MIDTRANS_CALLBACK = os.Getenv(ENV_MIDTRANS_CALLBACK)

	config.TIMEZONE = os.Getenv(ENV_TIMEZONE)
	config.REST_PORT = os.Getenv(ENV_REST_PORT)
	config.GO_ENV = os.Getenv(ENV_GO_ENV)
	config.SWAGGER_HOST = os.Getenv(ENV_SWAGGER_HOST)
	config.JWT_SECRET = os.Getenv(ENV_JWT_SECRET)
	config.DOMAIN_IMAGE = os.Getenv(ENV_DOMAIN_IMAGE)

	return config
}
