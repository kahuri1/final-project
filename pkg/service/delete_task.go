package service

import (
	log "github.com/sirupsen/logrus"
)

func (s *Service) DeleteTask(id int) error {

	err := s.repo.DeleteTask(id)
	if err != nil {
		log.Errorf("failed delete task: %w", err)
	}
	return nil
}
