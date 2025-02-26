package postgres

import (
	"database/sql"
	"edutest/pkg/model"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type QuestionRepo interface {
	CreateQuestion(req *model.CreateQuestionReq) error
	UpdateQuestion(req *model.UpdateQuestionReq) error
	DeleteQuestion(id string) error
	GetQuestions(req *model.GetQuestionsReq) (*model.GetQuestionsResp, error)
	GetQuestionForTemplate(req *model.TemplateQuestionsReq)(*model.GetQuestionsResp, error)
}

type questionImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewQuestionRepo(db *sql.DB, log *slog.Logger) QuestionRepo {
	return &questionImpl{
		DB:  db,
		Log: log,
	}
}

func (Q *questionImpl) CreateQuestion(req *model.CreateQuestionReq) error {
	id := uuid.NewString()
	query := `
				INSERT INTO questions(
					id, subject_id, question_text, options, answer, question_image, options_image, answer_image, type)
				VALUES
					($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	options, err := json.Marshal(req.Options)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is marshal options of question: %v", err))
		return err
	}
	optionsUrl, err := json.Marshal(req.OptionImageUrl)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is marshal optionsUrl of question: %v", err))
		return err
	}
	_, err = Q.DB.Exec(query, id, req.SubjectId, req.QuestionText, options, req.Answer, req.QuestionImageUrl, optionsUrl, req.AnswerImageUrl, req.Type)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is insert data of question: %v", err))
		return err
	}
	return nil
}

func (Q *questionImpl) UpdateQuestion(req *model.UpdateQuestionReq) error {
	query := `
				UPDATE questions SET 
					question_text = $1, options = $2, answer = $3, question_image = $4, options_image = $5, answer_image = $6, type = $7
				WHERE 
					id = $8`
	options, err := json.Marshal(req.Options)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is marshal options of question: %v", err))
		return err
	}
	optionsUrl, err := json.Marshal(req.OptionImageUrl)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is marshal optionsUrl of question: %v", err))
		return err
	}
	_, err = Q.DB.Exec(query, req.QuestionText, options, req.Answer, req.QuestionImageUrl, optionsUrl, req.AnswerImageUrl, req.Type, req.Id)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is update data of question: %v", err))
		return err
	}
	return nil
}

func (Q *questionImpl) DeleteQuestion(id string) error {
	query := `
				UPDATE questions SET
					deleted_at = NOW()
				WHERE 
					id = $1`
	_, err := Q.DB.Exec(query, id)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is delete question: %v", err))
		return err
	}
	return nil
}

func (Q *questionImpl) GetQuestions(req *model.GetQuestionsReq) (*model.GetQuestionsResp, error) {
	var questions []model.Question
	query := `
				SELECT 
					id, subject_id, type, question_text, options, answer, question_image, options_image, answer_image
				FROM 
					questions
				WHERE 
					deleted_at IS NULL`
	if len(req.Id) > 0 {
		query += fmt.Sprintf(" AND id = '%s'", req.Id)
	}
	if len(req.SubjectId) > 0 {
		query += fmt.Sprintf(" AND subject_id = '%s'", req.SubjectId)
	}
	if len(req.Type) > 0 {
		query += fmt.Sprintf(" AND type = '%s'", req.Type)
	}
	rows, err := Q.DB.Query(query)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is get questions: %v", err))
		return nil, err
	}
	var options, optionsUrl []byte
	for rows.Next() {
		var q model.Question
		err = rows.Scan(&q.Id, &q.SubjectId, &q.Type, &q.QuestionText, &options, &q.Answer, &q.QuestionImageUrl, &optionsUrl, &q.AnswerImageUrl)
		if err != nil {
			Q.Log.Error(fmt.Sprintf("Error is scan question: %v", err))
			return nil, err
		}
		err = json.Unmarshal(options, &q.Options)
		if err != nil{
			Q.Log.Error(fmt.Sprintf("Error is unmarshal options: %v", err))
			return nil, err
		}
		err = json.Unmarshal(optionsUrl, &q.OptionImageUrl)
		if err != nil{
			Q.Log.Error(fmt.Sprintf("Error is unmaeshal optionsUrl: %v", err))
			return nil, err
		}
		questions = append(questions, q)
	}
	return &model.GetQuestionsResp{
		Questions: questions,
	}, nil
}

func(Q *questionImpl) GetQuestionForTemplate(req *model.TemplateQuestionsReq)(*model.GetQuestionsResp, error){
	var questions []model.Question
	query := `
				SELECT 
					id, subject_id, type, question_text, options, answer, question_image, options_image, answer_image
				FROM 
					questions
				WHERE 
					deleted_at IS NULL AND subject_id = $1
				ORDER BY
					RANDOM()
				LIMIT $2`
	rows, err := Q.DB.Query(query, req.SubjectId, req.Count)
	if err != nil {
		Q.Log.Error(fmt.Sprintf("Error is get questions: %v", err))
		return nil, err
	}
	var options, optionsUrl []byte
	for rows.Next() {
		var q model.Question
		err = rows.Scan(&q.Id, &q.SubjectId, &q.Type, &q.QuestionText, &options, &q.Answer, &q.QuestionImageUrl, &optionsUrl, &q.AnswerImageUrl)
		if err != nil {
			Q.Log.Error(fmt.Sprintf("Error is scan question: %v", err))
			return nil, err
		}
		err = json.Unmarshal(options, &q.Options)
		if err != nil{
			Q.Log.Error(fmt.Sprintf("Error is unmarshal options: %v", err))
			return nil, err
		}
		err = json.Unmarshal(optionsUrl, &q.OptionImageUrl)
		if err != nil{
			Q.Log.Error(fmt.Sprintf("Error is unmaeshal optionsUrl: %v", err))
			return nil, err
		}
		questions = append(questions, q)
	}
	return &model.GetQuestionsResp{
		Questions: questions,
	}, nil
}