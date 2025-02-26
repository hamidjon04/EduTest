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
		if err != nil{
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
		Questions: questions,
	})
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is create pdf file: %v", err))
		return err
	}
	return nil
}

func(S *Service) CheckStudentTest(ctx context.Context, req *model.CheckStudentTestReq)(*model.Result, error){
	student, err := S.Storage.Student().GetStudentByStringId(req.StudentId)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is get student: %v", err))
		return nil, err
	}
	template, err := S.Storage.Template().GetTemplates(&model.GetTemplatesReq{
		StudentId: student.Id,
		Day: req.Day,
	})
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is get student's template: %v", err))
		return nil, err
	}
	S.Log.Info(fmt.Sprintf("%v", template))
	answers, err := S.Storage.Template().GetTemplateAnswer(template.Templates[0].Id)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is get answers: %v", err))
		return nil, err
	}

	var result model.Result
	for _, j := range req.Answers{
		if answers[j.Number] == j.Answer{
			result.Correct++
		}else{
			result.InCorrect++
		}
	}
	result.Percent = float64(result.Correct)/float64(len(answers))
	return &result, nil
}
