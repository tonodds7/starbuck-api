package routes

import (
	"backend/pkg/userdata"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AppRouter struct {
	router *gin.Engine
}

func Router(db *gorm.DB) AppRouter {
	router := AppRouter{router: gin.Default()}

	// main := router.router.Group("/")
	api := router.router.Group("/api")

	// users.AddUserController(api)
	// authentication.AddAuthController(api, db)
	// credential.CredentialHandler(api, db)
	userdata.UserDatasController(api, db)

	return router
}

func (router AppRouter) Run(addr ...string) error {
	return router.router.Run()
}
