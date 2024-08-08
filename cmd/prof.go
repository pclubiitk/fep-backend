package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/middleware"
	"github.com/pclubiitk/fep-backend/prof"
	"github.com/spf13/viper"
)

func profServer() *http.Server {
	PORT := viper.GetString("PORT.PROF")
	engine := gin.New()
	engine.Use(middleware.CORS())
	engine.Use(middleware.Authenticator())
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	// rc.CompanyRouter(engine)
	// application.CompanyRouter(engine)
	prof.ProfRouter(engine)

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      engine,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server
}
