package service

import (
	"edutest/pkg/function"
	"edutest/pkg/model"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (s *Service) GetQuestionsForTest(req *model.GetTest)(*model.GetTestResp, error){
	var tests = []model.Question{}
	id := uuid.NewString()

	questions, err := s.Storage.Question().GetQuestionForTemplate(&model.TemplateQuestionsReq{
		SubjectId: req.Subject_Id,
		Count: req.Count,
	})
	if err != nil{
		s.Log.Error(fmt.Sprintf("Error is get questions: %v", err))
		return nil, err
	}

	for _, v := range questions.Questions{
		tests = append(tests, v)
	}

	tests, answers := function.RandomOptions(tests)

	err = s.Storage.Template().CreateTemplateAnswer(&model.CreateTemplateAnswer{
		TemplateId: id,
		Answers:    answers,
	})
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is save answers of tamplate to database: %v", err))
		return nil, err
	}

	return &model.GetTestResp{
		TestId: id,
		Questions: tests,
	}, nil
}

func (s *Service) CheckTest(req model.CheckReq) (*model.TestResult, error){
	
	answers, err := s.Storage.Template().GetTemplateAnswer(req.TemplateId)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is get answers: %v", err))
		return nil, err
	}
	var uncorrect = len(answers)
	var correct int

	for _, v := range req.Answers{
		if strings.ToUpper(answers[v.Number]) == strings.ToUpper(v.Answer){
			correct++
			uncorrect--
		}
	}

	return &model.TestResult{
		Correct: correct,
		Uncorrect: uncorrect,
	}, nil
}