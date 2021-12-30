/*
   @Time : 2021/12/29 14:16
   @Author : 铁甲
   @File : calendar
   @Software: GoLand
*/

package feishu

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// -----------------------------------------------创建日历访问控制---------------------------------------------------------

// CreateACLForCalendarReq 创建日历访问控制-请求体对外结构
type CreateACLForCalendarReq struct {
	CalendarId string                        `json:"calendar_id"` // 日历ID
	Params     CreateACLForCalendarReqParams `json:"params"`      // 查询参数
	Body       CreateACLForCalendarReqBody   `json:"body"`        // 请求体
}

// CreateACLForCalendarReqParams 创建日历访问控制-请求参数
type CreateACLForCalendarReqParams struct {
	UserIdType string `json:"user_id_type"` // 用户 ID 类型可选值有： open_id：用户的 open id union_id：用户的 union id user_id：用户的 user id
}

// CreateACLForCalendarReqBody 创建日历访问控制-请求体
type CreateACLForCalendarReqBody struct {
	Role  string                       `json:"role"`  // 对日历的访问权限 可选值有： unknown：未知权限 free_busy_reader：游客，只能看到忙碌/空闲信息 reader：订阅者，查看所有日程详情 writer：编辑者，创建及修改日程 owner：管理员，管理日历及共享设置
	Scope CreateACLForCalendarReqScope `json:"scope"` // 权限范围
}

// CreateACLForCalendarReqScope 创建日历访问控制-请求体Scope
type CreateACLForCalendarReqScope struct {
	Type   string `json:"type"`    // 权限类型，当type为User时，值为open_id/user_id/union_id 可选值有： user：用户
	UserId string `json:"user_id"` // 用户ID
}

// CreateACLForCalendarRes 创建日历访问控制-响应体
type CreateACLForCalendarRes struct {
	ResponseCode
	Data CreateACLForCalendarResData `json:"data"`
}

// CreateACLForCalendarResData 创建日历访问控制-响应体data
type CreateACLForCalendarResData struct {
	AclId string `json:"acl_id"`
	CreateACLForCalendarReqBody
}

// CreateACLForCalendar 创建日历访问控制
func (c *Client) CreateACLForCalendar(req CreateACLForCalendarReq) (*CreateACLForCalendarRes, error) {
	bodyByte, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	if req.Params.UserIdType != "" {
		params.Add("user_id_type", req.Params.UserIdType)
	}
	request, _ := http.NewRequest(http.MethodPost, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/acls?"+
		params.Encode(), strings.NewReader(string(bodyByte)))

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data CreateACLForCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------删除日历访问控制------------------------------------------------------

// DeleteACLForCalendarReq 删除日历访问控制-请求体对外结构
type DeleteACLForCalendarReq struct {
	CalendarId string `json:"calendar_id"` // 日历ID
	AclId      string `json:"acl_id"`      // acl资源ID
}

// DeleteACLForCalendarRes 删除日历访问控制-响应体
type DeleteACLForCalendarRes struct {
	ResponseCode
	Data interface{} `json:"data"`
}

// DeleteACLForCalendar 删除日历访问控制
func (c *Client) DeleteACLForCalendar(req DeleteACLForCalendarReq) (*DeleteACLForCalendarRes, error) {
	request, _ := http.NewRequest(http.MethodDelete, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/acls/"+
		req.AclId, nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data DeleteACLForCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------获取日历访问控制列表---------------------------------------------------

// GetACLForCalendarReq 获取日历访问控制列表-请求体对外结构
type GetACLForCalendarReq struct {
	CalendarId string                     `json:"calendar_id"` // 日历ID
	Params     GetACLForCalendarReqParams `json:"params"`      // 查询参数
}

// GetACLForCalendarReqParams 获取日历访问控制列表-查询参数
type GetACLForCalendarReqParams struct {
	UserIdType string `json:"user_id_type"` // 用户 ID 类型 可选值有： open_id：用户的 open id union_id：用户的 union id user_id：用户的 user id
	PageToken  string `json:"page_token"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
	PageSize   int64  `json:"page_size"`    // 分页大小，最小值10，最大值50
}

// GetACLForCalendarRes 获取日历访问控制列表-响应体
type GetACLForCalendarRes struct {
	ResponseCode
	Data GetACLForCalendarResData `json:"data"`
}

// GetACLForCalendarResData  获取日历访问控制列表-响应体Data
type GetACLForCalendarResData struct {
	Acls      []GetACLForCalendarResDataAcls `json:"acls"`       // 入参日历对应的acl列表
	HasMore   bool                           `json:"has_more"`   // 是否还有更多项
	PageToken string                         `json:"page_token"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
}

// GetACLForCalendarResDataAcls  获取日历访问控制列表-响应体Data的Acl部分
type GetACLForCalendarResDataAcls struct {
	AclId string                            `json:"acl_id"` // acl资源ID
	Role  string                            `json:"role"`   // 对日历的访问权限，可选值有： unknown：未知权限 free_busy_reader：游客，只能看到忙碌/空闲信息 reader：订阅者，查看所有日程详情 writer：编辑者，创建及修改日程 owner：管理员，管理日历及共享设置
	Scope GetACLForCalendarResDataAclsScope `json:"scope"`  // 权限范围
}

// GetACLForCalendarResDataAclsScope 获取日历访问控制列表-响应体Data的Scope
type GetACLForCalendarResDataAclsScope struct {
	Type   string `json:"type"`    // 权限类型，当type为User时，值为open_id/user_id/union_id 可选值有：user：用户
	UserId string `json:"user_id"` // 用户ID
}

// GetACLForCalendar 获取日历访问控制列表
func (c *Client) GetACLForCalendar(req GetACLForCalendarReq) (*GetACLForCalendarRes, error) {
	params := url.Values{}
	if req.Params.UserIdType != "" {
		params.Add("user_id_type", req.Params.UserIdType)
	}
	if req.Params.PageToken != "" {
		params.Add("page_token", req.Params.PageToken)
	}
	if req.Params.PageSize > 0 {
		params.Add("page_size", fmt.Sprintf("%v", req.Params.PageSize))
	}

	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/calendar/v4/calendars/"+
		req.CalendarId+"/acls?"+params.Encode(), nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data GetACLForCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------创建日历------------------------------------------------------------

// CreateCalendarReq 创建日历-请求体对外结构
type CreateCalendarReq struct {
	Body CreateCalendarReqBody `json:"body"`
}

// CreateCalendarReqBody 创建日历-请求体
type CreateCalendarReqBody struct {
	Summary      string `json:"summary,omitempty"`       // 日历标题 数据校验规则： 最大长度：255 字符
	Description  string `json:"description,omitempty"`   // 日历描述 数据校验规则： 最大长度：255 字符
	Permissions  string `json:"permissions,omitempty"`   // 日历公开范围 可选值有：private：私密 show_only_free_busy：仅展示忙闲信息 public：他人可查看日程详情
	Color        int64  `json:"color,omitempty"`         // 日历颜色，颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效
	SummaryAlias string `json:"summary_alias,omitempty"` // 日历备注名，修改或添加后仅对当前身份生效 数据校验规则： 最大长度：255 字符
}

// CreateCalendarRes 创建日历-响应体
type CreateCalendarRes struct {
	ResponseCode
	Data CreateCalendarResData `json:"data"`
}

// CreateCalendarResData 创建日历-响应体data
type CreateCalendarResData struct {
	Calendar CreateCalendarResCalendar `json:"calendar"`
}

// CreateCalendarResCalendar 创建日历-响应体data的Calendar结构部分
type CreateCalendarResCalendar struct {
	CalendarId   string `json:"calendar_id"`    // 日历ID
	Summary      string `json:"summary"`        // 日历标题
	Description  string `json:"description"`    // 日历描述
	Permissions  string `json:"permissions"`    // 日历公开范围 可选值有： private：私密 show_only_free_busy：仅展示忙闲信息 public：他人可查看日程详情
	Color        int64  `json:"color"`          // 日历颜色，颜色RGB值的int32表示。客户端展示时会映射到色板上最接近的一种颜色。仅对当前身份生效
	Type         string `json:"type"`           // 日历类型 可选值有： unknown：未知类型 primary：用户或应用的主日历 shared：由用户或应用创建的共享日历 google：用户绑定的谷歌日历 resource：会议室日历 exchange：用户绑定的Exchange日历
	SummaryAlias string `json:"summary_alias"`  // 日历备注名，修改或添加后仅对当前身份生效
	IsDeleted    bool   `json:"is_deleted"`     // 对于当前身份，日历是否已经被标记为删除
	IsThirdParty bool   `json:"is_third_party"` // 当前日历是否是第三方数据；三方日历及日程只支持读，不支持写入
	Role         string `json:"role"`           // 当前身份对于该日历的访问权限 可选值有： unknown：未知权限 free_busy_reader：游客，只能看到忙碌/空闲信息 reader：订阅者，查看所有日程详情 writer：编辑者，创建及修改日程 owner：管理员，管理日历及共享设置
}

// CreateCalendar 创建日历
func (c *Client) CreateCalendar(req CreateCalendarReq) (*CreateCalendarRes, error) {
	bodyByte, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPost, ServerUrl+"/open-apis/calendar/v4/calendars",
		strings.NewReader(string(bodyByte)))

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data CreateCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------删除日历------------------------------------------------------------

// DeleteCalendarReq 删除日历-请求体对外结构
type DeleteCalendarReq struct {
	CalendarId string `json:"calendar_id"` // 日历ID
}

// DeleteCalendarRes 删除日历-响应体
type DeleteCalendarRes struct {
	ResponseCode
	Data interface{} `json:"data"`
}

// DeleteCalendar 删除日历
func (c *Client) DeleteCalendar(req DeleteCalendarReq) (*DeleteCalendarRes, error) {
	request, _ := http.NewRequest(http.MethodDelete, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId, nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data DeleteCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------获取日历------------------------------------------------------------

// GetCalendarReq 获取日历-请求体对外结构
type GetCalendarReq struct {
	CalendarId string `json:"calendar_id"` // 日历ID
}

// GetCalendarRes 获取日历-响应体
type GetCalendarRes struct {
	ResponseCode
	Data GetCalendarResData `json:"data"`
}

// GetCalendarResData 获取日历-响应体的data
type GetCalendarResData CreateCalendarResCalendar

// GetCalendar 获取日历
func (c *Client) GetCalendar(req GetCalendarReq) (*GetCalendarRes, error) {
	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId, nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data GetCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------获取日历列表---------------------------------------------------------

// GetCalendarListReq 获取日历列表-请求体对外结构
type GetCalendarListReq struct {
	Params GetCalendarListReqParams `json:"params"` // 查询参数
}

// GetCalendarListReqParams 获取日历列表-查询参数
type GetCalendarListReqParams struct {
	PageToken string `json:"page_token"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果
	PageSize  int64  `json:"page_size"`  // 一次请求要求返回最大数量，默认500，取值范围为[50. 1000]
	SyncToken string `json:"sync_token"` // 上次请求Response返回的增量同步标记，分页请求未结束时为空
}

// GetCalendarListRes 获取日历列表-响应体
type GetCalendarListRes struct {
	ResponseCode
	Data GetCalendarListResData `json:"data"`
}

type GetCalendarListResData struct {
	HasMore      bool                             `json:"has_more"`      // 是否有下一页数据
	PageToken    string                           `json:"page_token"`    // 下次请求需要带上的分页标记，90 天有效期
	SyncToken    string                           `json:"sync_token"`    // 下次请求需要带上的增量同步标记，90 天有效期
	CalendarList []GetCalendarListResCalendarData `json:"calendar_list"` // 分页加载的日历数据列表
}

// GetCalendarListResCalendarData 获取日历列表-响应体的data的日历部分
type GetCalendarListResCalendarData CreateCalendarResCalendar

// GetCalendarList 获取日历列表
func (c *Client) GetCalendarList(req GetCalendarListReq) (*GetCalendarListRes, error) {
	params := url.Values{}
	if req.Params.SyncToken != "" {
		params.Add("sync_token", req.Params.SyncToken)
	}
	if req.Params.PageToken != "" {
		params.Add("page_token", req.Params.PageToken)
	}
	if req.Params.PageSize > 0 {
		params.Add("page_size", fmt.Sprintf("%v", req.Params.PageSize))
	}
	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/calendar/v4/calendars?"+params.Encode(), nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data GetCalendarListRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------更新日历------------------------------------------------------------

// UpdateCalendarReq 更新日历-请求体对外结构
type UpdateCalendarReq struct {
	CalendarId string                `json:"calendar_id"` // 日历ID
	Body       UpdateCalendarReqBody `json:"body"`        // 请求体
}

// UpdateCalendarReqBody 更新日历-请求体
type UpdateCalendarReqBody struct {
	Summary      string `json:"summary,omitempty"`
	Description  string `json:"description,omitempty"`
	Permissions  string `json:"permissions,omitempty"`
	Color        int64  `json:"color,omitempty"`
	SummaryAlias string `json:"summary_alias,omitempty"`
}

// UpdateCalendarRes 更新日历-响应体
type UpdateCalendarRes struct {
	ResponseCode
	Data UpdateCalendarResData `json:"data"`
}

// UpdateCalendarResData 更新日历-响应体Data
type UpdateCalendarResData struct {
	Calendar UpdateCalendarResCalendar `json:"calendar"`
}

// UpdateCalendarResCalendar 更新日历-响应体日历部分
type UpdateCalendarResCalendar CreateCalendarResCalendar

// UpdateCalendar 更新日历
func (c *Client) UpdateCalendar(req UpdateCalendarReq) (*UpdateCalendarRes, error) {
	bodyByte, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPatch, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId,
		strings.NewReader(string(bodyByte)))

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data UpdateCalendarRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------创建日程------------------------------------------------------------

// CreateCalendarEventsReq 创建日程-请求体对外结构
type CreateCalendarEventsReq struct {
	CalendarId string                      `json:"calendar_id"` // 日历ID
	Body       CreateCalendarEventsReqBody `json:"body"`        // 请求体
}

// CreateCalendarEventsReqBody 创建日程-请求体
type CreateCalendarEventsReqBody struct {
	Summary          string                    `json:"summary,omitempty"`           // 日程标题 数据校验规则： 最大长度：1000 字符
	Description      string                    `json:"description,omitempty"`       // 日程描述；目前不支持编辑富文本描述，如果日程描述通过客户端编辑过，更新描述会导致富文本格式丢失 数据校验规则： 最大长度：40960 字符
	NeedNotification bool                      `json:"need_notification,omitempty"` // 更新日程是否给日程参与人发送bot通知，默认为true
	StartTime        CalendarEventsTime        `json:"start_time"`                  // 日程开始时间
	EndTime          CalendarEventsTime        `json:"end_time"`                    // 日程结束时间
	Vchat            *CalendarEventsVchat      `json:"vchat,omitempty"`             // 视频会议信息
	Visibility       string                    `json:"visibility,omitempty"`        // 日程公开范围，新建日程默认为Default；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效 可选值有： default：默认权限，跟随日历权限，默认仅向他人显示是否“忙碌” public：公开，显示日程详情 private：私密，仅自己可见详情
	AttendeeAbility  string                    `json:"attendee_ability,omitempty"`  // 参与人权限 可选值有： none：无法编辑日程、无法邀请其它参与人、无法查看参与人列表 can_see_others：无法编辑日程、无法邀请其它参与人、可以查看参与人列表 can_invite_others：无法编辑日程、可以邀请其它参与人、可以查看参与人列表 can_modify_event：可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus   string                    `json:"free_busy_status,omitempty"`  // 日程占用的忙闲状态，新建日程默认为Busy；仅新建日程时对所有参与人生效，之后修改该属性仅对当前身份生效 可选值有： busy：忙碌 free：空闲
	Location         *CalendarEventsLocation   `json:"location,omitempty"`          // 日程地点
	Color            int64                     `json:"color,omitempty"`             // 日程颜色，颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders        []CalendarEventsReminders `json:"reminders,omitempty"`         // 日程提醒列表
	Recurrence       string                    `json:"recurrence,omitempty"`        // 重复日程的重复性规则；参考rfc5545(https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10) ； 不支持COUNT和UNTIL同时出现； 预定会议室重复日程长度不得超过两年。 示例值："FREQ=DAILY;INTERVAL=1" 数据校验规则： 最大长度：2000 字符
	Schemas          []CalendarEventsSchemas   `json:"schemas,omitempty"`           // 日程自定义信息；控制日程详情页的ui展示
}

// CalendarEventsTime 日程开始时间
type CalendarEventsTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段，如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳，如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称，使用IANA Time Zone Database标准，如Asia/Shanghai；全天日程时区固定为UTC，非全天日程时区默认为Asia/Shanghai
}

// CalendarEventsVchat 视频会议信息
type CalendarEventsVchat struct {
	VcType      string `json:"vc_type,omitempty"`     // 视频会议类型 vc：飞书视频会议，取该类型时，其他字段无效。 third_party：第三方链接视频会议，取该类型时，icon_type、description、meeting_url字段生效。 no_meeting：无视频会议，取该类型时，其他字段无效。 lark_live：飞书直播，内部类型，飞书客户端使用，API不支持创建，只读。 unknown：未知类型，做兼容使用，飞书客户端使用，API不支持创建，只读。
	IconType    string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型；可以为空，为空展示默认icon。vc：飞书视频会议icon live：直播视频会议icon default：默认icon
	Description string `json:"description,omitempty"` // 第三方视频会议文案，可以为空，为空展示默认文案 数据校验规则： 长度范围：0 ～ 500 字符
	MeetingUrl  string `json:"meeting_url,omitempty"` // 视频会议URL  数据校验规则： 长度范围：1 ～ 2000 字符
}

// CalendarEventsLocation 日程地点
type CalendarEventsLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称 数据校验规则： 长度范围：1 ～ 512 字符
	Address   string  `json:"address,omitempty"`   // 地点地址  数据校验规则： 长度范围：1 ～ 255 字符
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息，对于国内的地点，采用GCJ-02标准，海外地点采用WGS84标准
}

// CalendarEventsReminders 日程提醒列表
type CalendarEventsReminders struct {
	Minutes int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量，正数时表示在日程开始前X分钟提醒，负数时表示在日程开始后X分钟提醒新建或更新日程时传入该字段，仅对当前身份生效 数据校验规则： 取值范围：-20160 ～ 20160
}

// CalendarEventsSchemas 日程自定义信息；控制日程详情页的ui展示
type CalendarEventsSchemas struct {
	UiName   string `json:"ui_name,omitempty"`   // UI名称。取值范围如下： ForwardIcon: 日程转发按钮 MeetingChatIcon: 会议群聊按钮 MeetingMinutesIcon: 会议纪要按钮 MeetingVideo: 视频会议区域 RSVP: 接受/拒绝/待定区域 Attendee: 参与者区域 OrganizerOrCreator: 组织者/创建者区域
	UiStatus string `json:"ui_status,omitempty"` // UI项自定义状态。目前只支持hide 可选值有： hide：隐藏显示 readonly：只读 editable：可编辑 unknown：未知UI项自定义状态，仅用于读取时兼容
	AppLink  string `json:"app_link,omitempty"`  // 按钮点击后跳转的链接 示例值："https://applink.feishu.cn/client/calendar/event/detail?calendarId=xxxxxx&key=xxxxxx&originalTime=xxxxxx&startTime=xxxxxx" 数据校验规则： 最大长度：2000 字符
}

// CreateCalendarEventsRes 创建日程-响应体
type CreateCalendarEventsRes struct {
	ResponseCode
	Data CreateCalendarEventsData `json:"data"`
}

// CreateCalendarEventsData 创建日程-响应体的data部分
type CreateCalendarEventsData struct {
	Event CreateCalendarEventsResEvent `json:"event"` // 新创建的日程实体
}

// CreateCalendarEventsResEvent 创建日程-响应体的data的日程部分
type CreateCalendarEventsResEvent struct {
	EventId          string `json:"event_id"`           // 日程ID
	Status           string `json:"status"`             // 日程标题
	IsException      bool   `json:"is_exception"`       // 日程是否是一个重复日程的例外日程
	RecurringEventId string `json:"recurring_event_id"` // 例外日程的原重复日程的event_id
	CreateCalendarEventsReqBody
}

// CreateCalendarEvent 创建日程
func (c *Client) CreateCalendarEvent(req CreateCalendarEventsReq) (*CreateCalendarEventsRes, error) {
	bodyByte, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPost, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/events",
		strings.NewReader(string(bodyByte)))

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data CreateCalendarEventsRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------删除日程------------------------------------------------------------

// DeleteCalendarEventsReq 删除日程-请求体对外结构
type DeleteCalendarEventsReq struct {
	CalendarId       string `json:"calendar_id"`       // 日历ID
	EventId          string `json:"event_id"`          // 日程ID
	NeedNotification bool   `json:"need_notification"` // 删除日程是否给日程参与人发送bot通知，默认为true 可选值有： true：true false：false
}

// DeleteCalendarEventsRes 删除日程-响应体
type DeleteCalendarEventsRes struct {
	ResponseCode
	Data interface{} `json:"data"`
}

// DeleteCalendarEvents 删除日程
func (c *Client) DeleteCalendarEvents(req DeleteCalendarEventsReq) (*DeleteCalendarEventsRes, error) {
	params := url.Values{}

	params.Add("need_notification", fmt.Sprintf("%v", req.NeedNotification))

	request, _ := http.NewRequest(http.MethodDelete, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/events/"+
		req.EventId, nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data DeleteCalendarEventsRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------获取日程------------------------------------------------------------

// GetCalendarEventsReq 获取日程-请求体对外结构
type GetCalendarEventsReq struct {
	CalendarId string `json:"calendar_id"` // 日历ID
	EventId    string `json:"event_id"`    // 日程ID
}

// GetCalendarEventsRes 获取日程-响应体
type GetCalendarEventsRes struct {
	ResponseCode
	Data GetCalendarEventsResData `json:"data"`
}

// GetCalendarEventsResData 创建日程-响应体的data部分
type GetCalendarEventsResData struct {
	Event GetCalendarEventsResDataEvent `json:"event"` // 新创建的日程实体
}

type GetCalendarEventsResDataEvent CreateCalendarEventsResEvent

// GetCalendarEvents 获取日程
func (c *Client) GetCalendarEvents(req GetCalendarEventsReq) (*GetCalendarEventsRes, error) {
	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/events/"+
		req.EventId, nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data GetCalendarEventsRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------获取日程列表---------------------------------------------------------

// GetCalendarEventsListReq 获取日程列表-请求体对外结构
type GetCalendarEventsListReq struct {
	CalendarId string `json:"calendar_id"` // 日历ID
	PageSize   int64  `json:"page_size"`   // 一次请求要求返回最大数量，默认500，取值范围为[50, 1000]
	AnchorTime string `json:"anchor_time"` // 拉取anchor_time之后的日程，为timestamp
	PageToken  string `json:"page_token"`  // 上次请求Response返回的分页标记，首次请求时为空
	SyncToken  string `json:"sync_token"`  // 上次请求Response返回的增量同步标记，分页请求未结束时为空
	StartTime  string `json:"start_time"`  // 日程开始Unix时间戳，单位为秒
	EndTime    string `json:"end_time"`    // 日程结束Unix时间戳，单位为秒
}

// GetCalendarEventsListRes 获取日程列表-响应体
type GetCalendarEventsListRes struct {
	ResponseCode
	Data GetCalendarEventsListResData `json:"data"`
}

// GetCalendarEventsListResData 获取日程列表-响应体的data部分
type GetCalendarEventsListResData struct {
	HasMore   bool                                `json:"has_more"`
	PageToken string                              `json:"page_token"`
	SyncToken string                              `json:"sync_token"`
	Items     []GetCalendarEventsListResDataEvent `json:"items"`
}

type GetCalendarEventsListResDataEvent CreateCalendarEventsResEvent

// GetCalendarEventsList 获取日程列表
func (c *Client) GetCalendarEventsList(req GetCalendarEventsListReq) (*GetCalendarEventsListRes, error) {
	params := url.Values{}
	if req.SyncToken != "" {
		params.Add("sync_token", req.SyncToken)
	}
	if req.PageToken != "" {
		params.Add("page_token", req.PageToken)
	}
	if req.PageSize > 0 {
		params.Add("page_size", fmt.Sprintf("%v", req.PageSize))
	}
	if req.AnchorTime != "" {
		params.Add("anchor_time", req.AnchorTime)
	}
	if req.StartTime != "" {
		params.Add("start_time", req.StartTime)
	}
	if req.EndTime != "" {
		params.Add("end_time", req.EndTime)
	}

	request, _ := http.NewRequest(http.MethodGet, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/events?"+
		params.Encode(), nil)

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data GetCalendarEventsListRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

// --------------------------------------------------更新日程------------------------------------------------------------

// UpdateCalendarEventsReq 更新日程-请求体对外结构
type UpdateCalendarEventsReq struct {
	CalendarId string                      `json:"-"`        // 日历ID
	EventId    string                      `json:"event_id"` // 日程ID
	Body       UpdateCalendarEventsReqBody `json:"body"`     // 请求体
}

type UpdateCalendarEventsReqBody CreateCalendarEventsReqBody

// UpdateCalendarEventsRes 更新日程-响应体
type UpdateCalendarEventsRes struct {
	ResponseCode
	Data UpdateCalendarEventsResData `json:"data"`
}

type UpdateCalendarEventsResData CreateCalendarEventsData

// UpdateCalendarEvents 更新日程
func (c *Client) UpdateCalendarEvents(req UpdateCalendarEventsReq) (*UpdateCalendarEventsRes, error) {
	bodyByte, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest(http.MethodPatch, ServerUrl+"/open-apis/calendar/v4/calendars/"+req.CalendarId+"/events/"+
		req.EventId, strings.NewReader(string(bodyByte)))

	AccessToken, err := c.TokenManager.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(request, AccessToken)
	if err != nil {
		return nil, err
	}
	var data UpdateCalendarEventsRes
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}
