package services

import (
	"go-starter/internal/common/apihelpers"
	"go-starter/internal/models"
	"go-starter/internal/repositories"
	"log"

	uuid "github.com/satori/go.uuid"
)

var userRepository = repositories.NewUserRepository()

func AddUser(user models.AddUserRequest) apihelpers.CreationResponse {
	log.Printf("received user - %v\n", user)

	var userId = uuid.NewV4()

	var _, err = userRepository.InsertUser(user, userId)
	if err != nil {
		log.Printf("internal server error - %s\n", err)

		return apihelpers.CreationResponse{
			Status:  500,
			Message: "Internal server error",
		}
	}

	log.Printf("inserted user - %s\n", userId)

	return apihelpers.CreationResponse{
		Status: 201,
		ID:     userId,
	}
}
