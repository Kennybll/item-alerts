package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type DatabaseService struct {
	db *sqlx.DB
}

func NewDatabaseService() *DatabaseService {
	return &DatabaseService{}
}

func (s *DatabaseService) GetDb() *sqlx.DB {
	if s.db == nil {
		panic("Database not initialized")
	}
	return s.db
}

func (s *DatabaseService) Init() error {
	var err error
	s.db, err = sqlx.Connect("mysql", os.Getenv("DATABASE_URL"))
	return err
}

func (s *DatabaseService) Close() error {
	return s.GetDb().Close()
}
