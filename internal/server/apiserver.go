package server

import (
	"github.com/arsu4ka/todo-auth/internal/dbs"
	"github.com/arsu4ka/todo-auth/internal/handlers"
	middleware "github.com/arsu4ka/todo-auth/internal/middlewares"
	"github.com/arsu4ka/todo-auth/internal/services/sqlservices"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiServer struct {
	router *gin.Engine
	db     *gorm.DB
	config *Config
}

func NewController(config *Config) (*ApiServer, error) {
	db, err := dbs.GetPostgresNoAuth(config.DBConf)
	if err != nil {
		return nil, err
	}

	return &ApiServer{
		router: gin.Default(),
		db:     db,
		config: config,
	}, nil
}

func (s *ApiServer) Start() error {
	s.configureServer()
	return s.router.Run(":" + s.config.Port)
}

func (s *ApiServer) configureServer() {
	s.router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Default(),
	)

	handler := handlers.RequestsHandler{
		User: sqlservices.NewUserService(s.db),
		Todo: sqlservices.NewTodoService(s.db),
	}
	api := s.router.Group("api/")

	authGroup := api.Group("auth/")
	authGroup.POST("/signup", handler.RegisterHandler())
	authGroup.POST("/login", handler.LoginHandler(s.config.TokenSecret, s.config.TokenExpiration))

	todoGroup := api.Group("todo/", middleware.JWTMiddleware(s.config.TokenSecret))
	todoGroup.POST("/", handler.CreateTodo())
	todoGroup.GET("/", handler.GetAllTodos())
	todoGroup.PUT("/:id", handler.UpdateTodo())
	todoGroup.DELETE("/:id", handler.DeleteTodo())
}