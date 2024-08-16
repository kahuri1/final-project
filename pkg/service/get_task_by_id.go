package service

import (
	"github.com/kahuri1/final-project/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) GetTaskById(id int) (*model.Task, error) {

	task, err := s.repo.GetTaskById(id)
	if err != nil {
		log.Errorf("failed to get task: %d", err)
		return nil, err
	}
	return task, nil
}
