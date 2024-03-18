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
            "url": "http://cloud-barista.github.io",
            "email": "contact-to-cloud-barista@googlegroups.com"
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
        "/health": {
            "get": {
                "description": "Check API server is running",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[System] Utility"
                ],
                "summary": "Check API server is running",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/httpVersion": {
            "get": {
                "description": "Checks and logs the HTTP version of the incoming request to the server console.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[System] Utility"
                ],
                "summary": "Check HTTP version of incoming request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}": {
            "delete": {
                "description": "Clear the entire directories and configuration files",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[ResourceGroup] Resource group"
                ],
                "summary": "Clear the entire directories and configuration files",
                "parameters": [
                    {
                        "type": "string",
                        "default": "rg-01",
                        "description": "Resource group ID",
                        "name": "ResourceGroupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}/vpn/gcp-aws": {
            "post": {
                "description": "Create network resources for VPN tunnel in GCP and AWS",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Create network resources for VPN tunnel in GCP and AWS",
                "parameters": [
                    {
                        "type": "string",
                        "default": "rg-01",
                        "description": "Resource group ID",
                        "name": "ResourceGroupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Destroy network resources that were used to configure GCP as an AWS VPN tunnel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Destroy network resources that were used to configure GCP as an AWS VPN tunnel",
                "parameters": [
                    {
                        "type": "string",
                        "default": "rg-01",
                        "description": "Resource group ID",
                        "name": "ResourceGroupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}/vpn/gcp-aws/": {
            "delete": {
                "description": "Clear the entire directory and configuration files",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Clear the entire directory and configuration files",
                "parameters": [
                    {
                        "type": "string",
                        "default": "rg-01",
                        "description": "Resource group ID",
                        "name": "ResourceGroupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}/vpn/gcp-aws/blueprint": {
            "post": {
                "description": "Create a blueprint to configure GCP to AWS VPN tunnels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Create a blueprint to configure GCP to AWS VPN tunnels",
                "parameters": [
                    {
                        "description": "Parameters requied to create a blueprint to configure GCP to AWS VPN tunnels",
                        "name": "ParamsForBlueprint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateBluprintOfGcpAwsVpnRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}/vpn/gcp-aws/init": {
            "post": {
                "description": "Initialize GCP and AWS to configure VPN tunnels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Initialize GCP and AWS to configure VPN tunnels",
                "parameters": [
                    {
                        "description": "Init GCP and AWS",
                        "name": "InitCSPs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.InitGcpAndAwsForVpnTunnelRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}/vpn/gcp-aws/plan": {
            "post": {
                "description": "Show changes required by the current blueprint to configure GCP to AWS VPN tunnels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Show changes required by the current blueprint to configure GCP to AWS VPN tunnels",
                "parameters": [
                    {
                        "type": "string",
                        "default": "rg-01",
                        "description": "Resource group ID",
                        "name": "ResourceGroupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/rg/{resourceGroupId}/vpn/gcp-aws/state": {
            "get": {
                "description": "Get the current state of a saved plan to configure GCP to AWS VPN tunnels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[VPN] GCP to AWS VPN tunnel configuration"
                ],
                "summary": "Get the current state of a saved plan to configure GCP to AWS VPN tunnels",
                "parameters": [
                    {
                        "type": "string",
                        "default": "rg-01",
                        "description": "Resource group ID",
                        "name": "ResourceGroupId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/sample/users": {
            "get": {
                "description": "Get information of all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Get a list of users",
                "responses": {
                    "200": {
                        "description": "(sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUsersResponse"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/sample/users/{id}": {
            "get": {
                "description": "Get information of a user with a specific ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Get specific user information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetUserResponse"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User information to update",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deletion successful",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "patch": {
                "description": "Patch a user with the given information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Sample] Users"
                ],
                "summary": "Patch a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User information to update",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.PatchUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "(Sample) This is a sample description for success response in Swagger UI",
                        "schema": {
                            "$ref": "#/definitions/handlers.PatchUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "User Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/tofuVersion": {
            "get": {
                "description": "Check Tofu version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[System] Utility"
                ],
                "summary": "Check Tofu version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CreateBluprintOfGcpAwsVpnRequest": {
            "type": "object",
            "properties": {
                "resourceGroupId": {
                    "type": "string",
                    "default": "rg-01"
                },
                "tfVars": {
                    "$ref": "#/definitions/models.TfVarsGcpAwsVpnTunnel"
                }
            }
        },
        "handlers.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.GetUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.GetUsersResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MyUser"
                    }
                }
            }
        },
        "handlers.InitGcpAndAwsForVpnTunnelRequest": {
            "type": "object",
            "properties": {
                "resourceGroupId": {
                    "type": "string",
                    "default": "rg-01"
                }
            }
        },
        "handlers.PatchUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.PatchUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.UpdateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.MyUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "text": {
                    "type": "string",
                    "example": "Any text"
                }
            }
        },
        "models.TfVarsGcpAwsVpnTunnel": {
            "type": "object",
            "properties": {
                "my-imported-aws-subnet-id": {
                    "type": "string"
                },
                "my-imported-aws-vpc-id": {
                    "type": "string"
                },
                "my-imported-gcp-subnet-name": {
                    "type": "string"
                },
                "my-imported-gcp-vpc-name": {
                    "type": "string"
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
	Version:          "latest",
	Host:             "",
	BasePath:         "/mc-net",
	Schemes:          []string{},
	Title:            "POC-MC-Net-TF REST API",
	Description:      "POC-MC-Net-TF REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
