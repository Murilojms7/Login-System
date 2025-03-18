package handler

import (
	"fmt"
	"net/http"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.Users{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
		Birth:    request.Birth,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user on database ")
		return
	}
	sendSucess(ctx, "createUser", fmt.Sprintf("user: %v created successfully", user.Email))
}
