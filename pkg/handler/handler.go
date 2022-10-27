package handler

import (
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/service"
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
	router.LoadHTMLGlob("./ui/html/*")
	router.Static("/ui/static", "./ui/static/")
	router.GET("/", h.home)

	auth := router.Group("/auth")
	{
		auth.GET("/register", h.register)
		auth.GET("/login", h.login)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-out", h.signOut)
	}

	api := router.Group("/api", h.userIdentity)
	{
		notes := api.Group("/notes")
		{
			notes.GET("/", h.getAllNotes)
			notes.GET("/create", h.createNote)
			notes.POST("/", h.createdNote)
			notes.GET("/:id", h.getNoteById)
			notes.POST("/update/:id", h.updatedNote)
			notes.GET("/update/:id", h.updateNote)
			notes.POST("/:id", h.deleteNote)
		}
	}

	return router
}
