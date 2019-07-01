package controllers

import(
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/biezhi/gorm-paginator/pagination"
	"cart/models"
	"fmt"
	"strconv"
)

type UserController struct {
	Db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Db: db,
	}
}

func (contr *UserController) DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	d := contr.Db.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func (contr *UserController) UpdatePerson(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	if err := contr.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)

	contr.Db.Save(&user)
	c.JSON(200, user)
}

func (contr *UserController) CreatePerson(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	contr.Db.Create(&user)
	c.JSON(200, user)
}

func (contr *UserController) GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := contr.Db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func (contr *UserController) GetPeople(c *gin.Context) {
	var user []models.User
	if err := contr.Db.Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func (contr *UserController) GetPeopleWithPaginator(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))
	var user []models.User

    paginator := pagination.Paging(&pagination.Param{
        DB:      contr.Db,
        Page:    page,
        Limit:   limit,
		ShowSQL: true,
		OrderBy: []string{"id desc"},
    }, &user)
    c.JSON(200, paginator)
}