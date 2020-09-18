package routers

import(
	userControllers "goApiPractice/controllers"
	"github.com/gin-gonic/gin"
)

// UserRouter return router
func UserRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/AddUser", userControllers.Register)
	return router
}