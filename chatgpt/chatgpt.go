package chatgpt

// * ChatCompletion Logic * //

const (
	endpoint = "https://api.openai.com/v1/chat/completions"

	ModelGPT4       ChatModel = "gpt-4"
	ModelGPT35Turbo ChatModel = "gpt-3.5-turbo"
	// TODO: Add more models

	ChatGPTRoleUser      ChatGPTRole = "user"
	ChatGPTRoleSystem    ChatGPTRole = "system"
	ChatGPTRoleAssistant ChatGPTRole = "assistant"
)

type ChatModel string
type ChatGPTRole string

// * ChatCompletion *//
type ChatCompletion struct {
	apikey   string
	rh       ChatGPTRequestHandler
	jt       JsonTranslatorHandler
	request  *ChatCompletionRequest
	response *ChatCompletionResponse
}

func NewChatCompletion(
	apikey string,
	model ChatModel,
	maxtokens int,
	requestHandler ChatGPTRequestHandler,
	jsonTranslator JsonTranslatorHandler) *ChatCompletion {

	return &ChatCompletion{
		apikey: apikey,
		rh:     requestHandler,
		jt:     jsonTranslator,
		request: &ChatCompletionRequest{
			Model:     model,
			MaxTokens: maxtokens,
		},
	}
}

func (c *ChatCompletion) AddMessage(role ChatGPTRole, content string) error {
	// TODO: Add validation of role and content
	c.request.Messages = append(c.request.Messages, Message{
		Role:    role,
		Content: content,
	})
	return nil
}

func (c *ChatCompletion) ClearMessages() {
	c.request.Messages = []Message{}
}

func (c *ChatCompletion) HandleRequest() (*ChatCompletionResponse, error) {

	request, err := c.jt.Marshal(c.request)
	if err != nil {
		return nil, err
	}
	response, err := c.rh(endpoint, c.apikey, request)
	if err != nil {
		return nil, err
	}

	var ccr ChatCompletionResponse
	if err := c.jt.Unmarshal(response, &ccr); err != nil {
		return nil, err
	}

	c.response = &ccr
	return c.response, nil
}
