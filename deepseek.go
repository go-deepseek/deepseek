package deepseek

import (
	"context"

	"github.com/go-deepseek/deepseek/client"
	"github.com/go-deepseek/deepseek/config"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

const DEFAULT_TIMEOUT_SECONDS = 60

const (
	DEEPSEEK_CHAT_MODEL     = "deepseek-chat"
	DEEPSEEK_REASONER_MODEL = "deepseek-reasoner"
)

type Client interface {
	// CallChatCompletionsChat calls chat api with model=deepseek-chat and stream=false.
	// It returns response from DeepSeek-V3 model.
	CallChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)

	// CallChatCompletionsReasoner calls chat api with model=deepseek-reasoner and stream=false.
	// It returns response from DeepSeek-R1 model.
	CallChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)

	// StreamChatCompletionsChat calls chat api with model=deepseek-chat and stream=true.
	// It returns response from DeepSeek-V3 model.
	StreamChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)

	// StreamChatCompletionsChat calls chat api with model=deepseek-reasoner and stream=true.
	// It returns response from DeepSeek-R1 model.
	StreamChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)

	// PingChatCompletions is a ping to check go deepseek client is working fine.
	PingChatCompletions(ctx context.Context, inputMessage string) (outputMessge string, err error)
}

// NewClient creates deeseek client with given api key.
func NewClient(apiKey string) (Client, error) {
	config := config.Config{
		ApiKey:         apiKey,
		TimeoutSeconds: DEFAULT_TIMEOUT_SECONDS,
	}
	return NewClientWithConfig(config)
}

// NewClient creates deeseek client with given client config.
func NewClientWithConfig(config config.Config) (Client, error) {
	return client.NewClient(config)
}
