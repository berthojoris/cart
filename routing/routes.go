package routing

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"gin-framework/controllers"
)

func RegisterRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userController := controllers.NewUserController(db)

	r.GET("/people", userController.GetPeople)
	r.GET("/paginator", userController.GetPeopleWithPaginator)
	r.GET("/people/:id", userController.GetPerson)
	r.POST("/people", userController.CreatePerson)
	r.PUT("/people/:id", userController.UpdatePerson)
	r.DELETE("/people/:id", userController.DeletePerson)

	return r
}