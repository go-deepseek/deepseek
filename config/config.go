package config

// Config for deepseek client.
//
//	ApiKey - deepseek API key.
//	TimeoutSeconds - http client timeout used by deepseek client.
//	DisableRequestValidation - disable request validation by deepseek client.
//	BaseURL - base URL for the deepseek API.
type Config struct {
	ApiKey                   string
	TimeoutSeconds           int
	DisableRequestValidation bool
	BaseURL                  string
}
