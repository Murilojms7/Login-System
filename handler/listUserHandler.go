package handler

import (
	"net/http"

	"github.com/Murilojms7/Login-System/schemas"
	"github.com/gin-gonic/gin"
)

func ListUserHandler(ctx *gin.Context) {
	users := []schemas.Users{}
	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing users")
		return
	}

	// Criando um slice de DTO
	var usersDTO []gin.H
	for _, user := range users {
		usersDTO = append(usersDTO, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"phone": user.Phone,
			"birth": user.Birth,
		})
	}

	sendSucess(ctx, "list-users", usersDTO)
}
