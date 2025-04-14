package fake

import (
	"context"

	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

type Callbacks struct {
	CallChatCompletionsChatCallback     func(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)
	CallChatCompletionsReasonerCallback func(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error)

	StreamChatCompletionsChatCallback     func(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)
	StreamChatCompletionsReasonerCallback func(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error)

	PingChatCompletionsCallback func(ctx context.Context, inputMessage string) (outputMessage string, err error)
}

type FakeCallbackClient struct {
	callbacks Callbacks
}

func NewFakeCallbackClient(callbacks Callbacks) *FakeCallbackClient {
	fc := &FakeCallbackClient{
		callbacks: callbacks,
	}
	return fc
}

func (c *FakeCallbackClient) CallChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	if c.callbacks.CallChatCompletionsChatCallback == nil {
		panic("err: CallChatCompletionsChatCallback is nil")
	}
	return c.callbacks.CallChatCompletionsChatCallback(ctx, chatReq)
}

func (c *FakeCallbackClient) CallChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	if c.callbacks.CallChatCompletionsReasonerCallback == nil {
		panic("err: CallChatCompletionsReasonerCallback is nil")
	}
	return c.callbacks.CallChatCompletionsReasonerCallback(ctx, chatReq)
}

func (c *FakeCallbackClient) StreamChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error) {
	if c.callbacks.StreamChatCompletionsChatCallback == nil {
		panic("err: StreamChatCompletionsChatCallback is nil")
	}
	return c.callbacks.StreamChatCompletionsChatCallback(ctx, chatReq)
}

func (c *FakeCallbackClient) StreamChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error) {
	if c.callbacks.StreamChatCompletionsReasonerCallback == nil {
		panic("err: StreamChatCompletionsReasonerCallback is nil")
	}
	return c.callbacks.StreamChatCompletionsReasonerCallback(ctx, chatReq)
}

func (c *FakeCallbackClient) PingChatCompletions(ctx context.Context, inputMessage string) (outputMessage string, err error) {
	if c.callbacks.PingChatCompletionsCallback == nil {
		panic("err: PingChatCompletionsCallback is nil")
	}
	return c.callbacks.PingChatCompletionsCallback(ctx, inputMessage)
}
