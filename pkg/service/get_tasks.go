package service

import (
	"github.com/kahuri1/final-project/pkg/model"
)

func (s *Service) GetTasks(search string) (model.TasksResp, error) {
	tasks, err := s.repo.GetTasks(search)
	if err != nil {
		return model.TasksResp{}, err
	}

	return tasks, nil
}
