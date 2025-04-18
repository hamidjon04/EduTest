package service

import (
	"context"
	"edutest/pkg/model"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
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
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is delete data of question: %v", err))
		return err
	}
	return nil
}

func (S *Service) GetQuestions(ctx context.Context, req *model.GetQuestionsReq) (*model.GetQuestionsResp, error) {
	resp, err := S.Storage.Question().GetQuestions(req)
	if err != nil {
		S.Log.Error(fmt.Sprintf("Error is get data of question: %v", err))
		return nil, err
	}
	return resp, nil
}

func (S *Service) OpenQuestionsExelFile(ctx context.Context, filePath string) (*model.QuestionsStatus, error) {
	var resp model.QuestionsStatus

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
		if len(row) < 8 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Question ma'lumoti to'liq emas: %v", row))
			continue
		}

		var question model.CreateQuestionReq
		var option model.Option
		var optionImageUrl model.Option

		if len(row[1]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else if !hmap[row[1]] {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Question subjecti noto'g'ri kiritilgan: %v", row))
			continue
		} else {
			question.SubjectId = row[1]
		}

		if len(row[2]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			question.QuestionText = row[2]
		}

		if len(row[3]) == 0 {
			question.QuestionImageUrl = ""
		} else {
			question.QuestionImageUrl = row[3]
		}

		if len(row[4]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			option.A = row[4]
		}

		if len(row[5]) == 0 {
			optionImageUrl.A = ""
		} else {
			optionImageUrl.A = row[5]
		}

		if len(row[6]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			option.B = row[6]
		}

		if len(row[7]) == 0 {
			optionImageUrl.B = ""
		} else {
			optionImageUrl.B = row[7]
		}

		if len(row[8]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			option.C = row[8]
		}

		if len(row[9]) == 0 {
			optionImageUrl.C = ""
		} else {
			optionImageUrl.C = row[9]
		}

		if len(row[10]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			option.D = row[10]
		}

		if len(row[11]) == 0 {
			optionImageUrl.D = ""
		} else {
			optionImageUrl.D = row[11]
		}

		if len(row[12]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			question.Answer = row[12]
		}

		if len(row[13]) == 0 {
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Info(fmt.Sprintf("Questionning ma'lumoti to'liq emas: %v", row))
			continue
		} else {
			question.Type = row[13]
		}

		question.Options = option
		question.OptionImageUrl = optionImageUrl

		err := S.Storage.Question().CreateQuestion(&question)
		if err != nil{
			resp.IncorrectQuestions = addIncorrectQuestionList(row, resp.IncorrectQuestions)
			S.Log.Error(fmt.Sprintf("Error is save data to database: %v", err))
			resp.Incorrect = len(resp.IncorrectQuestions)
			return &resp, err
		}

		resp.Correct++
	}

	resp.Incorrect = len(resp.IncorrectQuestions)

	return &resp, nil
}

func addIncorrectQuestionList(row []string, incorrectQuestionList []model.Incorrect) []model.Incorrect {
	nomer, err := strconv.Atoi(row[0])
	if err != nil {
		fmt.Printf("Savol raqami noto'g'ri kiritilgan: %v", err)
		return nil
	}
	incorrectQuestionList = append(incorrectQuestionList, model.Incorrect{
		Nomer: nomer,
		Name:  row[2],
	})

	return incorrectQuestionList
}
