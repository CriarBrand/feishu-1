package feishu

import (
	"encoding/json"
	"net/http"
	"strings"
)

type InteractiveV1CardUpdateParam struct {
	Token string                           `json:"token"`
	Card  InteractiveV1CardUpdateCardParam `json:"card"`
}

type InteractiveV1CardUpdateCardParam struct {
	OpenIds  []string `json:"open_ids"`
	Elements interface{}
}

type InteractiveV1CardUpdateRes struct {
	ResponseCode
}

// InteractiveV1CardUpdate 消息卡片延迟更新
func (c *Client) InteractiveV1CardUpdate(param InteractiveV1CardUpdateParam) (*InteractiveV1CardUpdateRes, error) {
	paramByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPost, ServerUrl+"/open-apis/interactive/v1/card/update", strings.NewReader(string(paramByte)))
	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data InteractiveV1CardUpdateRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
