package service

import (
	"context"
	"edutest/pkg/function"
	"edutest/pkg/model"
	"edutest/pkg/pdf"
	"fmt"
)

func (S *Service) CreateTemplate(ctx context.Context, req *model.CreateTemplateReq) error {
	id, err := S.Storage.Template().CreateTempl(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is save data of template at database: %v", err))
		return nil
	}
	student, err := S.Storage.Student().GetStudents(&model.StudentId{Id: req.StudentId})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is student's data: %v", err))
		return err
	}
	var questions []model.Question
	subject1, err := S.Storage.Question().GetQuestionForTemplate(&model.TemplateQuestionsReq{
		SubjectId: student.Students[0].Subject1,
		Count:     30,
	})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get questions of student's subject1: %v", err))
		return err
	}
	subject2, err := S.Storage.Question().GetQuestionForTemplate(&model.TemplateQuestionsReq{
		SubjectId: student.Students[0].Subject2,
		Count:     30,
	})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get questions of student's subject1: %v", err))
		return err
	}
	questions = append(questions, subject1.Questions...)
	questions = append(questions, subject2.Questions...)
	for i := 0; i < len(questions); i++ {
		err = S.Storage.Template().CreateTemplateQuestion(&model.CreateTemplateQuestionReq{
			TemplateId:  id,
			QuesttionId: questions[i].Id,
			Number:      i + 1,
		})
		if err != nil {
			S.Log.Error(fmt.Sprintf("Error is create question of template: %v", err))
			return nil
		}
	}

	questions, answers := function.RandomOptions(questions)

	err = S.Storage.Template().CreateTemplateAnswer(&model.CreateTemplateAnswer{
		TemplateId: id,
		Answers:    answers,
	})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is save answers of tamplate to database: %v", err))
		return err
	}

	_, err = pdf.CreateTestTemplate(model.CreatePdf{
		StudentId:  student.Students[0].Id,
		Name:       student.Students[0].Name,
		Lastname:   student.Students[0].Lastname,
		TemplateId: id,
		Subject1:   student.Students[0].Subject1,
		Subject2:   student.Students[0].Subject2,
		Questions:  questions,
	})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is create pdf file: %v", err))
		return err
	}
	return nil
}

func (S *Service) CheckStudentTest(ctx context.Context, req *model.CheckStudentTestReq) (*model.Result, error) {
	// student, err := S.Storage.Student().GetStudentByStringId(req.StudentId)
	// if err != nil {
	// 	S.Log.Error(fmt.Sprintf("Error is get student: %v", err))
	// 	return nil, err
	// }
	templateId, err := S.Storage.Template().GetTemplate(req.StudentId, req.Day)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get student's template: %v", err))
		return nil, err
	}
	answers, err := S.Storage.Template().GetTemplateAnswer(templateId)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get answers: %v", err))
		return nil, err
	}

	var result model.Result
	var results []model.QuestionResult
	var numbers = make(map[int]bool)
	var point float64
	for _, j := range req.Answers {
		var questionResult = model.QuestionResult{
			Number: j.Number,
		}
		if answers[j.Number] == j.Answer {
			result.Correct++
			questionResult.Status = true
			if j.Number <= 30 {
				point += 3.1
			} else if j.Number > 30 && j.Number <= 60 {
				point += 2.1
			}
		} else {
			result.InCorrect++
			questionResult.Status = false
		}
		results = append(results, questionResult)
		numbers[j.Number] = true
	}

	for i := 1; i <= len(answers); i++ {
		var questionResult = model.QuestionResult{
			Number: i,
		}
		if !numbers[i] {
			questionResult.Status = false
		}
		results = append(results, questionResult)
	}

	err = S.Storage.Student().CreateStudentResult(&model.CreateStudentResultReq{
		StudentId:  req.StudentId,
		TemplateId: templateId,
		Results:    results,
		Point:      point,
	})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is save student's result: %v", err))
	}

	result.Percent = float64(result.Correct) / float64(len(answers))
	return &result, nil
}

func (S *Service) GetStudentTemplates(ctx context.Context, req *model.GetTemplatesReq) (*[]byte, error) {
	templates, err := S.Storage.Template().GetTemplates(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get templates: %v", err))
		return nil, err
	}

	file, err := function.ReadPDFFile(templates.Templates[0].Id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get pdf of template: %v", err))
		return nil, err
	}
	return &file, nil
}

func (S *Service) GetStudentResult(ctx context.Context, req *model.GetStudentResultReq) (*model.GetStudentResultResp, error) {
	resp, err := S.Storage.Student().GetStudentResult(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get student's result: %v", err))
		return nil, err
	}
	return resp, nil
}
