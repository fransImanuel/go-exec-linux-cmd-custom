package schemas

type SchemaEnvironment struct {
	TIMEZONE     string
	VERSION      string
	REST_PORT    string
	GO_ENV       string
	SWAGGER_HOST string
	JWT_SECRET   string
	DOMAIN       string

	DB_USER    string
	DB_PASS    string
	DB_HOST    string
	DB_PORT    string
	DB_NAME    string
	DB_SSLMODE string

	REDIS_HOST     string
	REDIS_PASSWORD string
	REDIS_DB       string

	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
	SMTP_NAME     string

	ONE_SIGNAL_APP_ID  string
	ONE_SIGNAL_API_KEY string

	NEW_RELIC_LICENCE  string
	NEW_RELIC_APP_NAME string

	MINIO_HOST       string
	MINIO_LOCATION   string
	MINIO_ACCESS_KEY string
	MINIO_SECRET_KEY string
	MINIO_SSL        string

	SDH_API_AUTH        string
	EXTERNAL_CREDENTIAL string
	MIDTRANS_HOST       string
	MIDTRANS_API        string
	MIDTRANS_SERVER_KEY string
	MIDTRANS_CALLBACK   string

	DOMAIN_IMAGE string
}
