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
        "contact": {
            "name": "Data Platform",
            "email": "dataplatform@tiket.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth": {
            "post": {
                "description": "API will try to validate user credential and return token that can be used for calling other methods.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Authenticate user",
                "parameters": [
                    {
                        "description": "user credential",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserAuth"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    }
                }
            }
        },
        "/api/seller/detail": {
            "get": {
                "description": "Api will get user details by token.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer + [token]",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "description": "user",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    }
                }
            }
        },
        "/seller/product": {
            "post": {
                "description": "User can create products by this REST",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Seller create product",
                "parameters": [
                    {
                        "description": "create",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Product"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer + [token]",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    }
                }
            }
        },
        "/seller/type": {
            "get": {
                "description": "Api will get type of products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get type for products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer + [token]",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Wrapper"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Image": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "productId": {
                    "type": "integer"
                }
            }
        },
        "dto.Product": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "cod": {
                    "type": "boolean"
                },
                "delivery": {
                    "type": "boolean"
                },
                "delivery_cost": {
                    "type": "integer"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Image"
                    }
                },
                "index": {
                    "type": "integer"
                },
                "isactive": {
                    "type": "boolean"
                },
                "isclosed": {
                    "type": "boolean"
                },
                "istimeclosed": {
                    "type": "string"
                },
                "istimeopen": {
                    "type": "string"
                },
                "lat": {
                    "type": "number"
                },
                "long": {
                    "type": "number"
                },
                "price": {
                    "type": "integer"
                },
                "satuan": {
                    "type": "string"
                },
                "stocks": {
                    "type": "integer"
                },
                "typeId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "brandname": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lat": {
                    "type": "string"
                },
                "long": {
                    "type": "string"
                },
                "notif": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Product"
                    }
                },
                "roleId": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "dto.UserAuth": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "wrapper.Wrapper": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "serverTime": {
                    "type": "integer"
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
	Version:     "0.0.1",
	Host:        "0.0.0.0:8080",
	BasePath:    "",
	Schemes:     []string{"http", "https"},
	Title:       "DMS Core Module API",
	Description: "This API is used to serve DMS Front End",
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
