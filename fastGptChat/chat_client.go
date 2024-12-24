package fastGptChat

import (
	"fmt"
	"io"
	"net/http"

	httpclient "github.com/qukuqhd/fastgptGo/http_client"
)

type ChatClient struct {
	cli         *httpclient.HttpClient
	appKey      string
	baseUrl     string
	ChatBufSize int64
}

func NewChatClient(appKey string, baseUrl string, ChatBufSize int64) *ChatClient {
	cli := httpclient.NewHttpClient()
	cli.AddReqInterception(func(request *http.Request) error {
		// 权限密钥添加
		request.Header.Set("Authorization", "Bearer "+appKey)
		return nil
	})
	cli.AddRespInterception(func(response *http.Response) error {
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("http status code: %d", response.StatusCode)
		}
		return nil
	})
	return &ChatClient{
		cli:         cli,
		appKey:      appKey,
		baseUrl:     baseUrl,
		ChatBufSize: ChatBufSize,
	}
}

// NoStreamChat 非流式响应
func (c *ChatClient) NoStreamChat(req *ChatReqInfo) (NoStreamResp, error) {
	var respInfo NoStreamResp
	var ReqChat = chatReq{
		ChatReqInfo: *req,
		Stream:      false,
		Detail:      false,
	}
	err := c.cli.SendObjParse(http.MethodPost, c.baseUrl+"v1/chat/completions", ReqChat, &respInfo)
	return respInfo, err
}

// StreamChat 流式响应
func (c *ChatClient) StreamChat(req *ChatReqInfo) (func(writer io.Writer) bool, error) {

	var ReqChat = chatReq{
		ChatReqInfo: *req,
		Stream:      true,
		Detail:      false,
	}
	resp, err := c.cli.SendObj(http.MethodPost, c.baseUrl+"v1/chat/completions", ReqChat)
	if err != nil {
		return nil, err
	}
	return func(writer io.Writer) bool {
		r := resp.Body
		buf := make([]byte, c.ChatBufSize)
		_, err := r.Read(buf)
		if err != nil {
			return false
		}
		_, err = io.Copy(writer, r)
		if err != nil {
			return false
		}
		return true
	}, nil
}
