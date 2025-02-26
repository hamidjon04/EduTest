package postgres

import (
	"database/sql"
	"edutest/pkg/model"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type TemplateRepo interface {
	CreateTempl(req *model.CreateTemplateReq) (string, error)
	DeleteTempl(id string) error
	GetTemplates(req *model.GetTemplatesReq) (*model.GetTemplatesResp, error)
	CreateTemplateQuestion(req *model.CreateTemplateQuestionReq) error
	CreateTemplateAnswer(req *model.CreateTemplateAnswer) error
	GetTemplateAnswer(templateId string)(map[int]string, error)
}

type tempImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewTemplateRepo(db *sql.DB, log *slog.Logger) TemplateRepo {
	return &tempImpl{
		DB:  db,
		Log: log,
	}
}

func (T *tempImpl) CreateTempl(req *model.CreateTemplateReq) (string, error) {
	id := uuid.NewString()
	query := `
				INSERT INTO templates(
					id, student_id, day)
				VALUES
					($1, $2, $3)`
	_, err := T.DB.Exec(query, id, req.StudentId, req.Day)
	if err != nil {
		T.Log.Error(fmt.Sprintf("Error is insert data of template: %v", err))
		return "", err
	}
	return id, nil
}

func (T *tempImpl) DeleteTempl(id string) error {
	query := `
				DELETE FROM 
					templates
				WHERE 
					id = $1`
	_, err := T.DB.Exec(query, id)
	if err != nil {
		T.Log.Error(fmt.Sprintf("Error is delete template: %v", err))
		return err
	}
	return nil
}

func (T *tempImpl) GetTemplates(req *model.GetTemplatesReq) (*model.GetTemplatesResp, error) {
	var templates []model.TemplateId
	query := `
				SELECT 
					id
				FROM 
					templates
				WHERE 
					TRUE `
	if len(req.StudentId) > 0 {
		query += fmt.Sprintf(" AND student_id = '%s'", req.StudentId)
	}
	if len(req.Day) > 0 {
		query += fmt.Sprintf(" AND day = '%s'", req.Day)
	}
	rows, err := T.DB.Query(query)
	if err != nil {
		T.Log.Error(fmt.Sprintf("Error is get data of template: %v", err))
		return nil, err
	}
	for rows.Next() {
		var template model.TemplateId
		err = rows.Scan(&template.Id)
		if err != nil {
			T.Log.Error(fmt.Sprintf("Error is scan data of template: %v", err))
			return nil, err
		}
		templates = append(templates, template)
	}
	return &model.GetTemplatesResp{
		Templates: templates,
	}, nil
}

func (T *tempImpl) CreateTemplateQuestion(req *model.CreateTemplateQuestionReq) error {
	query := `
				INSERT INTO template_questions(
					template_id, question_id, question_number)
				VALUES
					($1, $2, $3)`
	_, err := T.DB.Exec(query, req.TemplateId, req.QuesttionId, req.Number)
	if err != nil {
		T.Log.Error(fmt.Sprintf("Error is insert question of template: %v", err))
		return err
	}
	return nil
}

func (T *tempImpl) CreateTemplateAnswer(req *model.CreateTemplateAnswer) error {
	query := `
				INSERT INTO templte_answers(
					template_id, answer)
				VALUES
					($1, $2)`
	answers, err := json.Marshal(req.Answers)
	if err != nil {
		T.Log.Error(fmt.Sprintf("Error is covert from map to json answers: %v", err))
		return err
	}
	_, err = T.DB.Exec(query, req.TemplateId, answers)
	if err != nil {
		T.Log.Error(fmt.Sprintf("Error is insert answers: %v", err))
		return err
	}
	return nil
}

func (T *tempImpl) GetTemplateAnswer(templateId string)(map[int]string, error){
	var answers = make(map[int]string)
	var answer []byte
	query := `
				SELECT 
					answer
				FROM 
					templte_answers
				WHERE 
					template_id = $1`
	err := T.DB.QueryRow(query, templateId).Scan(&answer)
	if err != nil{
		T.Log.Error(fmt.Sprintf("Error is get answers of template: %v", err))
		return nil, err
	}
	err = json.Unmarshal(answer, &answers)
	if err != nil{
		T.Log.Error(fmt.Sprintf("Error is unmarshal answer: %v", err))
		return nil, err
	}
	return answers, nil
}