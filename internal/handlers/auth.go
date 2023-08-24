package handlers

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/handlers/dto"
	middleware "github.com/arsu4ka/todo-auth/internal/middlewares"
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (rh *RequestsHandler) RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.CreateUserDto
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := &models.User{
			FullName: request.FullName,
			Email:    request.Email,
			Password: request.Password,
		}

		if err := user.HashPassword(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := rh.User.Create(user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, dto.NewResponseUserDto(user))
	}
}

func (rh *RequestsHandler) LoginHandler(tokenSecret string, tokenExpiration int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginCredentials dto.LoginUserDto
		if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := rh.User.FindByEmail(loginCredentials.Email)
		if err != nil || !user.ComparePassword(loginCredentials.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		authToken, err := middleware.GenerateToken(user.ID, tokenExpiration, []byte(tokenSecret))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "authenticated", "token": authToken})
	}
}
