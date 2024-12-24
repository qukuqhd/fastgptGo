package fastgpt

import (
	"io"
	"time"
)

// Resp 响应数据的基本格式
type Resp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// CommonResp 通用的响应结构
type CommonResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

// ReqListKnowledgeBase 知识库列表的请求结构
type ReqListKnowledgeBase struct {
	ParentId int `parm:"parentId"`
}

type ReqDetailKnowledgeBase struct {
	Id string `parm:"id"`
}

// KnowledgeBase 知识库信息
type KnowledgeBase struct {
	ID                string      `json:"_id"`
	Avatar            string      `json:"avatar"`
	Name              string      `json:"name"`
	Intro             string      `json:"intro"`
	Type              string      `json:"type"`
	Permission        Permission  `json:"permission"`
	VectorModel       VectorModel `json:"vectorModel"`
	DefaultPermission int         `json:"defaultPermission"`
	InheritPermission bool        `json:"inheritPermission"`
	TmbId             string      `json:"tmbId"`
	UpdateTime        string      `json:"updateTime"`
}

type Permission struct {
	Value          int            `json:"value"`
	IsOwner        bool           `json:"isOwner"`
	PermissionList PermissionList `json:"_permissionList"`
	HasManagePer   bool           `json:"hasManagePer"`
	HasWritePer    bool           `json:"hasWritePer"`
	HasReadPer     bool           `json:"hasReadPer"`
}

type PermissionList struct {
	Read   PermissionItem `json:"read"`
	Write  PermissionItem `json:"write"`
	Manage PermissionItem `json:"manage"`
}

type PermissionItem struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Value        int    `json:"value"`
	CheckBoxType string `json:"checkBoxType"`
}

type VectorModel struct {
	Model            string                 `json:"model"`
	Name             string                 `json:"name"`
	Avatar           string                 `json:"avatar"`
	CharsPointsPrice int                    `json:"charsPointsPrice"`
	DefaultToken     int                    `json:"defaultToken"`
	MaxToken         int                    `json:"maxToken"`
	Weight           int                    `json:"weight"`
	DefaultConfig    map[string]interface{} `json:"defaultConfig"`
	DbConfig         map[string]interface{} `json:"dbConfig"`
	QueryConfig      map[string]interface{} `json:"queryConfig"`
}

type CreateTrainUsageReq struct {
	DataSetId string `json:"datasetId"` // 知识库id
	Name      string `json:"name"`      //
}

type CreateKnowledgeBaseReq struct {
	ParentId    interface{} `json:"parentId"`
	Type        string      `json:"type"`
	Name        string      `json:"name"`
	Intro       string      `json:"intro"`
	Avatar      string      `json:"avatar"`
	VectorModel string      `json:"vectorModel"`
	AgentModel  string      `json:"agentModel"`
}

type CreateKnowledgeBaseResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       string `json:"data"`
}

type CreateUniversalSetReq struct {
	DataSetId     string   `json:"datasetId"`     // 知识库id
	ParentId      string   `json:"parentId"`      // 父级ID，不填则默认为根目录
	TrainingType  string   `json:"trainingType"`  // 训练模式：chunk: 按文本长度进行分割; qa: QA拆分; auto: 增强训练
	ChunkSize     int      `json:"chunkSize"`     // 训练模式为chunk时，每个chunk的长度
	ChunkSplitter string   `json:"chunkSplitter"` // 训练模式为chunk时，分割符
	QaPrompt      string   `json:"qaPrompt"`      // 训练模式为qa时，问题的前缀
	Tags          []string `json:"tags"`          // 集合标签，字符串数组
	CreateTime    string   `json:"createTime"`    // 创建时间(Date/String)
}

type CreateUniversalSetResp struct {
	CollectionId string `json:"collectionId"` // 集合ID
	InsertLen    int    `json:"insertLen"`    // 插入的块数量
}

type CreateTextSetReq struct {
	Text          string      `json:"text"`          // 文本内容
	DataSetId     string      `json:"datasetId"`     // 知识库id
	ParentId      interface{} `json:"parentId"`      // 父级ID，不填则默认为根目录
	Name          string      `json:"name"`          // 训练集名称
	TrainingType  string      `json:"trainingType"`  // 训练模式：chunk: 按文本长度进行分割; qa: QA拆分; auto: 增强训练
	ChunkSize     int         `json:"chunkSize"`     // 训练模式为chunk时，每个chunk的长度
	ChunkSplitter string      `json:"chunkSplitter"` // 训练模式为chunk时，分割符
	QaPrompt      string      `json:"qaPrompt"`      // 训练模式为qa时，问题的前缀
	Metadata      any         `json:"metadata"`      // 元数据
}

type Results struct {
	InsertLen int      `json:"insertLen"` // 插入的块数量
	OverToken []string `json:"overToken"` // 超过token限制的文本
	Repeat    []string `json:"repeat"`    // 重复的文本
	Error     []string `json:"error"`     // 错误的文本

}

type CreateSetData struct {
	CollectionId string  `json:"collectionId"` // 集合ID
	Results      Results `json:"results"`
}
type CreateTextSetResp struct {
	Code       int           `json:"code"`       // 状态码
	StatusText string        `json:"statusText"` // 状态信息
	Message    string        `json:"message"`    // 消息
	Data       CreateSetData `json:"data"`       // 数据
}

type CreateLinkSetReq struct {
	Link          string      `json:"link"`          // 链接地址
	DataSetId     string      `json:"datasetId"`     // 知识库id
	ParentId      interface{} `json:"parentId"`      // 父级ID，不填则默认为根目录
	Name          string      `json:"name"`          // 训练集名称
	TrainingType  string      `json:"trainingType"`  // 训练模式：chunk: 按文本长度进行分割; qa: QA拆分; auto: 增强训练
	ChunkSize     int         `json:"chunkSize"`     // 训练模式为chunk时，每个chunk的长度
	ChunkSplitter string      `json:"chunkSplitter"` // 训练模式为chunk时，分割符
	QaPrompt      string      `json:"qaPrompt"`      // 训练模式为qa时，问题的前缀
	Metadata      any         `json:"metadata"`      // 元数据
}

type CreateLinkSetResp struct {
	CollectionId string `json:"collectionId"` // 集合ID
	Code         int    `json:"code"`         // 状态码
	StatusText   string `json:"statusText"`   // 状态信息
	Message      string `json:"message"`      // 消息
	Data         struct {
		CollectionId string `json:"collectionId"` // 集合ID
	} `json:"data"` // 数据
}

type CreateLocalFileSetReq struct {
	File          string      `json:"file"`          // 文件地址; 目前支持：pdf, docx, md, txt, html, csv
	DataSetId     string      `json:"datasetId"`     // 知识库id
	ParentId      interface{} `json:"parentId"`      // 父级ID，不填则默认为根目录
	TrainingType  string      `json:"trainingType"`  // 训练模式：chunk: 按文本长度进行分割; qa: QA拆分; auto: 增强训练
	ChunkSize     int         `json:"chunkSize"`     // 训练模式为chunk时，每个chunk的长度
	ChunkSplitter string      `json:"chunkSplitter"` // 训练模式为chunk时，分割符
	QaPrompt      string      `json:"qaPrompt"`      // 训练模式为qa时，问题的前缀
	Metadata      any         `json:"metadata"`      // 元数据
}

type CreateLocalFileSetResp struct {
	Code       int           `json:"code"`       // 状态码
	StatusText string        `json:"statusText"` // 状态信息
	Message    string        `json:"message"`    // 消息
	Data       CreateSetData `json:"data"`       // 数据
}

type GetCollectionListReq struct {
	PageNum    int    `json:"pageNum"`    // 页码
	PageSize   int    `json:"pageSize"`   // 每页数量
	DataSetId  string `json:"datasetId"`  // 知识库id
	ParentId   string `json:"parentId"`   // 父级ID
	SearchText string `json:"searchText"` // 搜索文本
}

type GetCollectionListResp struct {
	Code       int    `json:"code"`       // 状态码
	StatusText string `json:"statusText"` // 状态信息
	Message    string `json:"message"`    // 消息

}

type DetailKnowledge struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		Id          string      `json:"_id"`
		ParentId    interface{} `json:"parentId"`
		TeamId      string      `json:"teamId"`
		TmbId       string      `json:"tmbId"`
		Type        string      `json:"type"`
		Status      string      `json:"status"`
		Avatar      string      `json:"avatar"`
		Name        string      `json:"name"`
		VectorModel struct {
			Model            string `json:"model"`
			Name             string `json:"name"`
			CharsPointsPrice int    `json:"charsPointsPrice"`
			DefaultToken     int    `json:"defaultToken"`
			MaxToken         int    `json:"maxToken"`
			Weight           int    `json:"weight"`
		} `json:"vectorModel"`
		AgentModel struct {
			Model            string `json:"model"`
			Name             string `json:"name"`
			MaxContext       int    `json:"maxContext"`
			MaxResponse      int    `json:"maxResponse"`
			CharsPointsPrice int    `json:"charsPointsPrice"`
		} `json:"agentModel"`
		Intro      string    `json:"intro"`
		Permission string    `json:"permission"`
		UpdateTime time.Time `json:"updateTime"`
		CanWrite   bool      `json:"canWrite"`
		IsOwner    bool      `json:"isOwner"`
	} `json:"data"`
}

// FileInfo 定义文件应该拥有的格式
type FileInfo struct {
	Name    string    `form_file:"name"`
	Content io.Reader `form_file:"content"`
}

type UploadFileSetReq struct {
	File FileInfo `form:"type:file;name:file"`
	Data struct {
		DatasetId     string                 `json:"datasetId"`
		ParentId      interface{}            `json:"parentId"`
		TrainingType  string                 `json:"trainingType"` // 训练方式 qa 或者是 chunk
		ChunkSize     int                    `json:"chunkSize"`
		ChunkSplitter string                 `json:"chunkSplitter"`
		QaPrompt      string                 `json:"qaPrompt"`
		Metadata      map[string]interface{} `json:"metadata"`
	} `form:"type:str;name:data"`
}

type UploadFileSetResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		CollectionId string `json:"collectionId"`
		Results      struct {
			InsertLen int           `json:"insertLen"`
			OverToken []interface{} `json:"overToken"`
			Repeat    []interface{} `json:"repeat"`
			Error     []interface{} `json:"error"`
		} `json:"results"`
	} `json:"data"`
}

type UpLoadLinkFileReq struct {
	Link          string      `json:"link"`
	DatasetId     string      `json:"datasetId"`
	ParentId      interface{} `json:"parentId"`
	TrainingType  string      `json:"trainingType"`
	ChunkSize     int         `json:"chunkSize"`
	ChunkSplitter string      `json:"chunkSplitter"`
	QaPrompt      string      `json:"qaPrompt"`
	Metadata      struct {
		WebPageSelector string `json:"webPageSelector"`
	} `json:"metadata"`
}

type UploadLinkFileResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		CollectionId string `json:"collectionId"`
		Results      struct {
			InsertLen int           `json:"insertLen"`
			OverToken []interface{} `json:"overToken"`
			Repeat    []interface{} `json:"repeat"`
			Error     []interface{} `json:"error"`
		} `json:"results"`
	} `json:"data"`
}

type UploadTextSetReq struct {
	Text          string      `json:"text"`
	DatasetId     string      `json:"datasetId"`
	ParentId      interface{} `json:"parentId"`
	Name          string      `json:"name"`
	TrainingType  string      `json:"trainingType"`
	ChunkSize     int         `json:"chunkSize"`
	ChunkSplitter string      `json:"chunkSplitter"`
	QaPrompt      string      `json:"qaPrompt"`
	Metadata      struct {
	} `json:"metadata"`
}

type UploadTextSetResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		CollectionId string `json:"collectionId"`
		Results      struct {
			InsertLen int           `json:"insertLen"`
			OverToken []interface{} `json:"overToken"`
			Repeat    []interface{} `json:"repeat"`
			Error     []interface{} `json:"error"`
		} `json:"results"`
	} `json:"data"`
}

type GetAppChatHistoryReq struct {
	AppId    string `json:"appId"`
	Offset   int    `json:"offset"`
	PageSize int    `json:"pageSize"`
	Source   string `json:"source"`
}

type GetAppChatHistoryResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		List []struct {
			ChatId      string    `json:"chatId"`
			UpdateTime  time.Time `json:"updateTime"`
			AppId       string    `json:"appId"`
			CustomTitle string    `json:"customTitle"`
			Title       string    `json:"title"`
			Top         bool      `json:"top"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}

type GetChatRecordsReq struct {
	AppId               string `json:"appId"`
	ChatId              string `json:"chatId"`
	Offset              int    `json:"offset"`
	PageSize            int    `json:"pageSize"`
	LoadCustomFeedbacks bool   `json:"loadCustomFeedbacks"`
}

type GetChatRecordsResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		List []struct {
			Id     string `json:"_id"`
			DataId string `json:"dataId"`
			Obj    string `json:"obj"`
			Value  []struct {
				Type string `json:"type"`
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
			} `json:"value"`
			CustomFeedbacks      []interface{} `json:"customFeedbacks"`
			LlmModuleAccount     int           `json:"llmModuleAccount,omitempty"`
			TotalQuoteList       []interface{} `json:"totalQuoteList,omitempty"`
			TotalRunningTime     float64       `json:"totalRunningTime,omitempty"`
			HistoryPreviewLength int           `json:"historyPreviewLength,omitempty"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}

type ChatReq struct {
	ChatId             string `json:"chatId"`
	Stream             bool   `json:"stream"`
	Detail             bool   `json:"detail"`
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

type ListKnowLedgeResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       []struct {
		Id                string `json:"_id"`
		Avatar            string `json:"avatar"`
		DefaultPermission int    `json:"defaultPermission"`
		InheritPermission bool   `json:"inheritPermission"`
		Intro             string `json:"intro"`
		Name              string `json:"name"`
		Permission        struct {
			PermissionList struct {
				Manage struct {
					CheckBoxType string `json:"checkBoxType"`
					Description  string `json:"description"`
					Name         string `json:"name"`
					Value        int    `json:"value"`
				} `json:"manage"`
				Read struct {
					CheckBoxType string `json:"checkBoxType"`
					Description  string `json:"description"`
					Name         string `json:"name"`
					Value        int    `json:"value"`
				} `json:"read"`
				Write struct {
					CheckBoxType string `json:"checkBoxType"`
					Description  string `json:"description"`
					Name         string `json:"name"`
					Value        int    `json:"value"`
				} `json:"write"`
			} `json:"_permissionList"`
			HasManagePer bool  `json:"hasManagePer"`
			HasReadPer   bool  `json:"hasReadPer"`
			HasWritePer  bool  `json:"hasWritePer"`
			IsOwner      bool  `json:"isOwner"`
			Value        int64 `json:"value"`
		} `json:"permission"`
		TmbId       string    `json:"tmbId"`
		Type        string    `json:"type"`
		UpdateTime  time.Time `json:"updateTime"`
		VectorModel struct {
			Avatar           string `json:"avatar"`
			CharsPointsPrice int    `json:"charsPointsPrice"`
			DbConfig         struct {
			} `json:"dbConfig"`
			DefaultConfig struct {
			} `json:"defaultConfig"`
			DefaultToken int    `json:"defaultToken"`
			MaxToken     int    `json:"maxToken"`
			Model        string `json:"model"`
			Name         string `json:"name"`
			QueryConfig  struct {
			} `json:"queryConfig"`
			Weight int `json:"weight"`
		} `json:"vectorModel"`
	} `json:"data"`
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

type ListDocReq struct {
	PageNum    int         `json:"pageNum"`
	PageSize   int         `json:"pageSize"`
	DatasetId  string      `json:"datasetId"`
	ParentId   interface{} `json:"parentId"`
	SearchText string      `json:"searchText"`
}

type ListDocResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		PageNum  int `json:"pageNum"`
		PageSize int `json:"pageSize"`
		Data     []struct {
			Id             string      `json:"_id"`
			ParentId       interface{} `json:"parentId"`
			TmbId          string      `json:"tmbId"`
			Type           string      `json:"type"`
			Name           string      `json:"name"`
			UpdateTime     time.Time   `json:"updateTime"`
			DataAmount     int         `json:"dataAmount"`
			TrainingAmount int         `json:"trainingAmount"`
			ExternalFileId string      `json:"externalFileId"`
			Tags           []string    `json:"tags"`
			Forbid         bool        `json:"forbid"`
			TrainingType   string      `json:"trainingType"`
			Permission     struct {
				Value        int64 `json:"value"`
				IsOwner      bool  `json:"isOwner"`
				HasManagePer bool  `json:"hasManagePer"`
				HasWritePer  bool  `json:"hasWritePer"`
				HasReadPer   bool  `json:"hasReadPer"`
			} `json:"permission"`
			RawLink string `json:"rawLink,omitempty"`
		} `json:"data"`
		Total int `json:"total"`
	} `json:"data"`
}

type GetDetailDocReq struct {
	Id string `parm:"id"`
}

type DeleteDocReq struct {
	Id string `parm:"id"`
}

type DeleteDocResp struct {
	Code       int         `json:"code"`
	StatusText string      `json:"statusText"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type UpdateDocReq struct {
	Id         string      `json:"id"`
	ParentId   interface{} `json:"parentId"`
	Name       string      `json:"name"`
	Tags       []string    `json:"tags"`
	Forbid     bool        `json:"forbid"`
	CreateTime time.Time   `json:"createTime"`
}

type GetPointsReq struct {
	Offset       int    `json:"offset"`
	PageSize     int    `json:"pageSize"`
	CollectionId string `json:"collectionId"`
	SearchText   string `json:"searchText"`
}

type GetPointReq struct {
	Id string `parm:"id"`
}
type GetPointResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
	Data       struct {
		Id         string `json:"id"`
		Q          string `json:"q"`
		A          string `json:"a"`
		ChunkIndex int    `json:"chunkIndex"`
		Indexes    []struct {
			DefaultIndex bool   `json:"defaultIndex"`
			Type         string `json:"type"`
			DataId       string `json:"dataId"`
			Text         string `json:"text"`
			Id           string `json:"_id"`
		} `json:"indexes"`
		DatasetId    string `json:"datasetId"`
		CollectionId string `json:"collectionId"`
		SourceName   string `json:"sourceName"`
		SourceId     string `json:"sourceId"`
		IsOwner      bool   `json:"isOwner"`
		CanWrite     bool   `json:"canWrite"`
	} `json:"data"`
}

type UpdatePointReq struct {
	DataId  string `json:"dataId"`
	Q       string `json:"q"`
	A       string `json:"a"`
	Indexes []struct {
		DataId       string `json:"dataId,omitempty"`
		DefaultIndex bool   `json:"defaultIndex,omitempty"`
		Text         string `json:"text"`
	} `json:"indexes"`
}

type UpdatePointResp struct {
	Code       int         `json:"code"`
	StatusText string      `json:"statusText"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
type DeletePointReq struct {
	Id string `parm:"id"`
}

type AddPointsReq struct {
	CollectionId string `json:"collectionId"`
	TrainingMode string `json:"trainingMode"`
	Prompt       string `json:"prompt"`
	BillId       string `json:"billId"`
	Data         []struct {
		Q       string `json:"q"`
		A       string `json:"a"`
		Indexes []struct {
			Text string `json:"text"`
		} `json:"indexes,omitempty"`
	} `json:"data"`
}

type AddPointsResp struct {
	Code       int    `json:"code"`
	StatusText string `json:"statusText"`
	Data       struct {
		InsertLen int           `json:"insertLen"`
		OverToken []interface{} `json:"overToken"`
		Repeat    []interface{} `json:"repeat"`
		Error     []interface{} `json:"error"`
	} `json:"data"`
}
