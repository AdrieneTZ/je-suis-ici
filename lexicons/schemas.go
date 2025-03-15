package lexicons

/*
Define UserSchema Lexicon
*/
var UserSchema = `{
  "lexicon": 1,
  "id": "app.checkin.user",
  "type": "record",
  "description": "user's account record",
  "record": {
    "type": "object",
    "required": ["did", "handle", "displayName"],
    "properties": {
      "did": {
        "type": "string",
        "format": "did"
      },
      "handle": {
        "type": "string",
        "format": "handle"
      },
      "displayName": {
        "type": "string",
        "maxLength": 64
      },
      "avatar": {
        "type": "string",
        "format": "uri"
      },
      "createdAt": {
        "type": "string",
        "format": "datetime"
      },
	  "updatedAt": {
		"type": "string",
		"format": "datetime"
	  }
    }
  }
}`

/*
Define CheckInSchema Lexicon
*/
var CheckInSchema = `{
  "lexicon": 1,
  "id": "app.checkin.post",
  "type": "record",
  "description": "user's check-in record",
  "record": {
    "type": "object",
    "required": ["userDid", "location", "createdAt"],
    "properties": {
      "id": {
        "type": "string"
      },
      "userDid": {
        "type": "string",
        "format": "did"
      },
      "text": {
        "type": "string",
        "maxLength": 5000
      },
      "location": {
        "type": "object",
        "required": ["latitude", "longitude", "locationName"],
        "properties": {
          "latitude": {
            "type": "number"
          },
          "longitude": {
            "type": "number"
          },
          "locationName": {
            "type": "string"
          }
        }
      },
      "images": {
        "type": "array",
        "items": {
          "type": "string",
          "format": "uri"
        }
      },
      "createdAt": {
        "type": "string",
        "format": "datetime"
      },
	  "updatedAt": {
        "type": "string",
        "format": "datetime"
      }
    }
  }
}`

/*
ATProtocolSchemas:
1. return all defined lexicon schemas
2. returned data type is map
*/
func ATProtocolSchemas() map[string]string {
	return map[string]string{
		"app.checkin.user": UserSchema,
		"app.checkin.post": CheckInSchema,
	}
}
