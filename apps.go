package main

import (
	"backend-services/handlers"
	"backend-services/handlers/users"
	"backend-services/services/constant"
	"backend-services/services/db/postgres"
	"backend-services/services/env"
	"backend-services/services/log"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/tylerb/graceful.v1"
)

// @title DMS Core Module API
// @version 0.0.1
// @description This API is used to serve DMS Front End
//
// @contact.name Data Platform
// @contact.email dataplatform@tiket.com
//
// @host 0.0.0.0:8080
//
// @schemes http https

const (
	defaultPort = ":8081"
	zeroDotZero = "0.0.0.0:8081"
	localhost   = "localhost:8081"
)

func main() {
	environ := env.Global()
	logger, file := log.New(environ.AppName, environ.LogDir)
	defer func() {
		logger.Infoln(constant.ComposeMessage(constant.LogClose, ""))
		_ = file.Close()
	}()

	logger.Infoln(constant.ComposeMessage(constant.PostgreConnect, ""))
	postgreConn := new(postgres.Postgres)
	err := postgreConn.Connect(
		environ.PostgresHost,
		environ.PostgresPort,
		environ.PostgresUsername,
		environ.PostgresPassword,
		environ.PostgresDatabase,
	)
	if err != nil {
		logger.Fatalln(err)
	}

	err = postgreConn.Extensi()
	if err != nil {
		logger.Fatalln(err)
	}
	logger.Infoln(constant.ComposeMessage(constant.ModelMigration, ""))
	err = postgreConn.ModelMigrate()
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infoln(constant.ComposeMessage(constant.ModelInit, ""))
	err = postgreConn.ModelInit()
	if err != nil {
		logger.Fatalln(err)
	}

	// err = postgreConn.CreateDefaultAdmin(environ)
	// if err != nil {
	// 	logger.Fatalln(err)
	// }

	// err = postgreConn.TypeProduct()
	// if err != nil {
	// 	logger.Fatalln(err)
	// }

	gin.SetMode(environ.GinMode)
	ginEngine := gin.Default()

	handler := handlers.Handler{
		Env:      environ,
		Log:      logger,
		LogFile:  file,
		Postgres: postgreConn,
		//Redis:    redisConn,
		Gin: ginEngine,
	}
	users.Router(handler, "/api")

	graceful.Run(environ.AppHost, 10*time.Second, ginEngine)
}
