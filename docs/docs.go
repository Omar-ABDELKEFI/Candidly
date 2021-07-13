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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "fiber@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/questions": {
            "post": {
                "description": "add by json question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question"
                ],
                "summary": "add an question",
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
                ]
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "$ref": "#/definitions/sql.NullTime"
        },
        "models.Question": {
            "type": "object",
            "required": [
                "difficulty",
                "expected_time",
                "name",
                "points"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "object",
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
                "name": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                },
                "question_text": {
                    "type": "string"
                },
                "skill_id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "sql.NullTime": {
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "tekab-test",
	Description: "interview assessment tests for interviewing out of the box",
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
