// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/expense/:year/:month": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Expense"
                        }
                    }
                }
            }
        },
        "/expenses": {
            "get": {
                "description": "Get all expenses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "Get all expenses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Expense"
                            }
                        }
                    }
                }
            }
        },
        "/split": {
            "post": {
                "description": "adds expense split percentage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "split"
                ],
                "summary": "adds expense split percentage",
                "parameters": [
                    {
                        "description": "Split Descriptor",
                        "name": "Split",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SplitSpecifier"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Split"
                        }
                    }
                }
            }
        },
        "/split/:year/:month": {
            "get": {
                "description": "gets expense split percentage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "split"
                ],
                "summary": "gets expense split percentage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Split"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Expense": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "expenseID": {
                    "type": "integer"
                },
                "expenseType": {
                    "type": "string"
                }
            }
        },
        "models.Split": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "herPercentage": {
                    "type": "number"
                },
                "myPercentage": {
                    "type": "number"
                },
                "splitID": {
                    "type": "integer"
                }
            }
        },
        "models.SplitSpecifier": {
            "type": "object",
            "properties": {
                "herNetIncome": {
                    "type": "number"
                },
                "herPercentage": {
                    "type": "number"
                },
                "myNetIncome": {
                    "type": "number"
                },
                "myPercentage": {
                    "type": "number"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Rent Cat",
	Description:      "Manages expense sharing",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
