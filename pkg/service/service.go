package service

import (
	"github.com/kahuri1/final-project/pkg/model"
	log "github.com/sirupsen/logrus"
)

type repo interface {
	CreateDbTask(task model.Task) (int64, error)
	GetTasks(search string) (model.TasksResp, error)
	GetTaskById(id int) (*model.Task, error)
	UpdateTask(task *model.Task) (bool, error)
	TaskDone(task *model.Task) (bool, error)
	DeleteTask(id int) error
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}
