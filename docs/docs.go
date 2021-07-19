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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/candidate": {
            "post": {
                "description": "create new candidate by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "candidate"
                ],
                "summary": "add new  candidate",
                "parameters": [
                    {
                        "description": "Add candidate",
                        "name": "candidate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Candidate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Candidate"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login to the app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login to the app",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/my-tests/candidates/:id": {
            "post": {
                "description": "create test_candidate by json and path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "test_candidate"
                ],
                "summary": "add new  test_candidate",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Add candidate",
                        "name": "test_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add candidate",
                        "name": "test_candidate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TestCandidate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TestCandidate"
                        }
                    }
                }
            }
        },
        "/questions": {
            "get": {
                "description": "find a question by type or difficulty",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "find a question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "question search by type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "question search by difficulty",
                        "name": "difficulty",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Question"
                            }
                        }
                    }
                }
            }
        },
        "/questions/edit": {
            "post": {
                "description": "create new question by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "add new  question",
                "parameters": [
                    {
                        "description": "Add question",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Question"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Question"
                        }
                    }
                }
            }
        },
        "/skill": {
            "post": {
                "description": "create new skill by json",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "skill"
                ],
                "summary": "add new  skill",
                "parameters": [
                    {
                        "description": "Add Skill",
                        "name": "skill",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Skill"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Question"
                        }
                    }
                }
            }
        },
        "/tests": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "get tests by skill",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tests"
                ],
                "summary": "get tests",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tests search by type",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Test"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Answer": {
            "type": "object",
            "required": [
                "question_id",
                "test_candidate_id"
            ],
            "properties": {
                "answer_choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.AnswerChoices"
                    }
                },
                "answer_file": {
                    "type": "string"
                },
                "answer_text": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "point": {
                    "type": "number"
                },
                "question_id": {
                    "type": "integer"
                },
                "test_candidate_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.AnswerChoices": {
            "type": "object",
            "required": [
                "answer_id",
                "choices_id"
            ],
            "properties": {
                "answer_id": {
                    "type": "integer"
                },
                "choices_id": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Candidate": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "test_candidate": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TestCandidate"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Choices": {
            "type": "object",
            "required": [
                "is_answer"
            ],
            "properties": {
                "answer_choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.AnswerChoices"
                    }
                },
                "choice_text": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_answer": {
                    "type": "boolean"
                },
                "question_id": {
                    "type": "integer"
                }
            }
        },
        "models.Question": {
            "type": "object",
            "required": [
                "difficulty",
                "expected_time",
                "max_points",
                "name",
                "question_text"
            ],
            "properties": {
                "answer": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Choices"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "difficulty": {
                    "type": "string"
                },
                "expected_time": {
                    "type": "integer"
                },
                "file_read_me": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "max_points": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "question_text": {
                    "type": "string"
                },
                "skill_id": {
                    "type": "integer"
                },
                "test_questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TestQuestion"
                    }
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Skill": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "question": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Question"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Test": {
            "type": "object",
            "required": [
                "archived",
                "description",
                "name",
                "passingScore",
                "showScore",
                "timingPolicy"
            ],
            "properties": {
                "archived": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "notifyEmails": {
                    "type": "string"
                },
                "passingScore": {
                    "type": "integer"
                },
                "showScore": {
                    "type": "boolean"
                },
                "test_candidate": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TestCandidate"
                    }
                },
                "test_questions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TestQuestion"
                    }
                },
                "timingPolicy": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.TestCandidate": {
            "type": "object",
            "required": [
                "candidate_id",
                "test_id"
            ],
            "properties": {
                "answer": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "candidate_id": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "score": {
                    "type": "number"
                },
                "test_id": {
                    "type": "integer"
                },
                "test_status": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.TestQuestion": {
            "type": "object",
            "required": [
                "question_id",
                "test_id"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "question_id": {
                    "type": "integer"
                },
                "test_id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "tekab-test",
	Description: "this is an application web of interview assessment tests for interviewing out of the box .",
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
