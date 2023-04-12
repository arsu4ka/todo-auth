package controller

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/controller/resptypes"
	"github.com/arsu4ka/todo-auth/internal/middleware"
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) registerHandler() gin.HandlerFunc {
	type Request struct {
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var request Request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := &models.User{
			FullName: request.FullName,
			Email:    request.Email,
			Password: request.Password,
		}

		if err := c.store.User().Create(user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, resptypes.NewUserResponse(user))
	}
}

func (c *Controller) loginHandler() gin.HandlerFunc {
	type LoginCredentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var loginCredentials LoginCredentials
		if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := c.store.User().FindByEmail(loginCredentials.Email)
		if err != nil || !user.HasPassword(loginCredentials.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		authToken, err := middleware.GenerateToken(user.ID, c.config.TokenExpiration, c.config.TokenSecret)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Header("Authorization", authToken)
		ctx.JSON(http.StatusOK, gin.H{"status": "authenticated"})
	}
}
