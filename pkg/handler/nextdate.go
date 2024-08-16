package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kahuri1/final-project/pkg/model"
	"github.com/kahuri1/final-project/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (h *Handler) NextDateHandler(c *gin.Context) {
	nowStr := c.Query("now")
	dateStr := c.Query("date")
	repeat := c.Query("repeat")

	// Парсинг текущей даты
	now, err := time.Parse(model.TimeFormat, nowStr)
	if err != nil {
		log.WithFields(log.Fields{
			"now": nowStr,
		}).Error("invalid 'now' date format")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid 'now' date format"})
		return
	}

	// Вызываем функцию NextDate
	nextDate, err := usecase.NextDate(now, dateStr, repeat)
	if err != nil {
		log.WithFields(log.Fields{
			"date":   dateStr,
			"repeat": repeat,
		}).Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем следующую дату как строку
	c.String(http.StatusOK, nextDate)
}
