package postgres

import (
	"database/sql"
	"edutest/pkg/model"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type StudentRepo interface {
	CreateStudent(req *model.CreateStudentReq) (*model.CreateStudentResp, error)
	UpdateStudent(req *model.UpdateStudentReq) error
	DeleteStudent(req *model.StudentId) error
	GetStudents(req *model.StudentId) (*model.GetStudentsResp, error)
	GetStudentByStringId(id string) (*model.Student, error)
	CreateStudentResult(req *model.CreateStudentResultReq) error
	GetStudentResult(req *model.GetStudentResultReq) (*model.GetStudentResultResp, error) 
}

type studentImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewStudentRepo(db *sql.DB, log *slog.Logger) StudentRepo {
	return &studentImpl{
		DB:  db,
		Log: log,
	}
}

func (S *studentImpl) CreateStudent(req *model.CreateStudentReq) (*model.CreateStudentResp, error) {
	id := uuid.NewString()

	tr, err := S.DB.Begin()
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is begin transaction: %v", err))
		return nil, err
	}
	defer tr.Commit()

	query := `
				INSERT INTO students(
					id, student_id, name, lastname, phone_number)
				VALUES
					($1, $2, $3, $4, $5)`
	_, err = tr.Exec(query, id, req.StudentId, req.Name, req.Lastname, req.PhoneNumber)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is insert student's data to database: %v", err))
		tr.Rollback()
		return nil, err
	}
	query = `
				INSERT INTO student_subjects(
					student_id, subject1, subject2)
				VALUES
					($1, $2, $3)`
	_, err = tr.Exec(query, id, req.Subject1, req.Subject2)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is insert student's subject data to database: %v", err))
		tr.Rollback()
		return nil, err
	}
	return &model.CreateStudentResp{
		StudentId: req.StudentId,
	}, nil
}

func (S *studentImpl) UpdateStudent(req *model.UpdateStudentReq) error {
	tr, err := S.DB.Begin()
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is begin transaction: %v", err))
		return err
	}
	defer tr.Commit()

	query := `
				UPDATE students SET
					name = $1, lastname = $2, phone_number = $3
				WHERE
					id = $4`
	_, err = tr.Exec(query, req.Name, req.Lastname, req.PhoneNumber, req.Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is update student's data: %v", err))
		tr.Rollback()
		return err
	}
	query = `
				UPDATE student_subjects SET
					subject1 = $1, subject2 = $2
				WHERE
					student_id = $3`
	_, err = tr.Exec(query, req.Subject1, req.Subject2, req.Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is update student's subject data: %v", err))
		tr.Rollback()
		return err
	}
	return nil
}

func (S *studentImpl) DeleteStudent(req *model.StudentId) error {
	tr, err := S.DB.Begin()
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is begin transaction: %v", err))
		return err
	}
	defer tr.Commit()

	query := `
				UPDATE students SET
					deleted_at = NOW()
				WHERE
					id = $1`
	_, err = tr.Exec(query, req.Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete student: %v", err))
		tr.Rollback()
		return err
	}
	query = `
				DELETE FROM 
					student_subjects
				WHERE
					student_id = $1`
	_, err = tr.Exec(query, req.Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete student's subjects: %v", err))
		tr.Rollback()
		return err
	}
	return nil
}

func (S *studentImpl) GetStudents(req *model.StudentId) (*model.GetStudentsResp, error) {
	var students = []model.Student{}
	query := `
				SELECT 
					id, student_id, name, lastname, phone_number
				FROM 
					students`
	if len(req.Id) != 0 {
		query += fmt.Sprintf(" WHERE id = '%s'", req.Id)
	}
	rows, err := S.DB.Query(query)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get student's data: %v", err))
		return nil, err
	}
	for rows.Next() {
		var student model.Student
		query := `
					SELECT 
						subject1, subject2
					FROM 
						student_subjects
					WHERE
						student_id = $1`
		err := rows.Scan(&student.Id, &student.StudentId, &student.Name, &student.Lastname, &student.PhoneNumber)
		if err != nil {
			S.Log.Error(fmt.Sprintf("Error is scan studnt's data: %v", err))
			return nil, err
		}
		err = S.DB.QueryRow(query, student.Id).Scan(&student.Subject1, &student.Subject2)
		if err != nil {
			S.Log.Error(fmt.Sprintf("Error is scan student's subjects: %v", err))
			return nil, err
		}
		students = append(students, student)
	}
	return &model.GetStudentsResp{
		Students: students,
	}, nil
}

func (S *studentImpl) GetStudentByStringId(id string) (*model.Student, error) {
	var student model.Student
	query := `
				SELECT 
					id, name, lastname, phone_number
				FROM 
					students
				WHERE 
					student_id = $1`
	err := S.DB.QueryRow(query, id).Scan(&student.Id, &student.Name, &student.Lastname, &student.PhoneNumber)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get student by studentId: %v", err))
		return nil, err
	}
	return &student, nil
}

func (S *studentImpl) CreateStudentResult(req *model.CreateStudentResultReq) error {
	result, err := json.Marshal(req.Results)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is marshal results: %v", err))
		return err
	}

	query := `
				INSERT INTO students_result(
					student_id, template_id, result, ball)
				VALUES
					($1, $2, $3, $4)`
	_, err = S.DB.Exec(query, req.StudentId, req.TemplateId, result, req.Point)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is insert student's result: %v", err))
		return err
	}
	return nil
}

func (S *studentImpl) GetStudentResult(req *model.GetStudentResultReq) (*model.GetStudentResultResp, error) {
	query := `
				SELECT
					template_id, result, ball
				FROM 
					students_result
				WHERE
					student_id = $1`
	if len(req.TemplateId) > 0{
		query += fmt.Sprintf(" AND template_id = '%v'", req.TemplateId)
	}
	
	rows, err := S.DB.Query(query, req.StudentId)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is get student's results: %v", err))
		return nil, err
	}
	defer rows.Close()

	var results []model.StudentResult
	for rows.Next() {
		var result model.StudentResult
		var subjects []byte
		err = rows.Scan(&result.TemplateId, &subjects, &result.Ball)
		if err != nil{
			S.Log.Error(fmt.Sprintf("Error is scan results: %v", err))
			return nil, err
		}
		err = json.Unmarshal(subjects, &result.Result)
		if err != nil{
			S.Log.Error(fmt.Sprintf("Error is unmarshaling result: %v", err))
			return nil, err
		}
		results = append(results, result)
	}
	return &model.GetStudentResultResp{
		Results: results,
	}, nil
}
