// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-22 15:30:28.477884 +0800 CST m=+0.054940846

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
        "/api/v1/menulist": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "获取JSON",
                "tags": [
                    "Menu"
                ],
                "summary": "Menu列表数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "menuName",
                        "name": "menuName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "visible",
                        "name": "visible",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/{roleId}": {
            "get": {
                "security": [
                    {
                        "": []
                    }
                ],
                "description": "获取JSON",
                "tags": [
                    "Role"
                ],
                "summary": "获取Role数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roleId",
                        "name": "roleId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/rolelist": {
            "get": {
                "security": [
                    {
                        "": []
                    }
                ],
                "description": "Get JSON",
                "tags": [
                    "Role"
                ],
                "summary": "角色列表数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roleName",
                        "name": "roleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "rolekey",
                        "name": "rolekey",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页条数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "pageIndex",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": 200, \"data\": [...]}",
                        "schema": {
                            "$ref": "#/definitions/response.PageResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "admin",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "admin@123",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Res"
                        }
                    }
                }
            }
        },
        "/query": {
            "post": {
                "description": "查询登录状态，并返回user",
                "tags": [
                    "user"
                ],
                "summary": "查询登录状态，并返回user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Res"
                        }
                    }
                }
            }
        },
        "/user/add": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "管理员添加用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "dahuang",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "大黄",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Res"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "注销登录",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Res"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Page": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "object"
                },
                "pageIndex": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "response.PageResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/response.Page"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "response.Res": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
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
