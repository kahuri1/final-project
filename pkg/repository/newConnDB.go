package repository

import (
	"github.com/jmoiron/sqlx"
	migration "github.com/kahuri1/final-project/migrations"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
	"os"
)

type Config struct {
	DBName string
}

func NewSqlLiteDB(cfg Config) (*sqlx.DB, error) {
	var install bool
	_, err := os.Stat(cfg.DBName)
	if err != nil {
		install = true
	}

	if install {
		_, err = os.Create(cfg.DBName)
		if err != nil {
			return nil, err
		}
	}

	db, err := sqlx.Open("sqlite", cfg.DBName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(migration.Schema)
	if err != nil {
		logrus.Fatalln("Failed to create table:", err)
	}

	return db, nil
}
