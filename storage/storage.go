package storage

import (
	"database/sql"
	"edutest/storage/postgres"
	"log/slog"
)

type Storage interface {
	Student() postgres.StudentRepo
	Subject() postgres.SubjectRepo
	Question() postgres.QuestionRepo
	Template() postgres.TemplateRepo
}

type storageImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewStorage(db *sql.DB, log *slog.Logger) Storage {
	return &storageImpl{
		DB:  db,
		Log: log,
	}
}

func (s *storageImpl) Student() postgres.StudentRepo {
	return postgres.NewStudentRepo(s.DB, s.Log)
}

func (s *storageImpl) Subject() postgres.SubjectRepo {
	return postgres.NewSubjectRepo(s.DB, s.Log)
}

func (s *storageImpl) Question() postgres.QuestionRepo{
	return postgres.NewQuestionRepo(s.DB, s.Log)
}

func (s *storageImpl) Template() postgres.TemplateRepo{
	return postgres.NewTemplateRepo(s.DB, s.Log)
}
