package handlers

import (
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/handlers/dto"
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func (rh *RequestsHandler) GetAllTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint("userId")

		todos, err := rh.Todo.FindByUser(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		todosSanitized := []*dto.ResponseTodoDto{}
		for _, todo := range todos {
			todosSanitized = append(todosSanitized, dto.NewResponseTodoDto(todo))
		}
		ctx.JSON(http.StatusOK, todosSanitized)
	}
}

func (rh *RequestsHandler) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.CreateTodoDto
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := ctx.GetUint("userId")
		todo := &models.Todo{
			Task:      request.Task,
			Completed: request.Completed,
			UserID:    userID,
		}

		if err := rh.Todo.Create(todo); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, dto.NewResponseTodoDto(todo))
	}
}

func (rh *RequestsHandler) UpdateTodo() gin.HandlerFunc {
	type RequestUri struct {
		ID uint `uri:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var requestBody dto.UpdateTodoDto
		var requestUri RequestUri

		if err := ctx.ShouldBindUri(&requestUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := rh.Todo.FindById(requestUri.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if todo.UserID != ctx.GetUint("userId") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you can edit only your todos"})
			return
		}

		updatedTodo := &models.Todo{
			Task:      requestBody.Task,
			Completed: requestBody.Completed,
			UserID:    todo.UserID,
		}

		if err := rh.Todo.Update(requestUri.ID, updatedTodo); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.JSON(http.StatusOK, dto.NewResponseTodoDto(updatedTodo))
	}
}

func (rh *RequestsHandler) DeleteTodo() gin.HandlerFunc {
	type RequestUri struct {
		ID uint `uri:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var requestUri RequestUri

		if err := ctx.ShouldBindUri(&requestUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := rh.Todo.FindById(requestUri.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if todo.UserID != ctx.GetUint("userId") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you can delete only your todos"})
			return
		}

		if err := rh.Todo.Delete(requestUri.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, dto.NewResponseTodoDto(todo))
	}
}
