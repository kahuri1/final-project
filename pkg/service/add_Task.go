package service

import (
	"github.com/kahuri1/final-project/pkg/model"
	log "github.com/sirupsen/logrus"
)

func (s *Service) CreateTask(t model.Task) (int64, error) {

	id, err := s.repo.CreateDbTask(t)
	if err != nil {
		log.Errorf("failed to create message: %w", err)
	}

	//log.Infof("message with id %d created", id)
	return id, nil
}
