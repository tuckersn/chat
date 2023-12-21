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
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/message/id/:messageKey": {
            "get": {
                "description": "Get the Message with the given id (messageKey)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message API"
                ],
                "summary": "Get a Message",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Deletes a Message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message API"
                ],
                "summary": "Deletes a Message",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/note/:path": {
            "get": {
                "description": "Returns the content of a Note based on it's path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Note API"
                ],
                "summary": "Returns the content of a Note",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/room/:roomKey/message": {
            "post": {
                "description": "Send a Message to a given Room with the given content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message API",
                    "Room API"
                ],
                "summary": "Message a Room",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/room/id/:roomKey": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Room API"
                ],
                "summary": "Get a Room's information",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/system/ping": {
            "get": {
                "description": "says hello to the server, used for overview page",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System API"
                ],
                "summary": "ping the server, used for overview page",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/user": {
            "post": {
                "description": "creates a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "create a new user",
                "parameters": [
                    {
                        "description": "User creation request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/id/:userId": {
            "post": {
                "description": "creates a page based on it's path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "creates a page",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/id/:username": {
            "get": {
                "description": "retrieves information about the user specified by the username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "get a user's info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username of the user to retrieve",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User information",
                        "schema": {
                            "$ref": "#/definitions/api.UserGetResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid username",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "User not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "deletes a user based on it's path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "deletes a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/webhook": {
            "get": {
                "description": "Returns a list of WebHooks this user owns.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Webhook API"
                ],
                "summary": "List WebHooks that you own",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "[Two Factor Authentication Required] Creates a new WebHook with the provided subscription details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Webhook API"
                ],
                "summary": "[2FA] Create a new WebHook",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/webhook/:webhook": {
            "get": {
                "description": "Provided either the key, name, or id returns the details of the webhook.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Webhook API"
                ],
                "summary": "Get's the details of a WebHook",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "[Two Factor Authentication Required] Updates the details of an existing WebHook at the given WebHook key or id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Webhook API"
                ],
                "summary": "[2FA] Update the details of an existing WebHook",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "[Two Factor Authentication Required] Deletes a WebHook at the given WebHook key or id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Webhook API"
                ],
                "summary": "[2FA] Deletes a WebHook",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login/gitlab": {
            "get": {
                "description": "where the user's browser is sent by GitLab after completing the oauth2 flow",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "handles GitLab oauth2 callback",
                "responses": {
                    "302": {
                        "description": "Found"
                    }
                }
            }
        },
        "/login/google": {
            "get": {
                "description": "{art 1 of the HTTP redirect to Google's OpenID Connect (OAuth 2.0) consent screen",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "HTTP redirect to Google's OpenID Connect (OAuth 2.0) consent screen",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login/google/disconnect": {
            "delete": {
                "description": "[Two Factor Authentication Required] Removes your Google identity record from your account. This will prevent you from logging in with Google.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "[2FA] Removes your Google account information from your account.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login/google/receive": {
            "get": {
                "description": "Part 2 of the HTTP redirect to Google's OpenID Connect (OAuth 2.0) consent screen",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Redirect URI receiving address for Googl'e OAuth 2.0 flow",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login/recent": {
            "get": {
                "description": "gets a list of recent logins for provided token's associated user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Gets your recent logins (up to 10)",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/note/:path": {
            "post": {
                "description": "Writes the provided content to the Note path provided in the url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Note API"
                ],
                "summary": "Writes a Note",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a note based on it's path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Note API"
                ],
                "summary": "Deletes a note",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.UserCreateRequest": {
            "type": "object",
            "properties": {
                "admin": {
                    "type": "boolean"
                },
                "displayName": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.UserGetResponse": {
            "type": "object",
            "properties": {
                "displayName": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Chat Backend API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
