package handler

import (
	"errors"
	"net/http"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.Users{}
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sendError(ctx, http.StatusNotFound, "User not found")
			return
		}
		sendError(ctx, http.StatusInternalServerError, "error to seach user")
		return
	}

	sendSuccess(ctx, "show-user", gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"birth": user.Birth,
	})
}
