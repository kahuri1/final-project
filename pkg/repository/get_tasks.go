package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/final-project/pkg/model"
	"time"
)

func (r *Repository) GetTasks(search string) (model.TasksResp, error) {
	tasks := make([]model.Task, 0)

	query := sq.
		Select("id", "date", "title", "comment", "repeat").
		From("scheduler").
		OrderBy("date ASC").
		PlaceholderFormat(sq.Dollar)
	if search != "" {
		// Попытка распознать дату из строки поиска
		date, err := time.Parse("02.01.2006", search)
		if err == nil {
			// Если строка поиска совпадает с форматом даты, фильтруем по дате
			query = query.Where(sq.Eq{"date": date})
		} else {
			// Если это не дата, ищем в заголовке и комментарии
			query = query.Where(sq.Or{
				sq.Like{"title": "%" + search + "%"},
				sq.Like{"comment": "%" + search + "%"},
			})
		}
	}
	sql, args, err := query.ToSql()

	if err != nil {
		return model.TasksResp{Tasks: tasks}, fmt.Errorf("failed to create task selection request: %w", err)
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return model.TasksResp{Tasks: tasks}, fmt.Errorf("failed to execute task selection query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.Id, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return model.TasksResp{Tasks: tasks}, fmt.Errorf("failed to scan task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return model.TasksResp{Tasks: tasks}, fmt.Errorf("failed during rows iteration: %w", err)
	}

	return model.TasksResp{Tasks: tasks}, nil
}
