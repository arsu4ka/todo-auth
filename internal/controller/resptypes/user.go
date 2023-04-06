package resptypes

import "github.com/arsu4ka/todo-auth/internal/models"

type UserResponse struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		Id:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	}
}
