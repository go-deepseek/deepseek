package response

// ChatCompletionsResponse is response payload for `POST /chat/completions` API.
type ChatCompletionsResponse struct {
	Id                string    `json:"id"`
	Choices           []*Choice `json:"choices"`
	Created           int       `json:"created"`
	Model             string    `json:"model"`
	SystemFingerprint string    `json:"system_fingerprint"`
	Object            string    `json:"object"`
	Usage             *Usage    `json:"usage,omitempty"`
}

type Choice struct {
	FinishReason string    `json:"finish_reason"`
	Index        int       `json:"index"`
	Message      *Message  `json:"message"`
	Delta        *Delta    `json:"delta"`
	Logprobs     *Logprobs `json:"logprobs"`
}

type Message struct {
	Role             string      `json:"role"`
	Content          string      `json:"content"`
	ReasoningContent string      `json:"reasoning_content"`
	ToolCalls        []*ToolCall `json:"tool_calls"`
}

type ToolCall struct {
	Id       string       `json:"id"`
	Type     string       `json:"type"`
	Function ToolFunction `json:"function"`
}

type ToolFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type Delta struct {
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content"`
}

type Usage struct {
	CompletionTokens        int                     `json:"completion_tokens"`
	PromptTokens            int                     `json:"prompt_tokens"`
	PromptCacheHitTokens    int                     `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens   int                     `json:"prompt_cache_miss_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type PromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type Logprobs struct {
	Content *[]Content `json:"content"`
}

type Content struct {
	TopLogprob
	TopLogprobs []*TopLogprob `json:"top_logprobs"`
}

type TopLogprob struct {
	Token   string `json:"token"`
	Logprob int    `json:"logprob"`
	Bytes   []int  `json:"bytes"`
}
