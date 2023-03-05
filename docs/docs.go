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
        "/expense-summary/{year}/{month}": {
            "get": {
                "description": "gets expense summary for a certain month and year",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "gets expense summary for a certain month and year",
                "parameters": [
                    {
                        "type": "string",
                        "description": "current year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "current month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ExpenseSummary"
                        }
                    }
                }
            }
        },
        "/expense/{expense_type}/{year}/{month}": {
            "get": {
                "description": "gets expenses for a certain month and year",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "gets expenses for a certain month and year",
                "parameters": [
                    {
                        "type": "string",
                        "description": "expense type",
                        "name": "expense_type",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "current year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "current month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
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
        "/expenses/{year}/{month}": {
            "get": {
                "description": "gets expenses for a certain month and year",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expenses"
                ],
                "summary": "gets expenses for a certain month and year",
                "parameters": [
                    {
                        "type": "string",
                        "description": "current year",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "current month",
                        "name": "month",
                        "in": "path",
                        "required": true
                    }
                ],
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
        "/split/latest": {
            "get": {
                "description": "gets latest expense split percentage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "split"
                ],
                "summary": "gets latest expense split percentage",
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
                    "$ref": "#/definitions/models.ExpenseType"
                },
                "status": {
                    "$ref": "#/definitions/models.ExpenseStatus"
                }
            }
        },
        "models.ExpenseStatus": {
            "type": "string",
            "enum": [
                "Paid",
                "Covered",
                "Unpaid"
            ],
            "x-enum-varnames": [
                "PaidStatus",
                "CoveredStatus",
                "UnpaidStatus"
            ]
        },
        "models.ExpenseSummary": {
            "type": "object",
            "properties": {
                "herAmount": {
                    "type": "number"
                },
                "myAmount": {
                    "type": "number"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.ExpenseType": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "splitType": {
                    "$ref": "#/definitions/models.SplitType"
                },
                "typeName": {
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
        },
        "models.SplitType": {
            "type": "string",
            "enum": [
                "FiftyFifty",
                "Default",
                "MeOnly",
                "HerOnly"
            ],
            "x-enum-varnames": [
                "FiftyFiftySplit",
                "DefaultSplit",
                "MeOnlySplit",
                "HerOnlySplit"
            ]
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
