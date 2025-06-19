package model

type GetTest struct {
	UserId     string `json:"user_id"`
	Subject_Id string `json:"subject_id"`
	Count      int    `json:"count"`
}

type GetTestResp struct{
	TestId string `json:"test_id"`
	Questions []Question `json:"question"`
}

type CheckReq struct{
	TemplateId string `json:"template_id"`
	Answers []QuestionAnswer `json:"answers"`
}

type TestResult struct{
	Correct int `json:"correct"`
	Uncorrect int `json:"uncorrect"`
}