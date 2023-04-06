package controller

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/controller/resptypes"
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

		todosSanitized := []*resptypes.TodoResponse{}
		for _, todo := range todos {
			todosSanitized = append(todosSanitized, resptypes.NewTodoResponse(todo))
		}
		ctx.JSON(http.StatusOK, todosSanitized)
	}
}

func (c *Controller) createTodo() gin.HandlerFunc {
	var request struct {
		Task string `json:"task" binding:"required"`
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

func (c *Controller) UpdateTodoStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
