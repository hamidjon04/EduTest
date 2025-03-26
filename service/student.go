package service

import (
	"context"
	"edutest/pkg/model"
	"fmt"
	"strconv"
)

func (S *Service) CreateStudent(ctx context.Context, req *model.CreateStudentReq) (*model.CreateStudentResp, error) {
	studentId := S.Storage.Student().StudentCount() + 200001
	req.StudentId = strconv.Itoa(studentId)
	resp, err := S.Storage.Student().CreateStudent(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is save data to database: %v", err))
		return nil, err
	}
	return &model.CreateStudentResp{
		StudentId: resp.StudentId,
	}, nil
}

func (S *Service) UpdateStudent(ctx context.Context, req *model.UpdateStudentReq) error {
	err := S.Storage.Student().UpdateStudent(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is update student's data at database: %v", err))
		return err
	}
	return nil
}

func (S *Service) DeleteStudent(ctx context.Context, req *model.StudentId) error {
	err := S.Storage.Student().DeleteStudent(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete student's data at database: %v", err))
		return err
	}
	return nil
}

func (S *Service) GetStudents(ctx context.Context, req *model.StudentId) (*model.GetStudentsResp, error) {
	resp, err := S.Storage.Student().GetStudents(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get student's data from database: %v", err))
		return nil, err
	}
	return resp, nil
}

func (S *Service) GetStudentsResult(ctx context.Context, req *model.GetStudentsResultReq) (*model.GetStudentsResultResp, error){
	resp, err := S.Storage.Student().GetStudentsResult(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is get student's results from database: %v", err))
		return nil, err
	}	
	return resp, nil
}
