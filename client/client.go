package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-deepseek/deepseek/config"
	"github.com/go-deepseek/deepseek/internal"
	"github.com/go-deepseek/deepseek/request"
	"github.com/go-deepseek/deepseek/response"
)

type Client struct { // TODO: VN -- move to internal pkg
	*http.Client
	config.Config
	BaseURL string
}

func NewClient(config config.Config) (*Client, error) {
	if config.ApiKey == "" {
		return nil, errors.New("err: api key should not be blank")
	}
	if config.TimeoutSeconds == 0 {
		return nil, errors.New("err: timeout seconds should not be 0")
	}

	c := &Client{
		Config: config,
		Client: &http.Client{
			Timeout: time.Second * time.Duration(config.TimeoutSeconds),
		},
		BaseURL: config.BaseURL,
	}
	return c, nil
}

func (c *Client) CallChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	// validate request
	if ctx == nil {
		ctx = context.Background()
	}
	if err := validateChatParams(chatReq, false, "deepseek-chat"); err != nil {
		return nil, err
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(ctx, chatReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	body, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, errors.New("err: service unavailable")
	}

	chatResp := &response.ChatCompletionsResponse{}
	err = json.Unmarshal(body, chatResp)
	return chatResp, err
}

func (c *Client) StreamChatCompletionsChat(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error) {
	// validate request
	if ctx == nil {
		ctx = context.Background()
	}
	if err := validateChatParams(chatReq, true, "deepseek-chat"); err != nil {
		return nil, err
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(ctx, chatReq)
	if err != nil {
		return nil, err
	}

	sr := response.NewStreamReader(respBody)
	return sr, nil
}

func (c *Client) CallChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (*response.ChatCompletionsResponse, error) {
	// validate request
	if ctx == nil {
		ctx = context.Background()
	}
	if err := validateChatParams(chatReq, false, "deepseek-reasoner"); err != nil {
		return nil, err
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(ctx, chatReq)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	body, err := io.ReadAll(respBody)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, errors.New("err: service unavailable")
	}

	chatResp := &response.ChatCompletionsResponse{}
	err = json.Unmarshal(body, chatResp)
	return chatResp, err
}

func (c *Client) StreamChatCompletionsReasoner(ctx context.Context, chatReq *request.ChatCompletionsRequest) (response.StreamReader, error) {
	// validate request
	if ctx == nil {
		ctx = context.Background()
	}
	if err := validateChatParams(chatReq, true, "deepseek-reasoner"); err != nil {
		return nil, err
	}
	if !c.DisableRequestValidation {
		err := request.ValidateChatCompletionsRequest(chatReq)
		if err != nil {
			return nil, err
		}
	}

	// call api
	respBody, err := c.do(ctx, chatReq)
	if err != nil {
		return nil, err
	}

	sr := response.NewStreamReader(respBody)
	return sr, nil
}

func (c *Client) PingChatCompletions(ctx context.Context, inputMessage string) (outputMessge string, err error) {
	chatReq := &request.ChatCompletionsRequest{
		Model:  "deepseek-chat",
		Stream: false,
		Messages: []*request.Message{
			{
				Role:    "user",
				Content: inputMessage,
			},
		},
	}
	chatResp, err := c.CallChatCompletionsChat(context.Background(), chatReq)
	if err != nil {
		return "", err
	}

	if chatResp != nil && len(chatResp.Choices) > 0 && chatResp.Choices[0].Message != nil {
		outputMessge = chatResp.Choices[0].Message.Content
	} else {
		return "", errors.New("err: invalid response")
	}
	return outputMessge, nil
}

func (c *Client) do(ctx context.Context, chatReq *request.ChatCompletionsRequest) (io.ReadCloser, error) {
	url := fmt.Sprintf(`%s/chat/completions`, c.BaseURL)

	in := new(bytes.Buffer)
	err := json.NewEncoder(in).Encode(chatReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, in)
	if err != nil {
		return nil, err
	}
	setDefaultHeaders(req, c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		return nil, processError(resp.Body, resp.StatusCode)
	}

	return resp.Body, nil
}

func setDefaultHeaders(req *http.Request, apiKey string) {
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, apiKey))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func processError(respBody io.Reader, statusCode int) error {
	errBody, err := io.ReadAll(respBody)
	if err != nil {
		return err
	}
	errResp, err := internal.ParseError(errBody)
	if err != nil {
		return fmt.Errorf("err: %s; http_status_code=%d", errBody, statusCode)
	}
	return fmt.Errorf("err: %s; http_status_code=%d", errResp.Error.Message, statusCode)
}

func validateChatParams(chatReq *request.ChatCompletionsRequest, wantStream bool, wantModel string) error {
	if chatReq == nil {
		return errors.New("err: chat completions request should not be nil")
	}
	if chatReq.Stream != wantStream {
		return fmt.Errorf(`err: stream should be %v`, wantStream)
	}
	if chatReq.Model != wantModel {
		return fmt.Errorf(`err: model should be %q`, wantModel)
	}
	return nil
}
