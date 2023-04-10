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
	return func(ctx *gin.Context) {
		var newTodo models.ToDo
		if err := ctx.ShouldBindJSON(&newTodo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := ctx.GetUint("userId")
		todo := &models.ToDo{
			Task:      newTodo.Task,
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

func (c *Controller) updateTodo() gin.HandlerFunc {
	var todoId struct {
		ID uint `uri:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&todoId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var todo models.ToDo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.store.ToDo().UpdateFull(&todo, todoId.ID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, todo)
	}
}
