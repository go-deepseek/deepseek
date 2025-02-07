package request

import (
	"errors"
	"fmt"
)

var (
	modelChat     = "deepseek-chat"
	modelReasoner = "deepseek-reasoner"
)

var roles map[string]struct{}
var rolesStr string
var modelStr string
var responseFormatStr string

func init() {
	rolesStr = fmt.Sprintf(`%s, %s, %s, %s`, RoleSystem, RoleUser, RoleAssistant, RoleTool)

	roles = make(map[string]struct{})
	roles[RoleSystem] = struct{}{}
	roles[RoleUser] = struct{}{}
	roles[RoleAssistant] = struct{}{}
	roles[RoleTool] = struct{}{}

	modelStr = fmt.Sprintf(`%s, %s`, modelChat, modelReasoner)
	responseFormatStr = fmt.Sprintf(`%s, %s`, ResponseFormatText, ResponseFormatJsonObject)
}

func ValidateChatCompletionsRequest(req *ChatCompletionsRequest) error {
	if req == nil {
		return errors.New("err: input request is nil")
	}

	if err := validateMessages(req.Messages); err != nil {
		return err
	}

	if err := validateModel(req); err != nil {
		return err
	}

	if err := validateMultipleFields(req); err != nil {
		return err
	}

	if err := validateResponseFormat(req); err != nil {
		return err
	}

	if err := validateStreamOptions(req); err != nil {
		return err
	}

	// TODO: VN -- tools validation
	// TODO: VN -- tool_choice validation
	if req.TopLogprobs != nil {
		if !req.Logprobs {
			return fmt.Errorf(`err: top_logprobs can not be set when "logprobs" is false`)
		}
		if !(*req.TopLogprobs >= 0 && *req.TopLogprobs <= 20) {
			return fmt.Errorf(`err: top_logprobs is invalid; it should be number between 0 and 20`)
		}
	}
	return nil
}

func validateMessages(messages []*Message) error {
	if messages == nil || len(messages) < 1 {
		return errors.New("err: at least one message required in request")
	}
	for ind, msg := range messages {
		if msg.Role == "" {
			return fmt.Errorf("err: role is blank for message at %d index; role must be one of [%s]", ind, rolesStr)
		}
		if _, found := roles[msg.Role]; !found {
			return fmt.Errorf("err: invalid role %q for message at %d index; role must be one of [%s]", msg.Role, ind, rolesStr)
		}
		if msg.Content == "" {
			return fmt.Errorf("err: content is blank for message at %d index", ind)
		}
		if msg.Role == RoleTool && msg.ToolCallId == "" {
			return fmt.Errorf("err: tool_call_id is blank for message at %d index; tool_call_id is mandatory with role %q", ind, msg.Role)
		}
	}
	return nil
}

func validateModel(req *ChatCompletionsRequest) error {
	if req.Model == "" {
		return errors.New("err: model required in request")
	}
	if !(req.Model == modelChat || req.Model == modelReasoner) {
		return fmt.Errorf("err: invalid model %q; model should be one of [%s]", req.Model, modelStr)
	}
	return nil
}

func validateMultipleFields(req *ChatCompletionsRequest) error {
	if !(req.FrequencyPenalty >= -2 && req.FrequencyPenalty <= 2) {
		return fmt.Errorf("err: frequency_penalty is invalid; it should be number between -2 and 2")
	}

	if !(req.MaxTokens == 0 || (req.MaxTokens >= 1 && req.MaxTokens <= 8192)) {
		return fmt.Errorf("err: max_tokens is invalid; it should be number between 1 and 8192 or 0")
	}

	if !(req.PresencePenalty >= -2 && req.PresencePenalty <= 2) {
		return fmt.Errorf("err: presence_penalty is invalid; it should be number between -2 and 2")
	}

	if req.Temperature != nil {
		if !(*req.Temperature >= 0 && *req.Temperature <= 2) {
			return fmt.Errorf("err: temperature is invalid; it should be number between 0 and 2")
		}
	}

	if req.TopP != nil {
		if !(*req.TopP > 0 && *req.TopP <= 1) {
			return fmt.Errorf("err: invalid top_p value; the valid range of top_p is (0, 1.0]")
		}
	}

	return nil
}

func validateResponseFormat(req *ChatCompletionsRequest) error {
	if req.ResponseFormat != nil {
		if !(req.ResponseFormat.Type == ResponseFormatText || req.ResponseFormat.Type == ResponseFormatJsonObject) {
			return fmt.Errorf(`err: invalid response_format type %q; it should be one of [%s]`, req.ResponseFormat.Type, responseFormatStr)
		}
	}
	return nil
}

func validateStreamOptions(req *ChatCompletionsRequest) error {
	if req.StreamOptions != nil {
		if !req.Stream {
			return errors.New(`err: stream_options should be set along with stream = true`)
		}
	}
	return nil
}
