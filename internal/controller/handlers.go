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
	type Request struct {
		Task      string `json:"task" binding:"required"`
		Completed bool   `json:"completed"`
	}
	return func(ctx *gin.Context) {
		var request Request
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := ctx.GetUint("userId")
		todo := &models.ToDo{
			Task:      request.Task,
			Completed: request.Completed,
			UserID:    userID,
		}

		if err := c.store.ToDo().Create(todo); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, resptypes.NewTodoResponse(todo))
	}
}

func (c *Controller) updateTodo() gin.HandlerFunc {
	type RequestUri struct {
		ID uint `uri:"id" binding:"required"`
	}
	type RequestBody struct {
		Task      string `json:"task" bindinig:"required"`
		Completed bool   `json:"completed" bindinig:"required"`
	}
	return func(ctx *gin.Context) {
		var requestBody RequestBody
		var requestUri RequestUri

		if err := ctx.ShouldBindUri(&requestUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := c.store.ToDo().FindById(requestUri.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if todo.UserID != ctx.GetUint("userId") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you can edit only your todos"})
			return
		}

		updatedTodo := &models.ToDo{
			ID:        requestUri.ID,
			Task:      requestBody.Task,
			Completed: requestBody.Completed,
			UserID:    todo.UserID,
		}

		if err := c.store.ToDo().UpdateFull(updatedTodo); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.JSON(http.StatusOK, resptypes.NewTodoResponse(updatedTodo))
	}
}

func (c *Controller) deleteTodo() gin.HandlerFunc {
	type RequestUri struct {
		ID uint `uri:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var requestUri RequestUri

		if err := ctx.ShouldBindUri(&requestUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := c.store.ToDo().FindById(requestUri.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if todo.UserID != ctx.GetUint("userId") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you can delete only your todos"})
			return
		}

		if err := c.store.ToDo().Delete(requestUri.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, resptypes.NewTodoResponse(todo))
	}
}
