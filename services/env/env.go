package env

import (
	"net"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type Environment struct {
	PostgresHost      string `mapstructure:"POSTGRES_HOST" default:"localhost"`
	PostgresPort      string `mapstructure:"POSTGRES_PORT" default:"5432"`
	PostgresDatabase  string `mapstructure:"POSTGRES_DATABASE" default:"postgres"`
	PostgresUsername  string `mapstructure:"POSTGRES_USER" default:"test"`
	PostgresPassword  string `mapstructure:"POSTGRES_PASSWORD" default:"test"`
	RedisHost         string `mapstructure:"REDIS_HOST" default:"localhost:6379"`
	RedisPassword     string `mapstructure:"REDIS_PASSWORD" default:""`
	GoogleProjectID   string `mapstructure:"GOOGLE_PROJECT_ID" default:"test"`
	GoogleBqProjectID string `mapstructure:"GOOGLE_BQ_PROJECT_ID" default:"prod-test"`
	GoogleCredential  string `mapstructure:"GOOGLE_APPLICATION_CREDENTIALS" default:"/home/wahyu/contoh.json"`
	GoogleBucket      string `mapstructure:"GOOGLE_BUCKET_NAME" default:"my_bucket"`
	SwaggerHost       string `mapstructure:"SWAGGER_HOST" default:"0.0.0.0:8080"`
	JwtSecretKey      string `mapstructure:"JWT_SECRET" default:"hilihkintil"`
	SendGridApiKey    string `mapstructure:"SENDGRID_API_KEY" default:"xxxxxxxxxx"`
	SendGridEndPoint  string `mapstructure:"SENDGRID_ENDPOINT" default:"/v3/mail/send"`
	SendGridHost      string `mapstructure:"SENDGRID_HOST" default:"https://api.sendgrid.com"`
	GinMode           string `mapstructure:"GIN_MODE" default:"debug"`
	AppName           string `mapstructure:"APP_NAME" default:"My-Unicorn"`
	AppHost           string `mapstructure:"APP_HOST" default:"0.0.0.0:8080"`
	AppCors           string `mapstructure:"APP_CORS" default:"*;gantidengandomain.com"`
	LogDir            string `mapstructure:"LOG_DIR" default:""`
	AdminAccount      string `mapstructure:"ADMIN_ACCOUNT" default:"wahyuhadi@gmail.com"`
	AdminPassword     string `mapstructure:"ADMIN_PASSWORD" default:"Test123!"`
	CreateAdmin       string `mapstructure:"CREATE_ADMIN" default:"true"`
	RunOnLocalMachine string `mapstructure:"RUN_ON_LOCAL_MACHINE" default:"true"`
}

var envCache = map[string]*Environment{}

func Global() *Environment {
	return Get("__global__")
}

func Get(instanceName string) *Environment {
	if env, ok := envCache[instanceName]; !ok {
		return New(instanceName)
	} else {
		return env
	}
}

func New(instanceName string) *Environment {
	var env Environment

	viper.SetConfigType("yaml")
	viper.SetConfigFile("env.yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	e := reflect.ValueOf(&env).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		key := t.Field(i).Tag.Get("mapstructure")
		def := t.Field(i).Tag.Get("default")

		viper.SetDefault(key, def)
	}

	_ = viper.Unmarshal(&env)
	if env.RunOnLocalMachine == "true" {
		_ = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", env.GoogleCredential)
	}

	envCache[instanceName] = &env
	return &env
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
