// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/class": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "강의개설",
                "parameters": [
                    {
                        "description": "강의정보",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Class"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "강의 생성에 실패 했습니다.",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/class/enroll/{classCode}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "수강신청",
                "parameters": [
                    {
                        "type": "string",
                        "description": "학수번호",
                        "name": "classCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "학수번호가 유효하지 않습니다.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "강의가 존재하지 않습니다..",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/class/{classQuizSetId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "ClassQuizSetId에 해당하는 Quiz set 을 가져온다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "불러오려는 QuizSetId",
                        "name": "classQuizSetId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "퀴즈 셋",
                        "schema": {
                            "$ref": "#/definitions/dto.ClassQuizSetGetForm"
                        }
                    },
                    "400": {
                        "description": "INVALID_PATH_PARAMETER",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/classes/{classCode}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "해당 강의가 가지고 있는 모든 퀴즈 셋을 반환한다",
                "parameters": [
                    {
                        "description": "퀴즈 셋을 조회할 강의의 학수 번호",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.QuizCreateForm"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/quizset": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "퀴즈셋을 생성한다",
                "parameters": [
                    {
                        "description": "퀴즈셋 정보",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.QuizSetCreateForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/quizset/{quizSetId}/class/{classCode}": {
            "post": {
                "description": "퀴즈 셋 고유 아이디와 학수번호를 받아서 클래스의 퀴즈셋 목록에 추가한다.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "유저의 퀴즈셋에 있는 퀴즈를 강의의 퀴즈셋으로 가져온다",
                "parameters": [
                    {
                        "type": "string",
                        "description": "클래스에 불러오려는 퀴즈셋의 아이디",
                        "name": "quizSetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "퀴즈를 추가하려는 강의의 학수번호",
                        "name": "classCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "요청이 올바르지 않습니다.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "INTERNAL_SERVER_ERROR =\u003e \u003cdetail message\u003e",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "퀴즈 셋 고유 아이디와 학수번호를 이용해서 클래스에서 퀴즈셋을 삭제한다.\n원본 퀴즈 셋은 삭제되지 않는다.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "클래스에 있는 퀴즈셋을 삭제한다",
                "parameters": [
                    {
                        "type": "string",
                        "description": "강의 에서 삭제 하려는 퀴즈셋의 아이디",
                        "name": "quizSetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "강의의 학수번호",
                        "name": "classCode",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "CLASS_QUIZ_SET_NOT_EXISTS",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/quizset/{quizSetId}/quiz/{quizId}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "퀴즈셋에 있는 퀴즈를 삭제한다",
                "parameters": [
                    {
                        "type": "string",
                        "description": "퀴즈를 삭제 하려는 퀴즈셋의 아이디",
                        "name": "quizSetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "삭제 하려는 퀴즈의 아이디",
                        "name": "quizId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "성공",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "요청이 올바르지 않습니다.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/quizset/{quizsetId}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "퀴즈셋을 삭제한다",
                "parameters": [
                    {
                        "type": "string",
                        "description": "퀴즈 셋 고유 아이디",
                        "name": "quizsetId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "INVALID_PATH_PARAMETER",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/quizset/{quizsetId}/quiz": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "퀴즈셋에 퀴즈를 추가한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "퀴즈를 추가하려는 퀴즈셋의 아이디",
                        "name": "quizSetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "퀴즈 정보",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.QuizCreateForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "성공",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "요청이 올바르지 않습니다.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/result/users/{userName}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "유저가 가진 모든 퀴즈 셋 채점 결과를 가져온다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "채점 결과를 받아오려는 유저의 로그인 아이디(즉 이메일)",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "채점된 퀴즈셋 목록(배열)",
                        "schema": {
                            "$ref": "#/definitions/dto.GetQuizSetResults"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/score": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "Quiz set 을 채점한다.",
                "parameters": [
                    {
                        "description": "채점 하려는 퀴즈 셋",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.QuizSetForScoring"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "INVALID_PATH_PARAMETER",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/quizsets/users/{username}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Quiz set"
                ],
                "summary": "유저가 만든 모든 퀴즈셋 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "유저 로그인 아이디",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "퀴즈셋목록(배열)",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.QuizSetGetForm"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/classes/enrolled/{userName}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "수강하고 있는 강의목록 조회",
                "parameters": [
                    {
                        "type": "string",
                        "description": "유저 아이디",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "강의목록(배열)",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Class"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/classes/managing/{userName}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Class"
                ],
                "summary": "유저가 관리하고 있는 모든 강의들을 가져온다",
                "parameters": [
                    {
                        "type": "string",
                        "description": "유저 아이디",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "강의목록(배열)",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Class"
                            }
                        }
                    },
                    "400": {
                        "description": "학수번호가 유효하지 않습니다.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "강의가 존재하지 않습니다..",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info/{userName}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "유저 정보를 반환한다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "유저 아이디",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "유저 정보",
                        "schema": {
                            "$ref": "#/definitions/dto.UserGetForm"
                        }
                    },
                    "404": {
                        "description": "user not exists ",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "summary": "로그인",
                "parameters": [
                    {
                        "description": "로그인 정보",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "USER_NOT_EXIST",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "회원가입",
                "parameters": [
                    {
                        "description": "유저정보",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "tags": [
                    "Users"
                ],
                "summary": "유저 로그인 아이디 중복 확인",
                "parameters": [
                    {
                        "type": "string",
                        "description": "확인 하려는  로그인 아이디",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "USER_NOT_EXIST",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Class": {
            "type": "object",
            "properties": {
                "class_code": {
                    "type": "string",
                    "example": "형식언어"
                },
                "class_name": {
                    "type": "string",
                    "example": "CSE-1234"
                }
            }
        },
        "dto.ClassQuizSetGetForm": {
            "type": "object",
            "properties": {
                "class_quiz_set_id": {
                    "type": "integer",
                    "example": 1234
                },
                "quiz_set_author_name": {
                    "type": "string",
                    "example": "퀴즈셋 생성자 로그인 아이디"
                },
                "quiz_set_name": {
                    "type": "string",
                    "example": "퀴즈셋 이름"
                },
                "quizzes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.QuizGetForm"
                    }
                },
                "total_score": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "dto.GetQuizSetResult": {
            "type": "object",
            "properties": {
                "class_code": {
                    "type": "string"
                },
                "class_name": {
                    "type": "string"
                },
                "my_score": {
                    "type": "integer"
                },
                "quiz_set_name": {
                    "type": "string"
                },
                "total_score": {
                    "type": "integer"
                }
            }
        },
        "dto.GetQuizSetResults": {
            "type": "object",
            "properties": {
                "quiz_set_results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GetQuizSetResult"
                    }
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.QuizCreateForm": {
            "type": "object",
            "properties": {
                "quiz_answer": {
                    "type": "string",
                    "example": "퀴즈 정답"
                },
                "quiz_content": {
                    "type": "string",
                    "example": "퀴즈 렌더링 정보"
                },
                "quiz_title": {
                    "type": "string",
                    "example": "퀴즈 문제"
                },
                "quiz_type": {
                    "type": "string",
                    "example": "퀴즈 타입(MULTI 혹은 SIMPLE)"
                }
            }
        },
        "dto.QuizForScoring": {
            "type": "object",
            "properties": {
                "quiz_answer": {
                    "type": "string"
                },
                "quiz_id": {
                    "type": "integer"
                },
                "quiz_type": {
                    "type": "string"
                }
            }
        },
        "dto.QuizGetForm": {
            "type": "object",
            "properties": {
                "quiz_answer": {
                    "type": "string",
                    "example": "퀴즈 정답"
                },
                "quiz_content": {
                    "type": "string",
                    "example": "퀴즈 렌더링 정보"
                },
                "quiz_id": {
                    "type": "integer",
                    "example": 1234
                },
                "quiz_title": {
                    "type": "string",
                    "example": "퀴즈 문제"
                },
                "quiz_type": {
                    "type": "string",
                    "example": "퀴즈 타입(MULTI 혹은 SIMPLE)"
                }
            }
        },
        "dto.QuizSetCreateForm": {
            "type": "object",
            "properties": {
                "class_code": {
                    "type": "string"
                },
                "quiz_set_author_name": {
                    "type": "string",
                    "example": "퀴즈셋 생성자 로그인 아이디"
                },
                "quiz_set_name": {
                    "type": "string",
                    "example": "퀴즈셋 이름"
                },
                "quizzes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.QuizCreateForm"
                    }
                },
                "total_score": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "dto.QuizSetForScoring": {
            "type": "object",
            "properties": {
                "class_quiz_set_id": {
                    "type": "integer",
                    "example": 1234
                },
                "quiz_for_scorings": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.QuizForScoring"
                    }
                },
                "username": {
                    "description": "email 과 같음",
                    "type": "string",
                    "example": "유저 이메일"
                }
            }
        },
        "dto.QuizSetGetForm": {
            "type": "object",
            "properties": {
                "quiz_set_author_name": {
                    "type": "string",
                    "example": "퀴즈셋 생성자 로그인 아이디"
                },
                "quiz_set_id": {
                    "type": "integer",
                    "example": 1234
                },
                "quiz_set_name": {
                    "type": "string",
                    "example": "퀴즈셋 이름"
                },
                "quizzes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.QuizGetForm"
                    }
                },
                "total_score": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "donggukmail@dgu.ac.kr"
                },
                "nickname": {
                    "type": "string",
                    "example": "my nickname"
                },
                "password": {
                    "type": "string"
                },
                "student_code": {
                    "type": "integer",
                    "example": 2015123456
                },
                "username": {
                    "type": "string",
                    "example": "pigeon1234"
                }
            }
        },
        "dto.UserGetForm": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "donggukmail@dgu.ac.kr"
                },
                "nickname": {
                    "type": "string",
                    "example": "my nickname"
                },
                "professor": {
                    "type": "boolean",
                    "example": true
                },
                "student_code": {
                    "type": "integer",
                    "example": 2015123456
                },
                "username": {
                    "type": "string",
                    "example": "pigeon1234"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
