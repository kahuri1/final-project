package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/final-project/pkg/model"
)

func (r *Repository) GetTaskById(id int) (*model.Task, error) {
	sql, args, err := sq.
		Select("id", "date", "title", "comment", "repeat").
		From("scheduler").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to create task query request: %w", err)
	}

	var task model.Task

	err = r.db.Get(&task, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to process task get query: %w", err)
	}

	return &task, nil
}
