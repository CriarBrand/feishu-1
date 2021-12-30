package feishu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type UsersFindByDepartmentParam struct {
	UserIdType       string `json:"user_id_type"`
	DepartmentIdType string `json:"department_id_type"`
	DepartmentId     string `json:"department_id"`
	PageSize         int64  `json:"page_size"`
	PageToken        string `json:"page_token"`
}

type UsersFindByDepartmentRes struct {
	ResponseCode
	Data UsersFindByDepartmentResData `json:"data"`
}

type UsersFindByDepartmentResData struct {
	HasMore   bool                               `json:"has_more"`
	PageToken string                             `json:"page_token"`
	Items     []UsersFindByDepartmentResDataItem `json:"items"`
}

type UsersFindByDepartmentResDataItem struct {
	UnionId         string                                        `json:"union_id"`
	UserId          string                                        `json:"user_id"`
	OpenId          string                                        `json:"open_id"`
	Name            string                                        `json:"name"`
	EnName          string                                        `json:"en_name"`
	Nickname        string                                        `json:"nickname"`
	Email           string                                        `json:"email"`
	Mobile          string                                        `json:"mobile"`
	MobileVisible   bool                                          `json:"mobile_visible"`
	Gender          int64                                         `json:"gender"`
	AvatarKey       string                                        `json:"avatar_key"`
	Avatar          UsersFindByDepartmentResDataItemAvatar        `json:"avatar"`
	Status          UsersFindByDepartmentResDataItemStatus        `json:"status"`
	DepartmentIds   []string                                      `json:"department_ids"`
	LeaderUserId    string                                        `json:"leader_user_id"`
	City            string                                        `json:"city"`
	Country         string                                        `json:"country"`
	WorkStation     string                                        `json:"work_station"`
	JoinTime        int64                                         `json:"join_time"`
	IsTenantManager bool                                          `json:"is_tenant_manager"`
	EmployeeNo      string                                        `json:"employee_no"`
	EmployeeType    int64                                         `json:"employee_type"`
	Orders          []UsersFindByDepartmentResDataItemOrders      `json:"orders"`
	CustomAttrs     []UsersFindByDepartmentResDataItemCustomAttrs `json:"custom_attrs"`
	EnterpriseEmail string                                        `json:"enterprise_email"`
	JobTitle        string                                        `json:"job_title"`
	IsFrozen        bool                                          `json:"is_frozen"`
}

type UsersFindByDepartmentResDataItemAvatar struct {
	Avatar72     string `json:"avatar_72"`
	Avatar240    string `json:"avatar_240"`
	Avatar640    string `json:"avatar_640"`
	AvatarOrigin string `json:"avatar_origin"`
}

type UsersFindByDepartmentResDataItemStatus struct {
	IsFrozen    bool `json:"is_frozen"`
	IsResigned  bool `json:"is_resigned"`
	IsActivated bool `json:"is_activated"`
	IsExited    bool `json:"is_exited"`
	IsUnjoin    bool `json:"is_unjoin"`
}

type UsersFindByDepartmentResDataItemOrders struct {
	DepartmentId    string `json:"department_id"`
	UserOrder       int64  `json:"user_order"`
	DepartmentOrder int64  `json:"department_order"`
}

type UsersFindByDepartmentResDataItemCustomAttrs struct {
	Type  string                                           `json:"type"`
	Id    string                                           `json:"id"`
	Value UsersFindByDepartmentResDataItemCustomAttrsValue `json:"value"`
}

type UsersFindByDepartmentResDataItemCustomAttrsValue struct {
	Text        string                                                      `json:"text"`
	Url         string                                                      `json:"url"`
	PcUrl       string                                                      `json:"pc_url"`
	OptionId    string                                                      `json:"option_id"`
	OptionValue string                                                      `json:"option_value"`
	Name        string                                                      `json:"name"`
	PictureUrl  string                                                      `json:"picture_url"`
	GenericUser UsersFindByDepartmentResDataItemCustomAttrsValueGenericUser `json:"generic_user"`
}

type UsersFindByDepartmentResDataItemCustomAttrsValueGenericUser struct {
	Id   string `json:"id"`
	Type int64  `json:"type"`
}

// UsersFindByDepartment 获取部门直属用户列表
func (c *Client) UsersFindByDepartment(param UsersFindByDepartmentParam) (*UsersFindByDepartmentRes, error) {
	params := url.Values{}
	if param.UserIdType != "" {
		params.Add("user_id_type", param.UserIdType)
	}
	if param.DepartmentIdType != "" {
		params.Add("department_id_type", param.DepartmentIdType)
	}
	params.Add("department_id", param.DepartmentId)
	if param.PageSize > 0 {
		params.Add("page_size", fmt.Sprintf("%v", param.PageSize))
	}
	if param.PageToken != "" {
		params.Add("page_token", param.PageToken)
	}
	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/contact/v3/users/find_by_department?"+params.Encode(), nil)
	AccessToken, err := c.TokenManager.GetAccessToken()
	fmt.Println(AccessToken)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data UsersFindByDepartmentRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
