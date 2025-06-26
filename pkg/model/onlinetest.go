package model

type GetTest struct {
	UserId     string `json:"user_id"`
	Subject_Id string `json:"subject_id"`
	Count      int    `json:"count"`
}

type GetTestResp struct {
	TestId    string         `json:"test_id"`
	Questions []QuestionTest `json:"question"`
}

type CheckReq struct {
	TestId  string           `json:"test_id"`
	Answers []QuestionAnswer `json:"answers"`
}

type TestResult struct {
	TestId    string           `json:"test_id"`
	Results   []QuestionResult `json:"results"`
	Correct   int              `json:"correct"`
	Incorrect int              `json:"incorrect"`
}
