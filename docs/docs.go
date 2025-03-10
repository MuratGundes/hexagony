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
        "contact": {
            "name": "Cyro Dubeux",
            "url": "https://github.com/cyruzin/hexagony/issues/new",
            "email": "xorycx@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/cyruzin/hexagony/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/album": {
            "get": {
                "description": "lists all albums",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "album"
                ],
                "summary": "List of albums",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Album"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "add a new album",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "album"
                ],
                "summary": "Add an album",
                "parameters": [
                    {
                        "description": "add a new album",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Album"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    }
                }
            }
        },
        "/album/{uuid}": {
            "get": {
                "description": "lists an album by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "album"
                ],
                "summary": "List an album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "album uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Album"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "update an album by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "album"
                ],
                "summary": "Update an album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "album uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update an album by uuid",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Album"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete an album by uuid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "album"
                ],
                "summary": "Delete an album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "album uuid",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/rest.APIMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Album": {
            "type": "object",
            "required": [
                "length",
                "name"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "length": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "rest.APIMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Hexagony API",
	Description:      "Clean architecture example in Golang.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
