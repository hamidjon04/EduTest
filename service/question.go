package service

import (
	"context"
	"edutest/pkg/model"
	"fmt"
)

func (S *Service) CreateQuestion(ctx context.Context, req *model.CreateQuestionReq) error {
	err := S.Storage.Question().CreateQuestion(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is save data of question: %v", err))
		return err
	}
	return nil
}

func (S *Service) UpdateQuestion(ctx context.Context, req *model.UpdateQuestionReq) error {
	err := S.Storage.Question().UpdateQuestion(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is update data of question: %v", err))
		return err
	}
	return nil
}

func (S *Service) DeleteQuestion(ctx context.Context, id string) error {
	err := S.Storage.Question().DeleteQuestion(id)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is delete data of question: %v", err))
		return err
	}
	return nil
}

func (S *Service) GetQuestions(ctx context.Context, req *model.GetQuestionsReq)(*model.GetQuestionsResp, error){
	resp, err := S.Storage.Question().GetQuestions(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is get data of question: %v", err))
		return nil, err
	}
	return resp, nil
}
