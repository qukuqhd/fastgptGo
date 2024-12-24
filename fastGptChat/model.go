package fastGptChat

type chatReq struct {
	ChatReqInfo
	Stream bool `json:"stream"`
	Detail bool `json:"detail"`
}

type ChatReqInfo struct {
	ChatId             string `json:"chatId"`
	ResponseChatItemId string `json:"responseChatItemId"`
	Variables          struct {
		Uid  string `json:"uid"`
		Name string `json:"name"`
	} `json:"variables"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}
type NoStreamResp struct {
	Id    string `json:"id"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}
type StreamResp struct {
	Data struct {
		Id      string `json:"id"`
		Object  string `json:"object"`
		Created int    `json:"created"`
		Choices []struct {
			Delta struct {
				Content string `json:"content"`
			} `json:"delta"`
			Index        int         `json:"index"`
			FinishReason interface{} `json:"finish_reason"`
		} `json:"choices"`
	} `json:"data"`
}
