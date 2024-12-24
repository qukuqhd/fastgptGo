package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
)

type ReqInterception func(*http.Request) error
type RespInterception func(*http.Response) error

type HttpClient struct {
	netCLi            *http.Client       // 网络请求客户端
	reqInterceptions  []ReqInterception  // 请求拦截
	respInterceptions []RespInterception // 响应拦截
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		netCLi:            http.DefaultClient,
		reqInterceptions:  []ReqInterception{},
		respInterceptions: []RespInterception{},
	}
}

// AddReqInterception 添加请求拦截器
func (h *HttpClient) AddReqInterception(reqInterceptions ...ReqInterception) {
	h.reqInterceptions = append(h.reqInterceptions, reqInterceptions...)
}

// AddRespInterception 添加响应拦截器
func (h *HttpClient) AddRespInterception(respInterceptions ...RespInterception) {
	h.respInterceptions = append(h.respInterceptions, respInterceptions...)
}

func (h *HttpClient) sendHttp(req *http.Request) (*http.Response, error) {
	for _, reqInterception := range h.reqInterceptions {
		err := reqInterception(req)
		if err != nil {
			return nil, err
		}
	}
	resp, err := h.netCLi.Do(req)

	if err != nil {
		return nil, err
	}
	for _, respInterception := range h.respInterceptions {
		err = respInterception(resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// SendObj 转换为json格式并发送请求
func (h *HttpClient) SendObj(method string, url string, reqObj any) (*http.Response, error) {
	info, err := json.Marshal(reqObj)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(info))
	if err != nil {
		return nil, err
	}
	// 设置请求格式为json
	req.Header.Set("Content-Type", "application/json")
	return h.sendHttp(req)
}

// SendParameter 发送请求参数
func (h *HttpClient) SendParameter(method string, urlPath string, obj any) (*http.Response, error) {
	req, err := http.NewRequest(method, urlPath, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	parameter, err := GetParam(obj)
	if err != nil {
		return nil, err
	}
	for key, value := range parameter {
		// 使用反射来获取类型
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Map | reflect.Slice:
			{
				data, err := json.Marshal(value)
				if err != nil {
					continue //错误就忽略
				}
				q.Set(key, string(data))
			}
		default:
			{
				q.Set(key, fmt.Sprintf("%v", value))
			}
		}
	}
	req.URL.RawQuery = q.Encode()
	return h.sendHttp(req)
}

type HttpFormOpt struct {
	Files []struct {
		FileName    string
		FieldName   string
		FileContent io.Reader
	}
	Fields []struct {
		FieldName    string
		FieldContent string
	}
}

// SendForm 发送表单请求
func (h *HttpClient) SendForm(method string, url string, obj any) (*http.Response, error) {
	opt, err := GetForm(obj)
	if err != nil {
		return nil, err
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, fileField := range opt.Files {
		part, err := writer.CreateFormFile(fileField.FieldName, fileField.FileName)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, fileField.FileContent)
		if err != nil {
			return nil, err
		}
	}
	for _, field := range opt.Fields {
		err := writer.WriteField(field.FieldName, field.FieldContent)
		if err != nil {
			return nil, err
		}
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return nil, err
	}
	return h.sendHttp(req)
}

func (h *HttpClient) Send(method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return h.sendHttp(req)
}

// SendObjParse 发送请求并解析响应对象
func (h *HttpClient) SendObjParse(method string, url string, reqObj any, respObj any) error {
	resp, err := h.SendObj(method, url, reqObj)
	if err != nil {
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(respObj)
	return err
}
