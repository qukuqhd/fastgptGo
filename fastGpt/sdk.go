package fastgpt

import (
	"encoding/json"
	"errors"
	"net/http"

	httpclient "github.com/qukuqhd/fastgptGo/http_client"
)

type FastGptSdkClient struct {
	baseUrl string
	httpCli *httpclient.HttpClient //请求FastGpt的客户端
	apiKey  string
}

func NewFastGptSdkClient(baseUrl string, apiKey string) *FastGptSdkClient {
	cli := httpclient.NewHttpClient()
	// 拦截添加认证请求头
	cli.AddReqInterception(func(r *http.Request) error {
		r.Header.Set("Authorization", "Bearer "+apiKey)
		return nil
	})
	// 拦截返回非200的状态码
	cli.AddRespInterception(func(r *http.Response) error {
		if r.StatusCode != http.StatusOK {
			return errors.New(http.StatusText(r.StatusCode))
		}
		return nil
	})
	// 拦截code不为200的错误消息
	cli.AddRespInterception(func(r *http.Response) error {
		if r.StatusCode != http.StatusOK {
			return errors.New(http.StatusText(r.StatusCode))
		}
		return nil
	})
	return &FastGptSdkClient{
		baseUrl: baseUrl,
		httpCli: cli,
		apiKey:  apiKey,
	}
}

// CreateKnowledgeBase 创建知识库
func (s *FastGptSdkClient) CreateKnowledgeBase(ReqInfo *CreateKnowledgeBaseReq) (*CreateKnowledgeBaseResp, error) {
	var respInfo CreateKnowledgeBaseResp
	err := s.httpCli.SendObjParse(http.MethodPost, s.baseUrl+"/core/dataset/create", &ReqInfo, &respInfo)
	if err != nil {
		return nil, err
	}
	return &respInfo, nil
}

// GetKnowledgeList 获取知识库列表
func (s *FastGptSdkClient) GetKnowledgeList(req *ReqListKnowledgeBase) (*ListKnowLedgeResp, error) {
	resp, err := s.httpCli.SendParameter(http.MethodPost, s.baseUrl+"/core/dataset/list", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respInfo ListKnowLedgeResp
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}
	return &respInfo, nil
}

// GetDetailKnowledge 获取知识库详细信息
func (s *FastGptSdkClient) GetDetailKnowledge(req *ReqDetailKnowledgeBase) (*DetailKnowledge, error) {
	resp, err := s.httpCli.SendParameter(http.MethodGet, s.baseUrl+"/core/dataset/detail", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respInfo DetailKnowledge
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}
	return &respInfo, nil
}

// UploadLocalFileSet 上传文件集合
func (s *FastGptSdkClient) UploadLocalFileSet(req *UploadFileSetReq) (*UploadFileSetResp, error) {
	resp, err := s.httpCli.SendForm(http.MethodPost, s.baseUrl+"/core/dataset/collection/create/localFile", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respInfo UploadFileSetResp
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return nil, err
	}
	return &respInfo, nil
}

// UploadLinkFileSet 添加链接集合
func (s *FastGptSdkClient) UploadLinkFileSet(req *UpLoadLinkFileReq) (*UploadLinkFileResp, error) {
	var resp UploadLinkFileResp
	err := s.httpCli.SendObjParse(http.MethodPost, "/core/dataset/collection/create/link", req, &resp)
	return &resp, err
}

// UploadTextSet 添加文本集合
func (s *FastGptSdkClient) UploadTextSet(req *UploadTextSetReq) (*UploadTextSetResp, error) {
	var resp UploadTextSetResp
	err := s.httpCli.SendObjParse(http.MethodPost, "/core/dataset/collection/create/text", req, &resp)
	return &resp, err
}

// GetAppChatHistory 获取应用程序的对话列表
func (s *FastGptSdkClient) GetAppChatHistory(req *GetAppChatHistoryReq) (*GetAppChatHistoryResp, error) {
	var resp GetAppChatHistoryResp
	err := s.httpCli.SendObjParse(http.MethodGet, "/core/chat/getHistories", req, &resp)
	return &resp, err
}

// GetAppChatHistoryRecords 获取应用程序聊天对话的记录列表
func (s *FastGptSdkClient) GetAppChatHistoryRecords(req *GetChatRecordsReq) (*GetChatRecordsResp, error) {
	var resp GetChatRecordsResp
	err := s.httpCli.SendObjParse(http.MethodPost, "/core/chat/getPaginationRecords", req, &resp)
	return &resp, err
}

//// NoStreamChat 非流式聊天
//func (s *FastGptSdkClient) NoStreamChat(req *ChatReq) (*NoStreamResp, error) {
//	var resp NoStreamResp
//	err := s.httpCli.SendObjParse(http.MethodPost, "/v1/chat/completions", req, &resp)
//	return &resp, err
//}
//
//// StreamChat 流式聊天
//func (s *FastGptSdkClient) StreamChat(req *ChatReq) (func(io.Writer) bool, error) {
//	resp, err := s.httpCli.SendObj(http.MethodPost, "/v1/chat/completions", req)
//	if err != nil {
//		return nil, err
//	}
//	return func(writer io.Writer) bool {
//		r := resp.Body
//		json.NewDecoder(r)
//		m := &StreamResp{}
//		if err := json.NewDecoder(r).Decode(m); err != nil {
//			r.Close()
//			return false
//		}
//		info, err := json.Marshal(resp)
//		if err != nil {
//			r.Close()
//			return false
//		}
//		_, err = writer.Write(info)
//		if err != nil {
//			r.Close()
//			return false
//		}
//		return true
//	}, nil
//}
