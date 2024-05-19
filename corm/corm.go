package corm

import (
	"database/sql"

	"github.com/HsiaoCz/something/corm/logger"
	"github.com/HsiaoCz/something/corm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		logger.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		logger.Error(err)
		return
	}
	e = &Engine{db: db}
	logger.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		logger.Error("Failed to close database")
	}
	logger.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
