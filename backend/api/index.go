package api

import (
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/chat/v1"
	"log"
	"oss/cli"
	"oss/dto"
	"oss/service"
	"oss/web"
	"strings"
)

const (
	UNAUTHORIZED = "UNAUTHORIZED"
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	INVALID_REQUEST_HEADER = "INVALID_REQUEST_HEADER"
	INVALID_PATH_PARAMETER = "INVALID_PATH_PARAMETER"
	INVALID_REQUEST_BODY = "INVALID_BODY"
	RETRY = "RETRY"
)

var (
	USER = false
	ADMIN = true
)


var str string = `{
"type": "MESSAGE",
"eventTime": "2020-05-15T12:17:01.210175Z",
"message": {
"name": "spaces/4d0pRgAAAAE/messages/Ao9gFyA-WRE.Ao9gFyA-WRE",
"sender": {
"name": "users/113080763144266498882",
"displayName": "김현준",
"avatarUrl": "https://lh3.googleusercontent.com/-dHC1yj6-5w0/AAAAAAAAAAI/AAAAAAAAAAA/AMZuucm58FTrUFdPCg2MLu9C
KKm2uJr8ew/photo.jpg",
"email": "cosmos96@dgu.ac.kr",
"type": "HUMAN",
"domainId": "2ar8rqj"
},
"createTime": "2020-05-15T12:17:01.210175Z",
"text": "응시 -학번 2015112391 -과목코드 CSE-2017-02 -퀴즈코드 01",
"thread": {
"name": "spaces/4d0pRgAAAAE/threads/Ao9gFyA-WRE",
"retentionSettings": {
"state": "PERMANENT"
}
},
"space": {
"name": "spaces/4d0pRgAAAAE",
"type": "DM",
"singleUserBotDm": true
},
"argumentText": "we"
},
"user": {
"name": "users/113080763144266498882",
"displayName": "김현준",
"avatarUrl": "https://lh3.googleusercontent.com/-dHC1yj6-5w0/AAAAAAAAAAI/AAAAAAAAAAA/AMZuucm58FTrUFdPCg2MLu9CKK
m2uJr8ew/photo.jpg",
"email": "cosmos96@dgu.ac.kr",
"type": "HUMAN",
"domainId": "2ar8rqj"
},
"space": {
"name": "spaces/4d0pRgAAAAE",
"type": "DM",
"singleUserBotDm": true
},
"configCompleteRedirectUrl": "https://chat.google.com/api/bot_config_complete?token\u003dAAJCfVWFsUHCqrGY2f5MJTWe
WFZ0dY4-gZJz_QYtInoQ5aFEC8iemHkMvLENgYwEcHV6dKncQQQxjUWZ3eZ1z2h3KnWVC-9F3vF6OkmQHDbXZ2KzZUFb9Kr2UpzyxeTO1nWOVHpkyqb
I7kgQcp8-"
}`

var StartTestCommand cli.Command = cli.Command{
	Cmd: "퀴즈응시",
	MaxParam: 4,
	MinParam: 2,
	MustParam: []int{0, 1},
	Options: map[string]int{
		"학번" : 0,
		"과목코드" : 1,
		"퀴즈코드" : 2,
	},
	Entry: cli.StartTestEntry,
}

var identityKey = "UserName"

var App *web.Context

func InitRouters(context *web.Context) {
	App = context
	initPublicRouters(context)
	authMiddleware, err := getAuthMiddleware()
	classMiddleware, err := getClassAuthMiddleware()
	quizsetsMiddleware, err := getQuizsetsAuthMiddleware()

	r := context.Engine
	r.Use(CORSMiddleWare())

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	r.GET("/", func (c *gin.Context) {
		c.JSON(200,"why???")
	})

	schdch := make(chan service.SubmittedQuiz)
	go service.ScheduleQuizSetScoring(web.Context0.Redis, schdch)
	r.POST("/", func (c *gin.Context) {
		str = strings.Replace(str, "\n", "", -1)
		//println(str)
		var obj chat.DeprecatedEvent
		err := json.Unmarshal([]byte(str), &obj)
		s := strings.Fields(obj.Message.Text)

		for i :=0; i < len(s); i++ {
		//	println(s[i])
		}
		schdch<- service.SubmittedQuiz{
			ScoringQueueIdent : &service.ScoringQueueIdent {
				ClassQuizSetId: 4444,
				Email: "4whomtbts@gmail.com",
			},
			QuizForScoring : &dto.QuizForScoring {
				QuizId: 4444,
				QuizType: "MULTI",
				QuizAnswer: "1",
			},
		}
		response, commandErr := StartTestCommand.ProcessCommand(s)

		if err != nil {
			c.JSON(commandErr.StatusCode, commandErr.DetailedError)
			return
		}

		c.JSON(200, response)
	})

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.POST("/login", Login(context, authMiddleware))
	r.POST("/register", Register(context))

	class := r.Group("/class")
	class.Use(CORSMiddleWare())
	class.Use(classMiddleware.MiddlewareFunc())
	class.POST("/", CreateClass(context))
	class.POST("/enroll/:classcode", JoinClass(context));

	auth := r.Group("/user")
	auth.Use(CORSMiddleWare())
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())

	auth.GET("/classes/enrolled/:userName", GetAllEnrolledClass(context))
	auth.GET("/classes/managing/:userName", 
		GetAllManagingClass(context))

	r.GET("/users/:name", GetUserByUserName(context))
	quizset := r.Group("/quizsets")
	quizset.Use(CORSMiddleWare())
	quizset.Use(quizsetsMiddleware.MiddlewareFunc())
	quizset.POST("/quizset", CreateQuizSet(context))
	// classQuizSetId 에 해당하는 퀴즈 셋을 가져온다
	quizset.GET("/class/:classQuizSetId", GetQuizSetByClassQuizSetId(context))
	quizset.POST("/score", ScoreQuizzes(context))
	quizset.DELETE("/quizset/:quizsetId", DeleteQuizSet(context))
	quizset.POST("/quizset/:quizsetId/quiz/", AddQuiz(context))
	quizset.DELETE("/quizset/:quizsetId/quiz/:quizId", DeleteQuizFromQuizSet(context))
	quizset.POST("/quizset/:quizsetId/class/:classCode", LoadQuizSetToClass(context))
	quizset.DELETE("/quizset/:quizsetId/class/:classCode", DeleteQuizSetFromClass(context))
	quizset.POST("/classes/:classCode", GetQuizSetsOfClass(context))
	quizset.GET("/users/:username", GetUserQuizSet(context))

	quiz := r.Group("/quiz")
	quiz.Use(CORSMiddleWare())
	quiz.Use(authMiddleware.MiddlewareFunc())
	quiz.POST("/", CreateQuizSet(context))
	quiz.DELETE("/", DeleteQuizSet(context))

	//r.POST("/:", ScoreQuizzes(context))
}

