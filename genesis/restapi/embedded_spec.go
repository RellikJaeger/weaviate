//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Weaviate Genesis Server is used to bootstrap the P2P network of Weaviate instances",
    "title": "Weaviate Genesis Server",
    "contact": {
      "name": "Weaviate",
      "url": "https://github.com/semi-technologies",
      "email": "hello@semi.technology"
    },
    "version": "0.1.0"
  },
  "paths": {
    "/peers": {
      "get": {
        "description": "List the registered peers",
        "operationId": "genesis.peers.list",
        "responses": {
          "200": {
            "description": "The list of registered peers",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Peer"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/peers/register": {
      "post": {
        "description": "Register a new Weaviate peer in the network",
        "operationId": "genesis.peers.register",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PeerUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully registred the peer to the network.",
            "schema": {
              "$ref": "#/definitions/PeerRegistrationResponse"
            }
          },
          "400": {
            "description": "The weaviate peer is not reachable from the Gensis service."
          },
          "403": {
            "description": "You are not allowed on the network."
          }
        }
      }
    },
    "/peers/{peerId}": {
      "delete": {
        "description": "Leave the weaviate network",
        "operationId": "genesis.peers.leave",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Name of the Weaviate peer",
            "name": "peerId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful left the network."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no such peer was found."
          }
        }
      }
    },
    "/peers/{peerId}/ping": {
      "post": {
        "description": "Ping the Genesis server, to make mark the peer as alive and udpate schema info",
        "operationId": "genesis.peers.ping",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Name of the Weaviate peer",
            "name": "peerId",
            "in": "path",
            "required": true
          },
          {
            "description": "Request Body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PeerPing"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ping received"
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no such peer was found."
          }
        }
      }
    }
  },
  "definitions": {
    "Contextionary": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "object",
          "properties": {
            "algorithm": {
              "description": "Hash algorithm",
              "type": "string"
            },
            "value": {
              "description": "The actual hash",
              "type": "string"
            }
          }
        },
        "url": {
          "description": "URL where the contextionary can be found",
          "type": "string",
          "format": "url"
        }
      }
    },
    "Peer": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/PeerUpdate"
        },
        {
          "type": "object",
          "properties": {
            "id": {
              "description": "Unique ID of this peer registration, will be updated if the peer conntects again to the network.",
              "type": "string",
              "format": "uuid"
            },
            "last_contact_at": {
              "description": "When we were received a ping from this peer from the last time",
              "type": "integer",
              "format": "int64"
            },
            "schema_hash": {
              "description": "The latest known hash of the local schema of the peer",
              "type": "string"
            }
          }
        }
      ]
    },
    "PeerPing": {
      "type": "object",
      "properties": {
        "schemaHash": {
          "description": "(base64 encoded) hash of the current schema",
          "type": "string",
          "format": "md5",
          "example": "59d41240e1b7024b6cdc1206696e62d2"
        }
      }
    },
    "PeerRegistrationResponse": {
      "type": "object",
      "properties": {
        "contextionary": {
          "$ref": "#/definitions/Contextionary"
        },
        "peer": {
          "$ref": "#/definitions/Peer"
        }
      }
    },
    "PeerUpdate": {
      "type": "object",
      "properties": {
        "peerName": {
          "description": "Name of the peer, must be valid DNS name",
          "type": "string"
        },
        "peerUri": {
          "description": "Host or IP of the peer, defaults to peerName",
          "type": "string",
          "format": "uri"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The Weaviate Genesis Server is used to bootstrap the P2P network of Weaviate instances",
    "title": "Weaviate Genesis Server",
    "contact": {
      "name": "Weaviate",
      "url": "https://github.com/semi-technologies",
      "email": "hello@semi.technology"
    },
    "version": "0.1.0"
  },
  "paths": {
    "/peers": {
      "get": {
        "description": "List the registered peers",
        "operationId": "genesis.peers.list",
        "responses": {
          "200": {
            "description": "The list of registered peers",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Peer"
              }
            }
          },
          "500": {
            "description": "Internal error"
          }
        }
      }
    },
    "/peers/register": {
      "post": {
        "description": "Register a new Weaviate peer in the network",
        "operationId": "genesis.peers.register",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PeerUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successfully registred the peer to the network.",
            "schema": {
              "$ref": "#/definitions/PeerRegistrationResponse"
            }
          },
          "400": {
            "description": "The weaviate peer is not reachable from the Gensis service."
          },
          "403": {
            "description": "You are not allowed on the network."
          }
        }
      }
    },
    "/peers/{peerId}": {
      "delete": {
        "description": "Leave the weaviate network",
        "operationId": "genesis.peers.leave",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Name of the Weaviate peer",
            "name": "peerId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful left the network."
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no such peer was found."
          }
        }
      }
    },
    "/peers/{peerId}/ping": {
      "post": {
        "description": "Ping the Genesis server, to make mark the peer as alive and udpate schema info",
        "operationId": "genesis.peers.ping",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Name of the Weaviate peer",
            "name": "peerId",
            "in": "path",
            "required": true
          },
          {
            "description": "Request Body",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PeerPing"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ping received"
          },
          "401": {
            "description": "Unauthorized or invalid credentials."
          },
          "403": {
            "description": "The used API-key has insufficient permissions."
          },
          "404": {
            "description": "Successful query result but no such peer was found."
          }
        }
      }
    }
  },
  "definitions": {
    "Contextionary": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "object",
          "properties": {
            "algorithm": {
              "description": "Hash algorithm",
              "type": "string"
            },
            "value": {
              "description": "The actual hash",
              "type": "string"
            }
          }
        },
        "url": {
          "description": "URL where the contextionary can be found",
          "type": "string",
          "format": "url"
        }
      }
    },
    "ContextionaryHash": {
      "type": "object",
      "properties": {
        "algorithm": {
          "description": "Hash algorithm",
          "type": "string"
        },
        "value": {
          "description": "The actual hash",
          "type": "string"
        }
      }
    },
    "Peer": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/PeerUpdate"
        },
        {
          "type": "object",
          "properties": {
            "id": {
              "description": "Unique ID of this peer registration, will be updated if the peer conntects again to the network.",
              "type": "string",
              "format": "uuid"
            },
            "last_contact_at": {
              "description": "When we were received a ping from this peer from the last time",
              "type": "integer",
              "format": "int64"
            },
            "schema_hash": {
              "description": "The latest known hash of the local schema of the peer",
              "type": "string"
            }
          }
        }
      ]
    },
    "PeerPing": {
      "type": "object",
      "properties": {
        "schemaHash": {
          "description": "(base64 encoded) hash of the current schema",
          "type": "string",
          "format": "md5",
          "example": "59d41240e1b7024b6cdc1206696e62d2"
        }
      }
    },
    "PeerRegistrationResponse": {
      "type": "object",
      "properties": {
        "contextionary": {
          "$ref": "#/definitions/Contextionary"
        },
        "peer": {
          "$ref": "#/definitions/Peer"
        }
      }
    },
    "PeerUpdate": {
      "type": "object",
      "properties": {
        "peerName": {
          "description": "Name of the peer, must be valid DNS name",
          "type": "string"
        },
        "peerUri": {
          "description": "Host or IP of the peer, defaults to peerName",
          "type": "string",
          "format": "uri"
        }
      }
    }
  }
}`))
}