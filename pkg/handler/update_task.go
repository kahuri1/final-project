package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/final-project/pkg/model"
	"github.com/kahuri1/final-project/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (h *Handler) UpdateTask(c *gin.Context) {
	var task model.Task
	d, err := c.GetRawData()

	err = json.Unmarshal(d, &task)
	if err != nil {
		log.Errorf("unmarshal handlerError")
		return
	}
	dateTaskMow := time.Now().Format(model.TimeFormat)
	err = CheckRequest(&task, dateTaskMow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.Date, err = usecase.NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		log.Printf("Failed to NextDate: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.UpdateTask(&task)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": fmt.Sprintf("%d", id)})
	log.Info("message created")
}
