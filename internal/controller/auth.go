package controller

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/middleware"
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) registerHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.store.User().Create(&user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		user.HidePassword()
		ctx.JSON(http.StatusCreated, &user)
	}
}

func (c *Controller) loginHandler() gin.HandlerFunc {
	var loginCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(ctx *gin.Context) {
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
