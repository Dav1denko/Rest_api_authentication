package handler

import (
	"Rest_api_authentication/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/:GUID", h.GetTokens)
		auth.GET("/:GUID/refresh/", h.RefreshTokens)
	}
	return router
}
