日程管理

接口文档

```
{
  "openapi": "3.0.1",
  "info": {
    "title": "默认模块",
    "description": "",
    "version": "1.0.0"
  },
  "tags": [],
  "paths": {
    "/schedule/query": {
      "post": {
        "summary": "查询某一天的日程",
        "deprecated": false,
        "description": "",
        "tags": [],
        "parameters": [],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "user_id": {
                    "type": "integer"
                  },
                  "year": {
                    "type": "integer"
                  },
                  "month": {
                    "type": "integer"
                  },
                  "day": {
                    "type": "integer"
                  }
                },
                "required": [
                  "user_id",
                  "year",
                  "month",
                  "day"
                ]
              },
              "example": {
                "user_id": 1,
                "year": 2025,
                "month": 11,
                "day": 29
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "integer"
                    },
                    "data": {
                      "type": "array",
                      "items": {
                        "type": "object",
                        "properties": {
                          "id": {
                            "type": "integer"
                          },
                          "user_id": {
                            "type": "integer"
                          },
                          "create_at": {
                            "type": "string"
                          },
                          "update_at": {
                            "type": "string"
                          },
                          "year": {
                            "type": "integer"
                          },
                          "month": {
                            "type": "integer"
                          },
                          "day": {
                            "type": "integer"
                          },
                          "start_time": {
                            "type": "string"
                          },
                          "end_time": {
                            "type": "string"
                          },
                          "content": {
                            "type": "string"
                          },
                          "priority": {
                            "type": "integer"
                          },
                          "status": {
                            "type": "integer"
                          }
                        }
                      }
                    },
                    "msg": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "code",
                    "data",
                    "msg"
                  ]
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
      }
    },
    "/schedule/store": {
      "post": {
        "summary": "新增日程",
        "deprecated": false,
        "description": "",
        "tags": [],
        "parameters": [],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "year": {
                    "type": "integer"
                  },
                  "month": {
                    "type": "integer"
                  },
                  "day": {
                    "type": "integer"
                  },
                  "user_id": {
                    "type": "integer"
                  },
                  "content": {
                    "type": "string"
                  },
                  "start": {
                    "type": "string"
                  },
                  "end": {
                    "type": "string"
                  },
                  "priority": {
                    "type": "integer"
                  }
                },
                "required": [
                  "year",
                  "month",
                  "day",
                  "user_id",
                  "content",
                  "start",
                  "end",
                  "priority"
                ]
              },
              "example": {
                "year": 2025,
                "month": 12,
                "day": 1,
                "user_id": 1,
                "content": "",
                "start": "",
                "end": ""
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "integer"
                    },
                    "data": {
                      "type": "null"
                    },
                    "msg": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "code",
                    "data",
                    "msg"
                  ]
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
      }
    },
    "/schedule/update": {
      "post": {
        "summary": "更新日程",
        "deprecated": false,
        "description": "",
        "tags": [],
        "parameters": [],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "id": {
                    "type": "integer"
                  },
                  "year": {
                    "type": "integer"
                  },
                  "month": {
                    "type": "integer"
                  },
                  "day": {
                    "type": "integer"
                  },
                  "start": {
                    "type": "string"
                  },
                  "end": {
                    "type": "string"
                  },
                  "content": {
                    "type": "string"
                  },
                  "status": {
                    "type": "integer"
                  },
                  "user_id": {
                    "type": "integer"
                  },
                  "priority": {
                    "type": "integer"
                  }
                },
                "required": [
                  "id",
                  "year",
                  "month",
                  "day",
                  "start",
                  "end",
                  "content",
                  "status",
                  "user_id",
                  "priority"
                ]
              },
              "example": {
                "id": 1,
                "year": 1,
                "month": 1,
                "day": 1,
                "start": "",
                "end": "",
                "content": "",
                "status": 1
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "code": {
                      "type": "integer"
                    },
                    "data": {
                      "type": "null"
                    },
                    "msg": {
                      "type": "string"
                    }
                  },
                  "required": [
                    "code",
                    "data",
                    "msg"
                  ]
                }
              }
            },
            "headers": {}
          }
        },
        "security": []
      }
    }
  },
  "components": {
    "schemas": {},
    "responses": {},
    "securitySchemes": {}
  },
  "servers": [],
  "security": []
}
```