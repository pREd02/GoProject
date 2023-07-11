package handler

import (
	"GoProject/internal/service"
	"GoProject/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
)

type Handler struct {
	services *service.Service
	logger   *logger.Logger
}

func NewHandler(services *service.Service, logger *logger.Logger) *Handler {
	return &Handler{services: services,
		logger: logger}
}
func (h *Handler) InitRoutes(writer io.Writer) *gin.Engine {
	gin.DefaultWriter = writer
	router := gin.Default()
	api := router.Group("api", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.POST("/signin", h.signIn)
			user.POST("/signup", h.signUp)
		}
		organization := api.Group("/organization")
		{
			organization.POST("/register", h.registerOrganization)
			organization.GET("/:id", h.idOrganization)

			projects := api.Group("projects")
			{
				projects.POST("/create", h.createProjects)
				projects.GET("/:id", h.idProjects)
				projects.GET("/:id/users", h.idUserProjects)
				projects.POST("/submit", h.submitProjects)
			}
		}
	}
	return router
}
