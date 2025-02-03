### Error handling
- error-code: https://api-docs.deepseek.com/quick_start/error_codes

code=401
```
{
  "error": {
    "message": "Authentication Fails (no such user)",
    "type": "authentication_error",
    "param": null,
    "code": "invalid_request_error"
  }
}
```

code=400
```
{
  "error": {
    "message": "Invalid frequency_penalty value, the valid range of frequency_penalty is [-2, 2]",
    "type": "invalid_request_error",
    "param": null,
    "code": "invalid_request_error"
  }
}
```

code=422 (when model not set)
```
Failed to deserialize the JSON body into the target type: missing field `model` at line 1 column 164
```

### Rate limit
- rate-limit: https://api-docs.deepseek.com/quick_start/rate_limit
- Non-streaming requests: Continuously return empty lines
- Streaming requests: Continuously return SSE keep-alive comments (: keep-alive)

### Sample Payload
```
{
  "messages": [
    {
      "content": "You are a helpful assistant",
      "role": "system"
    },
    {
      "content": "Hi",
      "role": "user"
    }
  ],
  "model": "deepseek-chat",
  "frequency_penalty": 1.1,
  "max_tokens": 2048,
  "presence_penalty": 0,
  "response_format": {
    "type": "text"
  },
  "stream": false,
  "temperature": 1,
  "top_p": 1,
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_weather",
        "parameters": {
          "type": "object",
          "properties": {
            "location": {
              "type": "string",
              "description": "The city and state, e.g. San Francisco, CA"
            }
          },
          "required": [
            "location"
          ]
        }
      }
    }
  ],
  "tool_choice": {"type": "function", "function": {"name": "get_weather"}},
  "logprobs": false,
  "top_logprobs": null
}
```