// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Rick Shang",
            "url": "https://gitee.com/nothing-is-nothing",
            "email": "2227627947@qq.com"
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
        "/api/v1/admin/login": {
            "post": {
                "description": "管理员登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台相关接口"
                ],
                "summary": "管理员登录接口",
                "parameters": [
                    {
                        "description": "查询参数",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ParamLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._RespAdminLogin"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/modules": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加xss模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前后台共用接口"
                ],
                "summary": "添加xss模块",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "desc",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "is_common",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "option_list",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "param_list",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "xss_payload",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Resp"
                        }
                    }
                }
            }
        },
        "/api/v1/module/{id}": {
            "get": {
                "description": "用户只可以获取公共模块的详情和属于自己的模块的详情",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前台相关接口"
                ],
                "summary": "用户获取模块详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "模块id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer 用户token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._RespModuleDetail"
                        }
                    }
                }
            }
        },
        "/api/v1/modules": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加xss模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前后台共用接口"
                ],
                "summary": "添加xss模块",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "desc",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "is_common",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "option_list",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "name": "param_list",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "xss_payload",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Resp"
                        }
                    }
                }
            }
        },
        "/api/v1/projectform": {
            "get": {
                "description": "获取添加项目需要的表单，因为模块是变化的，所以添加项目需要的填的表单项需要后端动态生成。创建项目时需要填的是一个空表单，无序提交id。配置项目时需要对有数据的表单进行修改，需要提交项目id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "前台相关接口"
                ],
                "summary": "获取创建项目/配置项目需要填的表单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "项目id",
                        "name": "object",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers._RespProjectForm"
                        }
                    }
                }
            }
        },
        "/basicAuth": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "BasicAuth 钓鱼",
                "tags": [
                    "攻击相关接口"
                ],
                "summary": "BasicAuth 钓鱼",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Base64编码的用户名和密码",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Resp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Resp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "object"
                }
            }
        },
        "controllers._RespAdminLogin": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "object"
                },
                "token": {
                    "description": "jwt令牌",
                    "type": "string"
                }
            }
        },
        "controllers._RespModuleDetail": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "module": {
                    "$ref": "#/definitions/models.Module"
                },
                "msg": {
                    "type": "object"
                }
            }
        },
        "controllers._RespProjectForm": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "form": {
                    "$ref": "#/definitions/models.ParamCreateProject"
                },
                "msg": {
                    "type": "object"
                }
            }
        },
        "models.Module": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_admin": {
                    "type": "boolean"
                },
                "is_common": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "option_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "param_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                },
                "xss_payload": {
                    "type": "string"
                }
            }
        },
        "models.ModuleItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_choosed": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "option_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OptionItem"
                    }
                }
            }
        },
        "models.OptionItem": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.ParamCreateProject": {
            "type": "object",
            "required": [
                "desc",
                "module_list",
                "name"
            ],
            "properties": {
                "desc": {
                    "type": "string"
                },
                "module_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ModuleItem"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.ParamLogin": {
            "type": "object",
            "required": [
                "g_recaptcha_response",
                "password",
                "username"
            ],
            "properties": {
                "g_recaptcha_response": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "XSS Platform",
	Description: "基于gin+vue+mysql+redis 实现的xss平台",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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