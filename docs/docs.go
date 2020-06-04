// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-04 10:50:25.247671 +0800 CST m=+0.097794066

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
        "/api/v1/getinfo": {
            "get": {
                "description": "获取JSON",
                "tags": [
                    "Info"
                ],
                "summary": "获取权限信息",
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
        "/api/v1/login": {
            "post": {
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.ReqLoginUser"
                        }
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
        "/api/v1/logout": {
            "post": {
                "tags": [
                    "User"
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
        },
        "/api/v1/menu": {
            "put": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "修改菜单",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"修改失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Menu"
                ],
                "summary": "创建菜单",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"添加失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/menu/{menuId}": {
            "get": {
                "description": "获取JSON",
                "tags": [
                    "Menu"
                ],
                "summary": "获取Menu数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "menuId",
                        "name": "menuId",
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
            },
            "delete": {
                "description": "删除数据",
                "tags": [
                    "Menu"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "menuId",
                        "name": "menuId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"删除失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/menulist": {
            "get": {
                "description": "获取JSON",
                "tags": [
                    "Menu"
                ],
                "summary": "Menu列表数据",
                "parameters": [
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
        "/api/v1/menurole": {
            "get": {
                "description": "获取JSON",
                "tags": [
                    "Menu"
                ],
                "summary": "根据角色名称获取菜单列表数据（左菜单使用）",
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
        "/api/v1/query": {
            "get": {
                "description": "查询登录状态，并返回user",
                "tags": [
                    "User"
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
        "/api/v1/role": {
            "put": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "修改用户角色",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Role"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"修改失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Role"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Role"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"添加失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/role/{roleId}": {
            "get": {
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
            },
            "delete": {
                "description": "删除数据",
                "tags": [
                    "Role"
                ],
                "summary": "删除用户角色",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roleId",
                        "name": "roleId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"删除失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/rolelist": {
            "get": {
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
        "/api/v1/user": {
            "put": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "修改用户数据",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.ReqUpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"修改失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "获取JSON",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "用户数据",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dao.ReqAddUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"添加失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{userId}": {
            "get": {
                "description": "获取JSON",
                "tags": [
                    "User"
                ],
                "summary": "获取用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
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
            },
            "delete": {
                "description": "删除数据",
                "tags": [
                    "User"
                ],
                "summary": "删除用户数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\": -1, \"message\": \"删除失败\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/userlist": {
            "get": {
                "description": "获取JSON",
                "tags": [
                    "User"
                ],
                "summary": "列表数据",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "status",
                        "name": "status",
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
                        "description": "{\"code\": -1, \"message\": \"抱歉未找到相关信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.ReqAddUser": {
            "type": "object",
            "required": [
                "nickname",
                "password",
                "roleId",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "createBy": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "dataScope": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "params": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "roleId": {
                    "type": "integer"
                },
                "salt": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dao.ReqLoginUser": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dao.ReqUpdateUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "createBy": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "dataScope": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "params": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "roleId": {
                    "type": "integer"
                },
                "salt": {
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Menu": {
            "type": "object",
            "required": [
                "menuType",
                "sort",
                "title"
            ],
            "properties": {
                "action": {
                    "type": "string"
                },
                "breadcrumb": {
                    "type": "string"
                },
                "component": {
                    "type": "string"
                },
                "createBy": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "isFrame": {
                    "type": "string"
                },
                "is_select": {
                    "type": "boolean"
                },
                "menuId": {
                    "type": "integer"
                },
                "menuType": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "noCache": {
                    "type": "boolean"
                },
                "params": {
                    "type": "string"
                },
                "parentId": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                },
                "paths": {
                    "type": "string"
                },
                "permission": {
                    "type": "string"
                },
                "routes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Menu"
                    }
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "visible": {
                    "type": "string"
                }
            }
        },
        "model.Role": {
            "type": "object",
            "required": [
                "menuIds",
                "roleKey",
                "roleName",
                "roleSort",
                "status"
            ],
            "properties": {
                "createBy": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "dataScope": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "menuIds": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "params": {
                    "type": "string"
                },
                "remark": {
                    "type": "string"
                },
                "roleId": {
                    "type": "integer"
                },
                "roleKey": {
                    "type": "string"
                },
                "roleName": {
                    "type": "string"
                },
                "roleSort": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "updateBy": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
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
