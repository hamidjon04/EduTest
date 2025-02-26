package function

import (
	"edutest/pkg/model"
	"math/rand"
	"time"
)

func RandomOptions(questions []model.Question) ([]model.Question, map[int]string) {
	var answers = make(map[int]string)

	for i := 0; i < len(questions); i++ {
		var options = []string{questions[i].Options.A, questions[i].Options.B, questions[i].Options.C, questions[i].Options.D}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		r.Shuffle(len(options), func(i, j int) {
			options[i], options[j] = options[j], options[i]
		})

		questions[i].Options.A = options[0]
		questions[i].Options.B = options[1]
		questions[i].Options.C = options[2]
		questions[i].Options.D = options[3]

		for j, v := range options {
			if v == questions[i].Answer {
				answers[i+1] = string(rune(65 + j))
				break
			}
		}
	}

	return questions, answers
}
