package handlers

import "github.com/arsu4ka/todo-auth/internal/services"

type RequestsHandler struct {
	User         services.IUserService
	Todo         services.ITodoService
	Verification services.IVerificationService
	Email        services.IEmailService
}
