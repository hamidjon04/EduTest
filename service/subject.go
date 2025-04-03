package service

import (
	"context"
	"edutest/pkg/model"
	"fmt"
)

func (S *Service) CreateSubject(ctx context.Context, req *model.CreateSubjectReq) error {
	err := S.Storage.Subject().CreateSubject(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is save data of subject ad database: %v", err))
		return err
	}
	return nil
}

func (S *Service) UpdateSubject(ctx context.Context, req *model.UpdateSubjectReq) error {
	err := S.Storage.Subject().UpdateSubject(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete subject: %v", err))
		return err
	}
	return nil
}

func (S *Service) GetSubjects(ctx context.Context, id string) (*model.GetSubjectsResp, error) {
	resp, err := S.Storage.Subject().GetSubjects(id)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get subjects from database: %v", err))
		return nil, err
	}
	return resp, nil
}
