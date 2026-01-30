package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func generateVacation(r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse{

}


func getVacation(id uuid.UUID)GetVacationIdeaResponse{
	v,err:=chainsGetVacationFromDb(id)
	if err!=nil{
		return GetVacationIdeaResponse{},err
	}
	return GetVacationIdeaResponse{Id:v.Id,Completed:v.Completed,Idea:v.Idea},nil
}

func getVacationRouter(router *gin.Engine) *gin.Engine {
	registrationRoutes := router.Group("/vacation")
	//setting two APIs
	registrationRoutes.POST("/create", func(c *gin.Context){
		var req GenerateVacationIdeaRequest
		end:=c.BindJSON(&req)
		if err! = nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"message":"Bad Request "
			})
			else{
				c.JSON(http.StatusOK. generateVacation(req))
			}
		}
	})
	registrationRoutes.GET(":/id", func(c *gin.Context){
		id,err = uuid.Parse(c.Param("id"))
		if err! = nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"message":"Bad Request "
			})
			else{
				resp,err :=getVacation(id)
				if err!=nil{
					c.JSON(http.StatusNotFound,gin.H{
				"message":"Id not found "
			})
				}else{
					c.JSON(http.StatusOK,resp)
				}
			}
		}
	})
}
}