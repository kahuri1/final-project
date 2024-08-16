package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/final-project/pkg/model"
)

func (r *Repository) TaskDone(task *model.Task) (bool, error) {
	var existingTask model.Task
	err := r.db.Get(&existingTask, "SELECT * FROM scheduler WHERE id=$1", task.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, fmt.Errorf("task not found") // Ошибка, если задача не найдена
		}
		return false, err // Если возникла другая ошибка
	}
	sql, args, err := sq.Update("scheduler").
		Set("date", task.Date).
		Where(sq.Eq{"id": task.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("failed to create update query: %w", err)
	}
	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return false, fmt.Errorf("failed to execute update query: %w", err)
	}

	return true, nil
}
