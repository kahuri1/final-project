package service

import (
	"github.com/kahuri1/final-project/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) TaskDone(task *model.Task) (bool, error) {
	_, err := s.repo.TaskDone(task)
	if err != nil {
		log.Errorf("failed to done task: %w", err)
		return false, err
	}
	return true, nil
}
