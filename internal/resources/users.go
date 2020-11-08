package resources

import (
	"fmt"
	"go-starter/internal/models"
	"go-starter/internal/services"

	"github.com/gin-gonic/gin"
)

func NewUserResource() *UserResource {
	return &UserResource{}
}

type UserResource struct {
}

func (resource UserResource) AddUser(c *gin.Context) {
	var user models.AddUserRequest
	e := c.BindJSON(&user)
	if e != nil {
		fmt.Println(e)
		c.JSON(400, "failed to read user request body.")
	}

	response := services.AddUser(user)
	if response.Status == 201 {
		c.JSON(201, gin.H{"id": response.ID})
	} else {
		c.JSON(response.Status, gin.H{"message": response.Message})
	}
}
