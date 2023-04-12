package controller

import (
	"github.com/arsu4ka/todo-auth/internal/dbs"
	"github.com/arsu4ka/todo-auth/internal/middleware"
	"github.com/arsu4ka/todo-auth/internal/store"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	router *gin.Engine
	store  *store.Store
	config *Config
}

func NewController(config *Config) (*Controller, error) {
	db, err := dbs.GetPostgresNoAuth(config.DBConf)
	if err != nil {
		return nil, err
	}

	store := store.NewStore(db)
	return &Controller{
		router: gin.Default(),
		store:  store,
		config: config,
	}, nil
}

func (c *Controller) Start() error {
	c.configureController()
	return c.router.Run(":" + c.config.Port)
}

func (c *Controller) configureController() {
	c.router.Use(
		gin.Recovery(),
		gin.Logger(),
		cors.Default(),
	)

	api := c.router.Group("api/")
	api.POST("/signup", c.registerHandler())
	api.POST("/login", c.loginHandler())

	protected := c.router.Group("api/todo", middleware.JWTMiddleware(c.config.TokenSecret))
	protected.POST("/", c.createTodo())
	protected.GET("/", c.getAllTodos())
	protected.PUT("/:id", c.updateTodo())
	protected.DELETE("/:id", c.deleteTodo())
}
