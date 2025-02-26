package postgres

import (
	"database/sql"
	"edutest/pkg/model"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type SubjectRepo interface {
	CreateSubject(req *model.CreateSubjectReq) error 
	DeleteSubject(id string) error
	GetSubjects(id string) (*model.GetSubjectsResp, error)
	CreateStudentSubject(req *model.CreateStudentReq) error
	UpdateStudentSubject(req *model.UpdateStudentReq) error 
	DeleteStudentSubject(req *model.StudentId) error
}

type subjectImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewSubjectRepo(db *sql.DB, log *slog.Logger) SubjectRepo {
	return &subjectImpl{
		DB:  db,
		Log: log,
	}
}

func (S *subjectImpl) CreateSubject(req *model.CreateSubjectReq) error {
	id := uuid.NewString()
	query := `
				INSERT INTO subjects(
					id, name)
				VALUES
					($1, $2)`
	_, err := S.DB.Exec(query, id, req.Name)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is insert data of subject: %v", err))
		return err
	}
	return nil
}

func (S *subjectImpl) DeleteSubject(id string) error {
	query := `
				DELETE FROM 
					subjects
				WHERE 
					id = $1`
	_, err := S.DB.Exec(query, id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete subject: %v", err))
		return err
	}
	return nil
}

func (S *subjectImpl) GetSubjects(id string) (*model.GetSubjectsResp, error) {
	var subjects = []model.Subject{}
	query := `
				SELECT 
					id, name
				FROM 
					subjects`
	if len(id) != 0 {
		query += fmt.Sprintf(" WHERE id = '%s'", id)
	}
	rows, err := S.DB.Query(query)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get subjects: %v", err))
		return nil, err
	}
	for rows.Next() {
		var subject model.Subject
		err := rows.Scan(&subject.Id, &subject.Name)
		if err != nil {
			S.Log.Error(fmt.Sprintf("Error is scan data of subject: %v", err))
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return &model.GetSubjectsResp{
		Subjects: subjects,
	}, nil
}

func (S *subjectImpl) CreateStudentSubject(req *model.CreateStudentReq) error {
	query := `
				INSERT INTO student_subjects(
					student_id, subject1, subject2)
				VALUES
					($1, $2, $3)`
	_, err := S.DB.Exec(query, req.StudentId, req.Subject1, req.Subject2)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is insert student's subject data to database: %v", err))
		return err
	}
	return nil
}

func (S *subjectImpl) UpdateStudentSubject(req *model.UpdateStudentReq) error {
	query := `
				UPDATE student_subjects SET
					subject1 = $1, subject2 = $2
				WHERE
					student_id = $3`
	_, err := S.DB.Exec(query, req.Subject1, req.Subject2, req.Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is update student's subject data: %v", err))
		return err
	}
	return nil
}

func (S *subjectImpl) DeleteStudentSubject(req *model.StudentId) error {
	query := `
				DELETE
					student_id = $1
				FROM 
					student_subjects`
	_, err := S.DB.Exec(query, req.Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete student's subjects: %v", err))
		return err
	}
	return nil
}
