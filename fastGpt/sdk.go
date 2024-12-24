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

// ListDoc 知识库的集合列表获取
func (s *FastGptSdkClient) ListDoc(req *ListDocReq) (*ListDocResp, error) {
	var resp ListDocResp
	err := s.httpCli.SendObjParse(http.MethodPost, "core/dataset/collection/list", req, &resp)
	return &resp, err
}

// GetDetailDoc 获取文档详细信息
func (s *FastGptSdkClient) GetDetailDoc(req *GetDetailDocReq) (*CommonResp, error) {
	resp, err := s.httpCli.SendParameter(http.MethodGet, s.baseUrl+"core/dataset/collection/detail", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respInfo CommonResp
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	return &respInfo, err
}

// DeleteDoc 删除文档
func (s *FastGptSdkClient) DeleteDoc(req *DeleteDocReq) (*DeleteDocResp, error) {
	resp, err := s.httpCli.SendParameter(http.MethodDelete, s.baseUrl+"core/dataset/collection/delete", req)
	defer resp.Body.Close()
	var respInfo DeleteDocResp
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	return &respInfo, err
}

// UpdateDoc 更新文档
func (s *FastGptSdkClient) UpdateDoc(req *UpdateDocReq) (*CommonResp, error) {
	var resp CommonResp
	err := s.httpCli.SendObjParse(http.MethodPut, s.baseUrl+"core/dataset/collection/update", req, &resp)
	return &resp, err
}

// GetPoints 获取文档知识点列表
func (s *FastGptSdkClient) GetPoints(req *GetPointsReq) (*CommonResp, error) {
	var resp CommonResp
	err := s.httpCli.SendObjParse(http.MethodPost, s.baseUrl+"core/dataset/data/v2/list", req, &resp)
	return &resp, err
}

// GetPointInfo 获取知识点信息
func (s *FastGptSdkClient) GetPointInfo(req *GetPointsReq) (*GetPointResp, error) {
	resp, err := s.httpCli.SendParameter(http.MethodGet, s.baseUrl+"core/dataset/data/detail", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respInfo GetPointResp
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	return &respInfo, err
}

// UpdatePoint 更新知识点
func (s *FastGptSdkClient) UpdatePoint(req *UpdatePointReq) (*UpdatePointResp, error) {
	var resp UpdatePointResp
	err := s.httpCli.SendObjParse(http.MethodPut, s.baseUrl+"core/dataset/data/update", req, &resp)

	return &resp, err
}

// DeletePoint 删除知识点
func (s *FastGptSdkClient) DeletePoint(req *DeletePointReq) (*CommonResp, error) {
	resp, err := s.httpCli.SendParameter(http.MethodDelete, s.baseUrl+"core/dataset/data/delete", req)
	if err != nil {
		return nil, err
	}
	var respInfo CommonResp
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	return &respInfo, err
}

// AddPoint 添加知识点
func (s *FastGptSdkClient) AddPoint(req *AddPointsReq) (*AddPointsResp, error) {
	var resp AddPointsResp
	err := s.httpCli.SendObjParse(http.MethodPost, s.baseUrl+"core/dataset/data/pushData", req, &resp)
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
