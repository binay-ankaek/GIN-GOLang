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
        "/api/auth/login": {
            "post": {
                "description": "Authenticate a user and return a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login to the account",
                "parameters": [
                    {
                        "description": "User login details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/profile": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Retrieve the profile of the currently authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.GetProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update user profile information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "description": "User profile update details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UpdateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.UpdateProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete the user profile associated with the provided token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "description": "Register a new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new account",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/contact/": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Retrieve all contacts for the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Get all contacts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/http.ContactResponse"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/contact/add": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Add a new contact for the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contacts"
                ],
                "summary": "Add a new contact",
                "parameters": [
                    {
                        "description": "Contact details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.AddContactRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.AddContactResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.AddContactRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "phone": {
                    "type": "string",
                    "example": "1234567890"
                }
            }
        },
        "http.AddContactResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "contact added successfully"
                }
            }
        },
        "http.ContactResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "phone": {
                    "type": "string",
                    "example": "1234567890"
                }
            }
        },
        "http.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Invalid input"
                }
            }
        },
        "http.GetProfileResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "format": "string",
                    "example": "user@example.com"
                },
                "image": {
                    "type": "string",
                    "format": "string",
                    "example": "http://example.com/image.jpg"
                },
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "John Doe"
                },
                "phone": {
                    "type": "string",
                    "format": "string",
                    "example": "1234567890"
                }
            }
        },
        "http.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "format": "string",
                    "example": "password123"
                },
                "phone": {
                    "type": "string",
                    "format": "string",
                    "example": "1234567890"
                }
            }
        },
        "http.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "login successful"
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                }
            }
        },
        "http.RegisterRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "format": "string",
                    "example": "password123"
                },
                "phone": {
                    "type": "string",
                    "format": "string",
                    "example": "1234567890"
                }
            }
        },
        "http.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "operation successful"
                }
            }
        },
        "http.UpdateProfileRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "format": "string",
                    "example": "newuser@example.com"
                },
                "image": {
                    "type": "string",
                    "format": "string",
                    "example": "http://example.com/newimage.jpg"
                },
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "Jane Doe"
                },
                "password": {
                    "type": "string",
                    "format": "string",
                    "example": "newpassword123"
                },
                "phone": {
                    "type": "string",
                    "format": "string",
                    "example": "1234567890"
                }
            }
        },
        "http.UpdateProfileResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "format": "string",
                    "example": "newuser@example.com"
                },
                "image": {
                    "type": "string",
                    "format": "string",
                    "example": "http://example.com/newimage.jpg"
                },
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "Jane Doe"
                },
                "phone": {
                    "type": "string",
                    "format": "string",
                    "example": "1234567890"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "Header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "aok-connect Business API",
	Description:      "This is a simple RESTful Service API written in Go using Gin web framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
