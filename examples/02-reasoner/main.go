package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

var apiKey = "your_deepseek_api_key"

func main() {
	if apiKeyEnv := os.Getenv("DEEPSEEK_API_KEY"); apiKeyEnv != "" {
		apiKey = apiKeyEnv
	}

	// create deepseek api client
	cli, err := deepseek.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	inputMessage := "Hello" // set your input message

	chatReq := &request.ChatCompletionsRequest{
		Model:  deepseek.DEEPSEEK_REASONER_MODEL,
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: inputMessage,
			},
		},
	}
	fmt.Printf("input message => %s\n", chatReq.Messages[0].Content)

	// call deepseek api
	chatResp, err := cli.CallChatCompletionsReasoner(context.Background(), chatReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("output reasoning => %s\n", chatResp.Choices[0].Message.ReasoningContent)
	fmt.Printf("output message => %s\n", chatResp.Choices[0].Message.Content)
}
