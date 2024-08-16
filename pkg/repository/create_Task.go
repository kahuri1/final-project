package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/final-project/pkg/model"
)

func (r *Repository) CreateDbTask(task model.Task) (int64, error) {
	sql, args, err := sq.
		Insert("scheduler").
		Columns("date", "title", "comment", "repeat").
		Values(task.Date, task.Title, task.Comment, task.Repeat).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create task creation request: %w", err)
	}

	var id int64

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to proccess news create query: %w", err)
	}

	return id, nil
}
