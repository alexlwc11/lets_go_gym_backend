// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
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
        "/app_info": {
            "get": {
                "description": "Get latest app info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AppInfo"
                ],
                "summary": "Get app info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AppInfo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/bookmarks": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get user bookmarked sports centers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bookmarks"
                ],
                "summary": "GetUserBookmarks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.userBookmarkOutDto"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user bookmarked sports centers",
                "tags": [
                    "Bookmarks"
                ],
                "summary": "UpdateUserBookmarks",
                "parameters": [
                    {
                        "description": "Updated sports centers IDs",
                        "name": "userUpdatedSportsCenterIds",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.putUserBookmarksInDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/districts": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get latest districts data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Districts"
                ],
                "summary": "Get all districts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.districtsOutDto"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "description": "Get new set of tokens with refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "refresh_token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.refreshInDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.sessionTokenOutDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/regions": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get latest regions data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Regions"
                ],
                "summary": "Get all regions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.regionsOutDto"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "New user registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "User info for registration",
                        "name": "user_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.userInfoInDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.sessionTokenOutDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/sign_in": {
            "post": {
                "description": "Existing user sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "User info for signing in",
                        "name": "user_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/apis.userInfoInDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.sessionTokenOutDto"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/sports_centers": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "getting latest sports centers data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sports centers"
                ],
                "summary": "Get all sports centers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.sportsCentersOutDto"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/sports_centers/{id}/details_url": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the details url for specified sports center",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sports centers"
                ],
                "summary": "Get details url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Sports center ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/apis.detailsUrlOutDto"
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "apis.detailsUrlOutDto": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "apis.districtsOutDto": {
            "type": "object",
            "properties": {
                "districts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.District"
                    }
                }
            }
        },
        "apis.putUserBookmarksInDto": {
            "type": "object",
            "properties": {
                "updated_sports_center_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "apis.refreshInDto": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "apis.regionsOutDto": {
            "type": "object",
            "properties": {
                "regions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Region"
                    }
                }
            }
        },
        "apis.sessionTokenOutDto": {
            "type": "object",
            "properties": {
                "refresh_expired_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "session_expired_at": {
                    "type": "string"
                },
                "session_token": {
                    "type": "string"
                }
            }
        },
        "apis.sportsCentersOutDto": {
            "type": "object",
            "properties": {
                "sports_centers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SportsCenter"
                    }
                }
            }
        },
        "apis.userBookmarkOutDto": {
            "type": "object",
            "properties": {
                "sports_center_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "apis.userInfoInDto": {
            "type": "object",
            "properties": {
                "device_uuid": {
                    "type": "string"
                }
            }
        },
        "models.AppInfo": {
            "type": "object",
            "properties": {
                "data_info": {
                    "$ref": "#/definitions/models.DataInfo"
                },
                "latest_build_version": {
                    "type": "integer"
                },
                "minimum_build_version": {
                    "type": "integer"
                }
            }
        },
        "models.DataInfo": {
            "type": "object",
            "properties": {
                "district_data_last_updated_at": {
                    "type": "string"
                },
                "region_data_last_updated_at": {
                    "type": "string"
                },
                "sports_center_data_last_updated_at": {
                    "type": "string"
                }
            }
        },
        "models.District": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name_en": {
                    "type": "string"
                },
                "name_zh": {
                    "type": "string"
                },
                "region_id": {
                    "type": "integer"
                }
            }
        },
        "models.Region": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name_en": {
                    "type": "string"
                },
                "name_zh": {
                    "type": "string"
                }
            }
        },
        "models.SportsCenter": {
            "type": "object",
            "properties": {
                "address_en": {
                    "type": "string"
                },
                "address_zh": {
                    "type": "string"
                },
                "district_id": {
                    "type": "integer"
                },
                "external_id": {
                    "type": "integer"
                },
                "hourly_quota": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "latitude_dd": {
                    "type": "string"
                },
                "latitude_dms": {
                    "type": "string"
                },
                "longitude_dd": {
                    "type": "string"
                },
                "longitude_dms": {
                    "type": "string"
                },
                "monthly_quota": {
                    "type": "integer"
                },
                "name_en": {
                    "type": "string"
                },
                "name_zh": {
                    "type": "string"
                },
                "phone_numbers": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type in format of \"Bearer --TOKEN--\".",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Let's go gym API",
	Description:      "Let's go gym API endpoints.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
