package storage

import (
	"database/sql"
	"edutest/pkg/config"
	"edutest/storage/minio"
	"edutest/storage/postgres"
	"log/slog"
)

type Storage interface {
	Student() postgres.StudentRepo
	Subject() postgres.SubjectRepo
	Question() postgres.QuestionRepo
	Template() postgres.TemplateRepo
	Minio() minio.Minio
}

type storageImpl struct {
	DB  *sql.DB
	Log *slog.Logger
	Min *minio.MinioStruct
	Cfg config.Config
}

func NewStorage(db *sql.DB, log *slog.Logger, min *minio.MinioStruct, cfg config.Config) Storage {
	return &storageImpl{
		DB:  db,
		Log: log,
		Min: min,
		Cfg: cfg,
	}
}

func (s *storageImpl) Student() postgres.StudentRepo {
	return postgres.NewStudentRepo(s.DB, s.Log)
}

func (s *storageImpl) Subject() postgres.SubjectRepo {
	return postgres.NewSubjectRepo(s.DB, s.Log)
}

func (s *storageImpl) Question() postgres.QuestionRepo {
	return postgres.NewQuestionRepo(s.DB, s.Log)
}

func (s *storageImpl) Template() postgres.TemplateRepo {
	return postgres.NewTemplateRepo(s.DB, s.Log)
}

func (s *storageImpl) Minio() minio.Minio {
	return minio.MinioConnect(*s.Min, s.Log, s.Cfg)
}
