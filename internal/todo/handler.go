package todo

import (
	"context"
	"net/http"
	"strconv"

	"github.com/berkantay/todo-app-example/internal/entity"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

type Service interface {
	CreateTodo(context.Context, *entity.Todo) (*entity.Todo, error)
	GetAllTodo(context.Context) ([]*entity.Todo, error)
	DeleteTodo(context.Context, int) (*entity.Todo, error)
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var req entity.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := h.Service.CreateTodo(c.Request.Context(), &entity.Todo{
		Description: req.Description,
		Deadline:    req.Deadline,
		Priority:    req.Priority,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Service.DeleteTodo(c.Request.Context(), todoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllTodo(c *gin.Context) {
	res, err := h.Service.GetAllTodo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
