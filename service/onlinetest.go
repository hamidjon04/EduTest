package service

import (
	"edutest/pkg/function"
	"edutest/pkg/model"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) GetQuestionsForTest(req *model.GetTest) (*model.GetTestResp, error) {
	var tests = []model.Question{}
	id := uuid.NewString()

	questions, err := s.Storage.Question().GetQuestionForTemplate(&model.TemplateQuestionsReq{
		SubjectId: req.Subject_Id,
		Count:     req.Count,
	})
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is get questions: %v", err))
		return nil, err
	}

	for _, v := range questions.Questions {
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

	var question = []model.QuestionTest{}

	for i, v := range tests {
		question = append(question, model.QuestionTest{
			Id:               v.Id,
			Nomer:            i + 1,
			SubjectId:        v.SubjectId,
			Type:             v.Type,
			QuestionText:     v.QuestionText,
			Options:          v.Options,
			QuestionImageUrl: v.QuestionImageUrl,
			OptionImageUrl:   v.OptionImageUrl,
		})
	}

	return &model.GetTestResp{
		TestId:    id,
		Questions: question,
	}, nil
}

func (s *Service) CheckTest(req model.CheckReq) (*model.TestResult, error) {

	answers, err := s.Storage.Template().GetTemplateAnswer(req.TestId)
	if err != nil {
		s.Log.Error(fmt.Sprintf("Error is get answers: %v", err))
		return nil, err
	}

	var result model.Result
	var results []model.QuestionResult
	var numbers = make(map[int]bool)
	for _, j := range req.Answers {
		var questionResult = model.QuestionResult{
			Number: j.Number,
		}
		if answers[j.Number] == j.Answer {
			result.Correct++
			questionResult.Status = true
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

	return &model.TestResult{
		TestId:    req.TestId,
		Results:   results,
		Correct:   result.Correct,
		Incorrect: len(results)-result.Correct,
	}, nil
}
