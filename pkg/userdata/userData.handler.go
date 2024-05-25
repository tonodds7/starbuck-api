package userdata

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type IUseDataHandler struct{}

type UseDataHandler struct {
	// UseDataService userdata.IUserData
}

func NewUseDataHandler() *UseDataHandler {
	return &UseDataHandler{
		// userdata.NewUseDataHandler(),
	}
}

func UserDatasController(routerGroup *gin.RouterGroup, db *gorm.DB) {
	// userDatasDb := UserDatasDbHandler{DB: db}
	userRoutes := routerGroup.Group("/user-data")

	userRoutes.POST("/", func(c *gin.Context) {
		// userDatasDb.RegisterUser()
	})

}
