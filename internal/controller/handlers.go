package controller

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (c *Controller) getAllTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint("userId")

		todos, err := c.store.ToDo().GetByUser(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, todos)
	}
}

func (c *Controller) createTodo() gin.HandlerFunc {
	var request struct {
		Task string `json:"task"`
	}
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := ctx.GetUint("userId")
		todo := &models.ToDo{
			Task:      request.Task,
			Completed: false,
			UserID:    userID,
		}

		if err := c.store.ToDo().Create(todo); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
	}
}
