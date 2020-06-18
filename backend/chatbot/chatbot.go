package chatbot

import (
	"google.golang.org/api/chat/v1"
	"oss/dto"
	"oss/hangout"
	"oss/models"
	"oss/service"
	_ "oss/service"
	"oss/web"
	"strconv"
	_ "strconv"
	"strings"
	_ "strings"
)

const (
	SUBMIT_MULTI_ANSWER = "SUBMIT_MULTI_ANSWER"
)

func analysisSubmitAnswer (){

}

func submitAnswerProcessor (c *web.Context, content string) *chat.Message {
	tokens := strings.Split(content, "-")
	email := tokens[0]
	classQuizSetId := tokens[1]
	quizId := tokens[2]
	buttonIndex := tokens[3]

	parsedClassQuizSetId, err := strconv.ParseInt(classQuizSetId, 10, 64)
	parsedQuizId, err := strconv.ParseInt(quizId, 10, 64)
	if err != nil {
		return InternalServerError()
	}
	service.ScoreQuizzes(c,
		service.ScoringQueueIdent{
		ClassQuizSetId: parsedClassQuizSetId,
		Email:          email },
		[]*dto.QuizForScoring{
			{
				QuizId:     parsedQuizId,
				QuizType:   models.QUIZ_TYPE_MULTI,
				QuizAnswer: buttonIndex,
			},
	})
	return GetNextQuiz(c, email)
}

func ProcessChat(c *web.Context, event hangout.DquizDeprecatedEvent) *chat.Message {
	//email := event.User.Email
	//message := strings.Fields(event.Message.Text)
	if event.Type == "CARD_CLICKED" {
		action := event.Action.ActionMethodName
		content := action[strings.Index(action, ":") + 1:]
		methodPrefixPosition := strings.Index(action, ":")
		methodPrefix := action[0 : methodPrefixPosition]
		switch methodPrefix {
			case "SUBMIT_MULTI_ANSWER":
				return submitAnswerProcessor(c, content)
			case "SUBMIT_SHORT_ANSWER":
				return submitAnswerProcessor(c, content)
		}
	} else {

	}
	return nil
}
