package handlers

import (
	"backend-services/services/db/postgres"
	"backend-services/services/env"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Env      *env.Environment
	Log      *logrus.Logger
	LogFile  *os.File
	Postgres *postgres.Postgres
	//Redis    *redis.Redis
	Gin *gin.Engine
}
