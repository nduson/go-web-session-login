package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"html/template"
	//"strings"

	globals "gin_session_auth/globals"
	middleware "gin_session_auth/middleware"
	routes "gin_session_auth/routes"
)

func main() {
	router := gin.Default()

	//router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.tmpl")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("localhost:8081")
}
