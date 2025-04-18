package model

type Student struct {
	Id          string `json:"id"`
	StudentId   string `json:"student_id"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
	Subject1    string `json:"subject1"`
	Subject2    string `json:"subject2"`
}

type CreateStudentReq struct {
	StudentId   string `json:"-"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
	Subject1    string `json:"subject1"`
	Subject2    string `json:"subject2"`
}

type CreateStudentResp struct {
	StudentId string `json:"student_id"`
}

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type UpdateStudentReq struct {
	Id          string `json:"-"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	PhoneNumber string `json:"phone_number"`
	Subject1    string `json:"subject1"`
	Subject2    string `json:"subject2"`
}

type StudentId struct {
	Id string `json:"id"`
}

type GetStudentsResp struct {
	Students []Student `json:"students"`
}

type Status struct {
	Message string `json:"message"`
}

type Subject struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateSubjectReq struct {
	Name string `json:"name"`
}

type GetSubjectsResp struct {
	Subjects []Subject `json:"subjects"`
}

type GetStudentSubjectResp struct {
	Subject1 string `json:"subject1"`
	Subject2 string `json:"subject2"`
}

type Option struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
}

type CreateQuestionReq struct {
	SubjectId        string `json:"subject_id"`
	Type             string `json:"type"`
	QuestionText     string `json:"question_text"`
	Options          Option `json:"options"`
	Answer           string `json:"answer"`
	QuestionImageUrl string `json:"question_image_url"`
	OptionImageUrl   Option `json:"option_image_url"`
	AnswerImageUrl   string `json:"answer_image_url"`
}

type UpdateQuestionReq struct {
	Id               string `json:"-"`
	Type             string `json:"type"`
	QuestionText     string `json:"question_text"`
	Options          Option `json:"options"`
	Answer           string `json:"answer"`
	QuestionImageUrl string `json:"question_image_url"`
	OptionImageUrl   Option `json:"option)image_url"`
	AnswerImageUrl   string `json:"answer_image_url"`
}

type GetQuestionsReq struct {
	Id        string `json:"id"`
	SubjectId string `json:"subject_id"`
	Type      string `json:"type"`
}

type Question struct {
	Id               string `josn:"id"`
	SubjectId        string `json:"subject_id"`
	Type             string `json:"type"`
	QuestionText     string `json:"question_text"`
	Options          Option `json:"options"`
	Answer           string `json:"answer"`
	QuestionImageUrl string `json:"question_image_url"`
	OptionImageUrl   Option `json:"option_image_url"`
	AnswerImageUrl   string `json:"answer_image_url"`
}

type GetQuestionsResp struct {
	Questions []Question `json:"questions"`
}

type CreateTemplateReq struct {
	StudentId string `json:"student_id"`
	Day       string `json:"day"`
}

type GetTemplatesReq struct {
	StudentId string `json:"student_id"`
	Day       string `json:"day"`
}

type TemplateId struct {
	Id string `json:"id"`
}

type GetTemplatesResp struct {
	Templates []TemplateId `json:"templates"`
}

type TemplateQuestionsReq struct {
	SubjectId string `json:"subject_id"`
	Count     int    `json:"count"`
}

type CreateTemplateQuestionReq struct {
	TemplateId  string `json:"template_id"`
	QuesttionId string `json:"question_id"`
	Number      int    `json:"number"`
}

type CreateTemplateAnswer struct {
	TemplateId string         `json:"template_id"`
	Answers    map[int]string `json:"answers"`
}

type CreatePdf struct {
	StudentId  string     `josn:"student_id"`
	Name       string     `json:"name"`
	Lastname   string     `json:"lastname"`
	TemplateId string     `json:"template_id"`
	Subject1   string     `josn:"subject1"`
	Subject2   string     `json:"subject2"`
	Questions  []Question `json:"questions"`
}

type QuestionAnswer struct {
	Number int    `json:"number"`
	Answer string `json:"answer"`
}

type CheckStudentTestReq struct {
	StudentId string           `json:"student_id"`
	Day       string           `json:"day"`
	Answers   []QuestionAnswer `json:"answers"`
}

type Result struct {
	Correct   int     `json:"correct"`
	InCorrect int     `json:"incorrect"`
	Percent   float64 `json:"percent"`
}

type QuestionResult struct {
	Number int  `json:"number"`
	Status bool `json:"status"`
}

type CreateStudentResultReq struct {
	StudentId  string           `json:"student_id"`
	Results    []QuestionResult `json:"results"`
	Point      float64          `json:"point"`
	TemplateId string           `json:"template_id"`
}

type GetStudentResultReq struct {
	StudentId  string `json:"student_id"`
	TemplateId string `json:"template_id"`
}

type StudentResult struct {
	TemplateId string           `json:"template_id"`
	Result     []QuestionResult `json:"result"`
	Ball       float64          `json:"ball"`
}

type GetStudentResultResp struct {
	Results []StudentResult `json:"results"`
}

type GetStudentsResultReq struct {
	Day      string `json:"day"`
	Subject1 string `json:"subject1"`
	Subject2 string `json:"subject2"`
}

type StudentReslt struct {
	StudentId string           `json:"student_id"`
	Name      string           `json:"name"`
	Lastname  string           `json:"lastname"`
	Subject1  string           `json:"subject1"`
	Subject2  string           `json:"subject2"`
	Day       string           `json:"day"`
	Result    []QuestionResult `json:"result"`
	Ball      float64          `json:"ball"`
}

type GetStudentsResultResp struct {
	StudentsResults []StudentReslt `json:"students_results"`
	Count           int            `json:"count"`
}

type UpdateSubjectReq struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type StudentsStatus struct {
	Correct           int       `json:"correct"`
	Incorrect         int       `json:"incorrect"`
	IncorrectStudents []Student `json:"incorrect_students"`
}

type Incorrect struct{
	Nomer int `json:"nomer"`
	Name string `json:"name"`
}

type QuestionsStatus struct {
	Correct            int        `json:"correct"`
	Incorrect          int        `json:"incorrect"`
	IncorrectQuestions []Incorrect `json:"incorrect_questions"`
}
