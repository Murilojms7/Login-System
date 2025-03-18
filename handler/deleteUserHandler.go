package handler

import (
	"fmt"
	"net/http"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
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

	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id: %v", id))
		return
	}
	sendSucess(ctx, "delete-user", gin.H{
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"Birth": user.Birth,
	})
}
