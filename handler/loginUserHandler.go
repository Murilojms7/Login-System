package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/Murilojms7/Login-System/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoginUserHandler(ctx *gin.Context) {
	request := LoginRequest{}
	ctx.BindJSON(&request)
	if err := request.validateLogin(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.Users{}
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "User not found")
			return
		}
		sendError(ctx, http.StatusInternalServerError, "error to seach user")
		return
	}

	match := utils.CheckPasswordHash(request.Password, user.Password)
	if !match {
		sendError(ctx, http.StatusNotFound, "Email or Password incorrect")
		return
	}

	sendSucess(ctx, "login-user", fmt.Sprintf("success login! Welcome %v", user.Name))
}
