package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetTasks(c *gin.Context) {
	search := c.Query("search")
	tasks, err := h.service.GetTasks(search)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed GetTasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
