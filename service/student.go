package service

import (
	"context"
	"edutest/pkg/model"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
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

func (S *Service) GetStudentsResult(ctx context.Context, req *model.GetStudentsResultReq) (*model.GetStudentsResultResp, error) {
	resp, err := S.Storage.Student().GetStudentsResult(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get student's results from database: %v", err))
		return nil, err
	}
	return resp, nil
}

func (S *Service) OpenStudentsExelFile(ctx context.Context, filePath string) (*model.StudentsStatus, error) {
	var resp model.StudentsStatus

	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
	    S.Log.Error(fmt.Sprintf("Faylni ochishda xatolik: %v", err))
		return nil, err
	}

	sheets := xlsx.GetSheetList()

	if len(sheets) == 0 {
	    S.Log.Info("Faylda hech qanday sahifa mavjud emas!")
		return nil, err
	}

	sheetName := sheets[0]
	S.Log.Info(fmt.Sprintf("Faylni ochish uchun ishlatiladigan sahifa: %v", sheetName))

	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
	    S.Log.Error(fmt.Sprintf("Sahifani o'qishda xatolik: %v", err))
		return nil, err
	}

	subjects, err := S.Storage.Subject().GetSubjects("")
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get subjects: %v", err))
	}
	var hmap = make(map[string]bool)
	for _, s := range subjects.Subjects {
		hmap[s.Id] = true
	}
	S.Log.Info(fmt.Sprintf("%v", hmap))	

	for _, row := range rows[1:] {
		if len(row) < 6 {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			continue
		}

		var student model.Student
		if len(row[1]) == 0 {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			student.Name = row[1]
		}

		if len(row[2]) == 0 {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			student.Lastname = row[2]
		}

		if len(row[3]) == 0 {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			student.PhoneNumber = row[3]
		}

		if len(row[4]) == 0 {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning ma'lumoti to'liq emas: %v", row))
			continue
		} else if !hmap[row[4]] {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning subject1 noto'g'ri kiritilgan: %v", row))
			continue
		} else {
			student.Subject1 = row[4]
		}

		if len(row[5]) == 0 {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning ma'lumoti to'liq emas: %v", row))
			continue
		} else if !hmap[row[4]] {
			resp.IncorrectStudents = addIncorrectStudentList(row, resp.IncorrectStudents)
			S.Log.Info(fmt.Sprintf("Studentning subject2 noto'g'ri kiritilgan: %v", row))
			continue
		} else {
			student.Subject2 = row[5]
		}

		studentId := S.Storage.Student().StudentCount() + 200001
		student.StudentId = strconv.Itoa(studentId)
		_, err := S.Storage.Student().CreateStudent(&model.CreateStudentReq{
			StudentId:   student.StudentId,
			Name:        student.Name,
			Lastname:    student.Lastname,
			PhoneNumber: student.PhoneNumber,
			Subject1:    student.Subject1,
			Subject2:    student.Subject2,
		})
		if err != nil {
			resp.IncorrectStudents = addIncorrectStudentList	(row, resp.IncorrectStudents)
			S.Log.Error(fmt.Sprintf("Error is save data to database: %v", err))
			resp.Incorrect = len(resp.IncorrectStudents)
			return &resp, err
		}

		resp.Correct++
	}
	resp.Incorrect = len(resp.IncorrectStudents)

	return &resp, nil
}

func addIncorrectStudentList(row []string, incorrectStudentList []model.Student) []model.Student {
	incorrectStudentList = append(incorrectStudentList, model.Student{
		Name:        row[1],
		Lastname:    row[2],
		PhoneNumber: row[3],
		Subject1:    row[4],
		Subject2:    row[5],
	})

	return incorrectStudentList
}
