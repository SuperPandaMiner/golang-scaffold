// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/swaggo/get/{id}": {
            "get": {
                "security": [
                    {
                        "ACCESS_TOKEN": []
                    }
                ],
                "description": "get",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "字符串",
                        "name": "string",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "布尔",
                        "name": "bool",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "数字",
                        "name": "number",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ViewVO"
                        }
                    }
                }
            }
        },
        "/swaggo/post": {
            "post": {
                "security": [
                    {
                        "ACCESS_TOKEN": []
                    }
                ],
                "description": "post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "api"
                ],
                "summary": "post",
                "parameters": [
                    {
                        "description": "params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ViewVO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.ViewVO"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ViewVO": {
            "type": "object",
            "properties": {
                "bool": {
                    "description": "布尔",
                    "type": "boolean"
                },
                "number": {
                    "description": "数字",
                    "type": "integer"
                },
                "string": {
                    "description": "字符串",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ACCESS_TOKEN": {
            "type": "apiKey",
            "name": "access_token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}