package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subhandev/go-todo-gitops/app/handlers"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Todo API running, this is home route.",
		})
	})

	h := handlers.NewTodoHandler(db)

	r.GET("/todos", h.List)
	r.GET("/todos/:id", h.Get)
	r.POST("/todos", h.Create)
	r.PUT("/todos/:id", h.Update)
	r.DELETE("/todos/:id", h.Delete)

	return r
}
