package deepseek_test

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/config"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const DEEPSEEK_API_KEY = `sk-123cd456c78d9be0b123de45cf6789b0` // replace with valid one

//go:embed testdata/*
var testdata embed.FS

func GetApiKey() string {
	if apiKey := os.Getenv("DEEPSEEK_API_KEY"); apiKey != "" {
		return apiKey
	}
	return DEEPSEEK_API_KEY
}

func TestPingChatCompletions(t *testing.T) {
	client, err := deepseek.NewClient(GetApiKey(), "https://custom.deepseek.com")
	require.NoError(t, err)

	output, err := client.PingChatCompletions(context.Background(), "Hello")

	assert.NoError(t, err)
	assert.NotEmpty(t, output)
}

func TestCallChat(t *testing.T) {
	// ts := NewFakeServer("testdata/01_resp_basic_chat.json")
	// defer ts.Close()

	client, err := deepseek.NewClient(GetApiKey(), "https://custom.deepseek.com")
	require.NoError(t, err)

	reqJson, err := testdata.ReadFile("testdata/01_req_basic_chat.json")
	require.NoError(t, err)
	req := &request.ChatCompletionsRequest{}
	err = json.Unmarshal(reqJson, req)
	require.NoError(t, err)

	ctx := context.Background()
	// ctx, _ = context.WithTimeout(ctx, time.Second*2)
	resp, err := client.CallChatCompletionsChat(ctx, req) // test

	require.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
}

func TestStreamChat(t *testing.T) {
	// ts := NewFakeServer("testdata/02_resp_stream_chat.json")
	// defer ts.Close()

	client, err := deepseek.NewClient(GetApiKey(), "https://custom.deepseek.com")
	require.NoError(t, err)

	reqJson, err := testdata.ReadFile("testdata/02_req_stream_chat.json")
	require.NoError(t, err)
	req := &request.ChatCompletionsRequest{}
	err = json.Unmarshal(reqJson, req)
	require.NoError(t, err)

	iter, err := client.StreamChatCompletionsChat(context.Background(), req) // test

	require.NoError(t, err)
	assert.NotNil(t, iter)

	for {
		resp, err := iter.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Id)
		fmt.Print(resp.Choices[0].Delta.Content)
	}
}

func TestCallReasoner(t *testing.T) {
	// ts := NewFakeServer("testdata/03_resp_basic_reasoner.json")
	// defer ts.Close()

	client, err := deepseek.NewClient(GetApiKey(), "https://custom.deepseek.com")
	require.NoError(t, err)

	reqJson, err := testdata.ReadFile("testdata/03_req_basic_reasoner.json")
	require.NoError(t, err)
	req := &request.ChatCompletionsRequest{}
	err = json.Unmarshal(reqJson, req)
	require.NoError(t, err)

	resp, err := client.CallChatCompletionsReasoner(context.Background(), req) // test

	require.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
}

func TestStreamReasoner(t *testing.T) {
	// ts := NewFakeSteamServer("testdata/04_resp_stream_reasoner.json")
	// defer ts.Close()

	config := config.Config{
		ApiKey:         GetApiKey(),
		TimeoutSeconds: 120,
		BaseURL:        "https://custom.deepseek.com",
	}
	client, err := deepseek.NewClientWithConfig(config)
	require.NoError(t, err)

	reqJson, err := testdata.ReadFile("testdata/04_req_stream_reasoner.json")
	require.NoError(t, err)
	req := &request.ChatCompletionsRequest{}
	err = json.Unmarshal(reqJson, req)
	require.NoError(t, err)

	iter, err := client.StreamChatCompletionsReasoner(context.Background(), req) // test

	require.NoError(t, err)
	assert.NotNil(t, iter)

	for {
		resp, err := iter.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Id)
		if resp.Choices[0].Delta.Content != "" {
			fmt.Print(resp.Choices[0].Delta.Content)
		} else {
			fmt.Print(resp.Choices[0].Delta.ReasoningContent)
		}
	}
}

func TestResponse(t *testing.T) {
	respJson, err := testdata.ReadFile("testdata/51_full_response.json")
	require.NoError(t, err)

	resp := &response.ChatCompletionsResponse{}
	json.Unmarshal(respJson, resp)

	wantStr := "dummy_string"
	wantInt := 1

	assert.NotNil(t, resp)

	assert.Equal(t, resp.Id, wantStr)
	assert.Len(t, resp.Choices, 1)
	assert.Equal(t, resp.Created, wantInt)
	assert.Equal(t, resp.Model, wantStr)
	assert.Equal(t, resp.SystemFingerprint, wantStr)
	assert.Equal(t, resp.Object, "chat.completion")
	assert.NotNil(t, resp.Usage)

	choice := resp.Choices[0]
	assert.Equal(t, choice.FinishReason, "stop")
	assert.Equal(t, choice.Index, wantInt)
	assert.NotNil(t, choice.Message)

	message := choice.Message
	assert.Equal(t, message.Content, wantStr)
	assert.Equal(t, message.Role, "assistant")
	// TODO: VN -- complete reasoning_content

	usage := resp.Usage
	assert.Equal(t, usage.CompletionTokens, wantInt)
	assert.Equal(t, usage.PromptTokens, wantInt)
	assert.Equal(t, usage.PromptCacheHitTokens, wantInt)
	assert.Equal(t, usage.PromptCacheMissTokens, wantInt)
	assert.Equal(t, usage.TotalTokens, wantInt)

	// TODO: VN -- enable below asserts
	// assert.NotNil(t, usage.PromptTokensDetails)
	// tokenDetails := usage.PromptTokensDetails
	// assert.Equal(t, tokenDetails.CachedTokens, wantInt)
}
