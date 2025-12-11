package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler contém a dependência do service
type Handler struct {
	service *Service
}

// Construtor
func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

// GET /todos
func (h *Handler) List(c *gin.Context) {
	todos := h.service.List()
	c.JSON(http.StatusOK, todos)
}

func (h *Handler) GetById(c *gin.Context) {
	var idStr = c.Params.ByName("id")
	n, _ := strconv.Atoi(idStr)

	todos, error := h.service.GetById(n)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
	}

	c.JSON(http.StatusOK, todos)
}

// POST /todos
func (h *Handler) Create(c *gin.Context) {
	var body struct {
		Title string `json:"title"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	todo, err := h.service.Create(body.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// PUT /todos/:id
func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	todo, err := h.service.Update(id, body.Title, body.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DELETE /todos/:id
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
