package repository

import (
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) DeleteTask(id int) error {
	sql, args, err := sq.Delete("scheduler").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}
