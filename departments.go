package feishu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DepartmentsChildrenParam struct {
	DepartmentId     string `json:"department_id"`
	UserIdType       string `json:"user_id_type"`
	DepartmentIdType string `json:"department_id_type"`
	FetchChild       bool   `json:"fetch_child"`
	PageSize         int64  `json:"page_size"`
	PageToken        string `json:"page_token"`
}

type DepartmentsChildrenRes struct {
	ResponseCode
	Data DepartmentsChildrenResData `json:"data"`
}

type DepartmentsChildrenResData struct {
	HasMore   bool                             `json:"has_more"`
	PageToken string                           `json:"page_token"`
	Items     []DepartmentsChildrenResDataItem `json:"items"`
}

type DepartmentsChildrenResDataItem struct {
	Name               string                                 `json:"name"`
	I18NName           DepartmentsChildrenResDataItemI18NName `json:"i18n_name"`
	ParentDepartmentId string                                 `json:"parent_department_id"`
	DepartmentId       string                                 `json:"department_id"`
	OpenDepartmentId   string                                 `json:"open_department_id"`
	LeaderUserId       string                                 `json:"leader_user_id"`
	ChatId             string                                 `json:"chat_id"`
	Order              string                                 `json:"order"`
	UnitIds            []string                               `json:"unit_ids"`
	MemberCount        int64                                  `json:"member_count"`
	Status             DepartmentsChildrenResDataItemStatus   `json:"status"`
	CreateGroupChat    bool                                   `json:"create_group_chat"`
}

type DepartmentsChildrenResDataItemI18NName struct {
	ZhCn string `json:"zh_cn"`
	JaJp string `json:"ja_jp"`
	EnUs string `json:"en_us"`
}

type DepartmentsChildrenResDataItemStatus struct {
	IsDeleted bool `json:"is_deleted"`
}

// DepartmentsChildren 获取子部门列表
func (c *Client) DepartmentsChildren(param DepartmentsChildrenParam) (*DepartmentsChildrenRes, error) {
	params := url.Values{}
	if param.UserIdType != "" {
		params.Add("user_id_type", param.UserIdType)
	}
	if param.DepartmentIdType != "" {
		params.Add("department_id_type", param.DepartmentIdType)
	}
	if param.FetchChild == true {
		params.Add("fetch_child", fmt.Sprintf("%v", param.FetchChild))
	}
	if param.PageSize > 0 {
		params.Add("page_size", fmt.Sprintf("%v", param.PageSize))
	}
	if param.PageToken != "" {
		params.Add("page_token", param.PageToken)
	}
	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/contact/v3/departments/"+param.DepartmentId+"/children?"+params.Encode(), nil)
	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data DepartmentsChildrenRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
