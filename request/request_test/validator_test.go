package request_test

import (
	"fmt"
	"testing"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
	"github.com/stretchr/testify/assert"
)

func TestValidateChatCompletionsRequest(t *testing.T) {
	temp := float32(2)
	req := &request.ChatCompletionsRequest{
		Messages: []*request.Message{
			{
				Role:    request.RoleUser,
				Content: "Hello",
			},
		},
		Model:            deepseek.DEEPSEEK_CHAT_MODEL,
		FrequencyPenalty: 1,
		MaxTokens:        1,
		PresencePenalty:  1,
		ResponseFormat: &request.ResponseFormat{
			Type: request.ResponseFormatText,
		},
		Stop:   []string{"MOD1"},
		Stream: true,
		StreamOptions: &request.StreamOptions{
			IncludeUsage: true,
		},
		Temperature: &temp,
		// TopP: nil,	// TODO: VN -- pass non nil
	}
	err := request.ValidateChatCompletionsRequest(req)
	assert.NoError(t, err)
	fmt.Println(err)
}
