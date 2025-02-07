package request

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleTool      = "tool"
)

const (
	ResponseFormatText       = "text"
	ResponseFormatJsonObject = "json_object"
)

// ChatCompletionsRequest is request payload for `POST /chat/completions` API.
type ChatCompletionsRequest struct {
	Messages         []*Message      `json:"messages"`
	Model            string          `json:"model"`
	FrequencyPenalty float32         `json:"frequency_penalty,omitempty"`
	MaxTokens        int             `json:"max_tokens,omitempty"`
	PresencePenalty  int             `json:"presence_penalty,omitempty"`
	ResponseFormat   *ResponseFormat `json:"response_format,omitempty"`
	Stop             []string        `json:"stop,omitempty"`
	Stream           bool            `json:"stream,omitempty"`
	StreamOptions    *StreamOptions  `json:"stream_options,omitempty"`
	Temperature      *float32        `json:"temperature,omitempty"`
	TopP             *float32        `json:"top_p,omitempty"`
	Tools            *[]Tool         `json:"tools,omitempty"`
	ToolChoice       any             `json:"tool_choice,omitempty"`
	Logprobs         bool            `json:"logprobs,omitempty"`
	TopLogprobs      *int            `json:"top_logprobs,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name,omitempty"`
	// Prefix  bool   `json:"prefix"` // TODO: VN -- applicable for assistant role; support prefix while enabling beta support
	// ReasoningContent string `json:"reasoning_content"` // TODO: VN -- applicable for assistant role; support prefix while enabling beta support
	ToolCallId string `json:"tool_call_id"`
}

type ResponseFormat struct {
	Type string `json:"type"` // Must be one of text or json_object
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage"`
}

// "tools": [
//
//	{
//	  "type": "function",
//	  "function": {
//	    "name": "get_weather",
//	    "description": "Get weather of an location, the user shoud supply a location first",
//	    "parameters": {
//	      "type": "object",
//	      "properties": {
//	        "location": {
//	          "type": "string",
//	          "description": "The city and state, e.g. San Francisco, CA"
//	        }
//	      },
//	      "required": [
//	        "location"
//	      ]
//	    }
//	  }
//	}
//
// ]
type Tool struct {
	Type     string        `json:"type"`
	Function *ToolFunction `json:"function"`
}

type ToolFunction struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters"`
}

type ToolChoiceString string

type ToolChoiceNamed struct {
	Type     string             `json:"type"`
	Function ToolChoiceFunction `json:"function"`
}

type ToolChoiceFunction struct {
	Name string `json:"name"`
}
