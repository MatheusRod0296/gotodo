package sorterUrl

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	service *URLService
}

func NewHandler(service *URLService) *UrlHandler {
	return &UrlHandler{service}
}

func (h *UrlHandler) Create(c *gin.Context) {
	var body struct {
		Url string `json:"url"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	shortUrl, err := h.service.CreateShortURL(body.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shortUrl)

}

func (h *UrlHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	url, err := h.service.GetOriginalURL(code)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.Redirect(302, url)
}

func (h *UrlHandler) List(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))

	shortUrls, err := h.service.ListShortURLs(offset, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(200, shortUrls)
}
