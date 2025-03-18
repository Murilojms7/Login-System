package handler

import (
	"fmt"
	"net/http"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.Users{}
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error to find user with id: %v", id))
		return
	}

	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Password != "" {
		user.Password = request.Password
	}
	if request.Phone > 0 {
		user.Phone = request.Phone
	}
	if request.Birth != "" {
		user.Birth = request.Birth
	}

	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user")
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}
	sendSucess(ctx, "update-user", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"birth": user.Birth,
	})
}
