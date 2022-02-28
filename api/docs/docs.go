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
        "contact": {},
        "license": {
            "name": "GNU General Public License v3.0",
            "url": "https://github.com/sundowndev/phoneinfoga/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "default"
                ],
                "summary": "Check if service is healthy.",
                "operationId": "healthCheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.healthResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    }
                }
            }
        },
        "/numbers": {
            "get": {
                "description": "This route is actually not used yet.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Numbers"
                ],
                "summary": "Fetch all previously scanned numbers.",
                "operationId": "getAllNumbers",
                "deprecated": true,
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.getAllNumbersResponse"
                        }
                    }
                }
            }
        },
        "/numbers/{number}/scan/googlesearch": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Numbers"
                ],
                "summary": "Perform a scan using Google Search engine.",
                "operationId": "googleSearchScan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Input phone number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.scanResultResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/scanners.GoogleSearchResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    }
                }
            }
        },
        "/numbers/{number}/scan/local": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Numbers"
                ],
                "summary": "Perform a scan using local phone number library.",
                "operationId": "localScan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Input phone number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.scanResultResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/scanners.Number"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    }
                }
            }
        },
        "/numbers/{number}/scan/numverify": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Numbers"
                ],
                "summary": "Perform a scan using Numverify's API.",
                "operationId": "numverifyScan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Input phone number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.scanResultResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/scanners.NumverifyScannerResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    }
                }
            }
        },
        "/numbers/{number}/scan/ovh": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Numbers"
                ],
                "summary": "Perform a scan using OVH's API.",
                "operationId": "ovhScan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Input phone number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.scanResultResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/scanners.OVHScannerResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    }
                }
            }
        },
        "/numbers/{number}/validate": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Numbers"
                ],
                "summary": "Check if a number is valid and possible.",
                "operationId": "validate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Input phone number",
                        "name": "number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.JSONResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.JSONResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "api.getAllNumbersResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "numbers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scanners.Number"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "api.healthResponse": {
            "type": "object",
            "properties": {
                "commit": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "api.scanResultResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "object"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "scanners.GoogleSearchDork": {
            "type": "object",
            "properties": {
                "dork": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "scanners.GoogleSearchResponse": {
            "type": "object",
            "properties": {
                "disposableProviders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scanners.GoogleSearchDork"
                    }
                },
                "general": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scanners.GoogleSearchDork"
                    }
                },
                "individuals": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scanners.GoogleSearchDork"
                    }
                },
                "reputation": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scanners.GoogleSearchDork"
                    }
                },
                "socialMedia": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/scanners.GoogleSearchDork"
                    }
                }
            }
        },
        "scanners.Number": {
            "type": "object",
            "properties": {
                "E164": {
                    "type": "string"
                },
                "carrier": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "countryCode": {
                    "type": "integer"
                },
                "international": {
                    "type": "string"
                },
                "local": {
                    "type": "string"
                },
                "rawLocal": {
                    "type": "string"
                }
            }
        },
        "scanners.NumverifyScannerResponse": {
            "type": "object",
            "properties": {
                "carrier": {
                    "type": "string"
                },
                "country_code": {
                    "type": "string"
                },
                "country_name": {
                    "type": "string"
                },
                "country_prefix": {
                    "type": "string"
                },
                "error": {
                    "$ref": "#/definitions/scanners.numverifyError"
                },
                "international_format": {
                    "type": "string"
                },
                "line_type": {
                    "type": "string"
                },
                "local_format": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                },
                "valid": {
                    "type": "boolean"
                }
            }
        },
        "scanners.OVHScannerResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "found": {
                    "type": "boolean"
                },
                "numberRange": {
                    "type": "string"
                },
                "zipCode": {
                    "type": "string"
                }
            }
        },
        "scanners.numverifyError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "info": {
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
	Version:     "v2",
	Host:        "demo.phoneinfoga.crvx.fr",
	BasePath:    "/api",
	Schemes:     []string{"http", "https"},
	Title:       "PhoneInfoga REST API",
	Description: "Advanced information gathering & OSINT framework for phone numbers.",
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
