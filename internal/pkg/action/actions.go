package action

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"os"
	"tew4fs/trivia-backend/internal/pkg/constant"
	"tew4fs/trivia-backend/internal/pkg/log"
	"tew4fs/trivia-backend/internal/pkg/model"
	"tew4fs/trivia-backend/internal/pkg/utils"

	"go.uber.org/zap"
)

var (
	logger               = log.GetLogger()
	questionListFilePath = "data/questions.json"
	availableQuestions   = getListOfQuestions()
)

func GetQuestion() *model.Question {
	if len(availableQuestions) <= 0 {
		return &model.Question{
			Question: "No More Questions",
		}
	}

	index := rand.Intn(len(availableQuestions))
	q := availableQuestions[index]
	availableQuestions = utils.Remove(availableQuestions, index)

	logger.Info("got question from available questions",
		zap.String(constant.QuestionLogKey, q.Question),
		zap.Int(constant.QuestionListSizeLogKey, len(availableQuestions)))

	return q
}

func getListOfQuestions() []*model.Question {
	f, err := os.Open(questionListFilePath)
	if err != nil {
		logger.Error("questions file could not be opened", zap.Error(err), zap.String(constant.FilePathLogKey, questionListFilePath))
	}
	defer f.Close()

	questions := []*model.Question{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		q := model.Question{}
		err := json.Unmarshal([]byte(scanner.Text()), &q)
		if err != nil {
			logger.Error("could not read question from file", zap.Error(err))
		}

		questions = append(questions, &q) // Maybe use copy here
	}

	return questions
}

func CheckAnswer(q model.Question, answer string) bool {
	return q.Answer == answer
}
