package feishu

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// SendMessagesParam 发送消息的请求结构体
type SendMessagesParam struct {
	ReceiveIdType string `json:"-"`
	ReceiveId     string `json:"receive_id"`
	Content       string `json:"content"`
	MsgType       string `json:"msg_type"`
}

// SendMessagesRes 发送消息的响应结构体
type SendMessagesRes struct {
	ResponseCode
	Data SendMessagesResData `json:"data"`
}

type SendMessagesResData struct {
	MessageId      string                        `json:"message_id"`
	RootId         string                        `json:"root_id"`
	ParentId       string                        `json:"parent_id"`
	MsgType        string                        `json:"msg_type"`
	CreateTime     string                        `json:"create_time"`
	UpdateTime     string                        `json:"update_time"`
	Deleted        bool                          `json:"deleted"`
	Updated        bool                          `json:"updated"`
	ChatId         string                        `json:"chat_id"`
	Sender         SendMessagesResDataSender     `json:"sender"`
	Body           SendMessagesResDataContent    `json:"body"`
	Mentions       []SendMessagesResDataMentions `json:"mentions"`
	UpperMessageId string                        `json:"upper_message_id"`
}

type SendMessagesResDataSender struct {
	Id         string `json:"id"`
	IdType     string `json:"id_type"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}

type SendMessagesResDataContent struct {
	Content string `json:"content"`
}

type SendMessagesResDataMentions struct {
	Key       string `json:"key"`
	Id        string `json:"id"`
	IdType    string `json:"id_type"`
	Name      string `json:"name"`
	TenantKey string `json:"tenant_key"`
}

// SendMessages 发送消息
func (c *Client) SendMessages(param SendMessagesParam) (*SendMessagesRes, error) {
	params := url.Values{}
	params.Add("receive_id_type", param.ReceiveIdType)
	jsonStr, _ := json.Marshal(param)
	request, _ := http.NewRequest(http.MethodPost, ServerUrl+"/open-apis/im/v1/messages?"+params.Encode(), strings.NewReader(string(jsonStr)))
	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data SendMessagesRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
