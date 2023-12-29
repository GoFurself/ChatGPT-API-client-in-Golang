package chatgpt

// * Contains types and interfaces * //

// Function interface for ChatCompletion Request Handler
type ChatGPTRequestHandler func(url string, apikey string, requestModel []byte) ([]byte, error)

// Interface for Json Translator Handler
type JsonTranslatorHandler interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// Request structs
type ChatCompletionRequest struct {
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Model       ChatModel `json:"model"`
	Stop        []string  `json:"stop,omitempty"`
	TopP        float64   `json:"top_p,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}
type Message struct {
	Role    ChatGPTRole `json:"role"`
	Content string      `json:"content"`
}

// Response structs
type ChatCompletionResponse struct {
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint *string  `json:"system_fingerprint"`
}
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
