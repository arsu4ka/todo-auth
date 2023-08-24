package middleware

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ActiveUserMiddleware(handler *handlers.RequestsHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint("userId")

		user, err := handler.User.FindByID(userID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
			return
		} else if !user.Active {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "email address not verified yet"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
