package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/arsu4ka/todo-auth/internal/handlers/dto"
	"github.com/arsu4ka/todo-auth/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateToken(userId uint, expTime int, secretKey []byte) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(expTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (rh *RequestsHandler) RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto.CreateUserDto
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := &models.User{
			FullName: request.FullName,
			Email:    request.Email,
			Password: request.Password,
			Active:   false,
		}

		if err := user.HashPassword(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := rh.User.Create(user); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		verif := models.NewVerification(user.ID)
		if err := rh.Verification.Create(verif); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err := rh.Email.SendVerificationLink(user.Email, user.FullName, verif.ID.String())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Success! Now you should verify your email."})
	}
}

func (rh *RequestsHandler) LoginHandler(tokenSecret string, tokenExpiration int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginCredentials dto.LoginUserDto
		if err := ctx.ShouldBindJSON(&loginCredentials); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := rh.User.FindByEmail(loginCredentials.Email)
		if err != nil || !user.ComparePassword(loginCredentials.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		if !user.Active {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "your email hasn't been verified yet"})
			return
		}

		authToken, err := GenerateToken(user.ID, tokenExpiration, []byte(tokenSecret))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "authenticated", "token": authToken})
	}
}

func (rh *RequestsHandler) VerifyHandler() gin.HandlerFunc {
	type RequestUri struct {
		VerifId string `uri:"id" binding:"required,uuid"`
	}
	return func(ctx *gin.Context) {
		var requestUri RequestUri

		if err := ctx.ShouldBindUri(&requestUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		verifUUID, err := uuid.Parse(requestUri.VerifId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		verif, err := rh.Verification.FindById(verifUUID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := rh.User.FindByID(verif.UserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		user.Active = true
		if err := rh.User.Update(user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rh.Verification.Delete(verif.ID)
		ctx.JSON(http.StatusOK, dto.NewResponseUserDto(user))
	}
}
