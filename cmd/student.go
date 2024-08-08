package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/mail"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/pclubiitk/fep-backend/project"
	"github.com/pclubiitk/fep-backend/student"
	"github.com/spf13/viper"
)

func studentServer(mail_channel chan mail.Mail) *http.Server {
	PORT := viper.GetString("PORT.STUDENT")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(gin.Logger())
	student.StudentRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
func resultServer() *http.Server {
	PORT := viper.GetString("PORT.RESULT")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	project.ProjectResultRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
