// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Will Kerwin",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/authors": {
            "get": {
                "description": "get the authors from database.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Get Authors.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Author"
                            }
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "creates an author asynchronously.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Create an author.",
                "parameters": [
                    {
                        "description": "author body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Author"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            }
        },
        "/authors/{id}": {
            "get": {
                "description": "get the author by id from database.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Get Author by its object id in hex format.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the author",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the author by id from database.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Delete Author by its object id in hex format.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the author",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "description": "get the books from database with filters.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get Books.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title of the book",
                        "name": "title",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "genre of the book",
                        "name": "genre",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "authorId of the book",
                        "name": "authorId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Book"
                            }
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "creates a book asynchronously.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create an book.",
                "parameters": [
                    {
                        "description": "book body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "description": "get the book by id from database.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get book by its object id in hex format.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the book",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the book by id from database.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete book by its object id in hex format.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the book",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update the book by id from database.",
                "consumes": [
                    "applicaiton/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Update book by its object id in hex format.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the book",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body of the book",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "$ref": "#/definitions/models.ApiErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ApiErrorResponse": {
            "type": "object",
            "additionalProperties": true
        },
        "models.Author": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "age": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Book": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "authorId": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "synopsis": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Go Microservice Bookstore API",
	Description:      "This is the api for the go bookstore microservices project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}