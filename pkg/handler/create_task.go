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

func (h *Handler) CreateTask(c *gin.Context) {

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
	if task.Repeat == "" || task.Date == dateTaskMow {
		task.Date = dateTaskMow
	} else {
		task.Date, err = usecase.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			log.Printf("Failed to NextDate: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	id, err := h.service.CreateTask(task)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})

}

func CheckRequest(task *model.Task, dateTaskNow string) error {
	if task.Title == "" {
		return fmt.Errorf("title is empty")
	}
	if task.Date == "" {
		task.Date = dateTaskNow
		return nil
	}
	_, err := time.Parse(model.TimeFormat, task.Date)
	if err != nil {
		return fmt.Errorf("date is invalid")
	}
	if task.Date < dateTaskNow && task.Repeat == "" {
		task.Date = dateTaskNow
	}

	return nil
}
