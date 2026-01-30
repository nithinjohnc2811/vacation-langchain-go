package routes
import "github.com/gin-gonic/gin"

func getVacationRouter(router *gin.Engine) *gin.Engine {
	registrationRoutes := router.Group("/vacation")
	//setting two APIs
	registrationRoutes.POST("/create", func(c *gin.Context)){
		
	}
}