package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arsu4ka/todo-auth/internal/handlers/dto"
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (rh *RequestsHandler) GetAllTodos() gin.HandlerFunc {
	type RequestQuery struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}
	return func(ctx *gin.Context) {
		var requestQuery RequestQuery
		if err := ctx.ShouldBindQuery(&requestQuery); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID := ctx.GetUint("userId")
		todos, err := rh.Todo.FindByUser(userID, requestQuery.Limit, requestQuery.Page)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if requestQuery.Page != 0 && requestQuery.Limit != 0 {
			recordNum, err := rh.Todo.GetTotalRecordCount()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx.Header("X-Total-Count", fmt.Sprint(recordNum))
		}

		todosSanitized := []*dto.ResponseTodoDto{}
		for _, todo := range todos {
			todosSanitized = append(todosSanitized, dto.NewResponseTodoDto(todo))
		}

		ctx.JSON(http.StatusOK, todosSanitized)
	}
}

func (rh *RequestsHandler) GetOneTodo() gin.HandlerFunc {
	type RequestUri struct {
		ID string `uri:"id" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var requestUri RequestUri

		if err := ctx.ShouldBindUri(&requestUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId := ctx.GetUint("userId")
		todo, err := rh.Todo.FindById(requestUri.ID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if todo.UserID != userId {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}

		ctx.JSON(http.StatusOK, dto.NewResponseTodoDto(todo))
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
			Task:        request.Task,
			Description: request.Description,
			Completed:   request.Completed,
			UserID:      userID,
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
		ID string `uri:"id" binding:"required"`
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
			Task:        requestBody.Task,
			Description: requestBody.Description,
			Completed:   requestBody.Completed,
			UserID:      todo.UserID,
		}

		if err := rh.Todo.Update(requestUri.ID, updatedTodo); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.JSON(http.StatusOK, dto.NewResponseTodoDto(updatedTodo))
	}
}

func (rh *RequestsHandler) DeleteTodo() gin.HandlerFunc {
	type RequestUri struct {
		ID string `uri:"id" binding:"required"`
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
