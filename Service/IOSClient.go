package Service

import (
	"encoding/json"
	"go-umeng-push/Constants"
	"go-umeng-push/Responses/Status"
	"go-umeng-push/Responses/TaskPush"
	"go-umeng-push/Responses/UniCast"
	"strconv"
	"strings"
	"time"
)

type IOSClient struct {
	abstractNotification
}

func NewIOSClient(appKey, appSecret, envMode string) *IOSClient {
	notification := newNotification(appKey, appSecret, envMode)
	ios := IOSClient{*notification}
	return &ios
}

type AlertParams struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Body     string `json:"body"`
}
type ApsParams struct {
	Alert            AlertParams `json:"alert"`             // 当content-available=1时(静默推送)，可选; 否则必填。
	Badge            string      `json:"badge"`             // 可选
	Sound            string      `json:"sound"`             // 可选
	ContentAvailable int         `json:"content-available"` // 可选，代表静默推送
	Category         string      `json:"category"`          // 可选，注意: ios8才支持该字段。
}
type PolicyParams struct {
	StartTime      string `json:"start_time"`       // 可选，定时发送时间，若不填写表示立即发送。
	ExpireTime     string `json:"expire_time"`      // 可选，消息过期时间，其值不可小于发送时间或者
	OutBizNo       string `json:"out_biz_no"`       // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送。
	ApnsCollapseId string `json:"apns_collapse_id"` // 可选，多条带有相同apns_collapse_id的消息，iOS设备仅展示
}
type Payload struct {
	Aps    ApsParams    `json:"aps"`
	Policy PolicyParams `json:"policy"`
}

type Customized struct {
	DeviceTokens string `json:"device_tokens"` /// 当type=unicast时, 必填, 表示指定的单个设备 当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
	AliasType    string `json:"alias_type"`    // 当type=customizedcast时, 必填
	Alias        string `json:"alias"`         // 当type=customizedcast时, 选填(此参数和file_id二选一)
	FileId       string `json:"file_id"`       // 当type=filecast时，必填，file内容为多条device_token，以回车符分割
	Description  string `json:"description"`   // 可选，发送消息描述，建议填写。
	Filter       string `json:"file_id"`       // 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等，参考附录G。@see https://developer.umeng.com/docs/66632/detail/68343#h2--g-14
}

//廣播
func (c *IOSClient) Broadcast(p Payload) (*TaskPush.TaskPush, error) {
	var result TaskPush.TaskPush
	var err error
	params, err := c.getParams(p, Constants.BROADCAST, Customized{})

	if err != nil {
		return &result, err
	}

	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	if err != nil {
		return &result, err
	}
	return TaskPush.New(response)
}

// 單一裝置推播
func (c *IOSClient) UniCast(p Payload, deviceToken string) (result *UniCast.UniCast, err error) {
	params, err := c.getParams(p, Constants.UNICAST, Customized{DeviceTokens: deviceToken})
	if err != nil {
		return result, err
	}
	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)
	if err != nil {
		return result, err
	}

	return UniCast.New(response)

}

// 清單推播
func (c *IOSClient) ListPush(p Payload, deviceToken []string) (result *UniCast.UniCast, err error) {
	tokens := strings.Join(deviceToken, ",")
	params, err := c.getParams(p, Constants.LISTS_PUSH, Customized{DeviceTokens: tokens})
	if err != nil {
		return result, err
	}
	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)
	if err != nil {
		return result, err
	}

	return UniCast.New(response)
}

//客製化推播
//@params aliasType 或 alias 其一需必填
func (c *IOSClient) CustomizedPush(p Payload, aliasType, alias string, fileIds []string) (result *TaskPush.TaskPush, err error) {
	ids := strings.Join(fileIds, ",")

	custom := Customized{
		Alias:     alias,
		AliasType: aliasType,
		FileId:    ids,
	}

	params, err := c.getParams(p, Constants.CUSTOMIZED_PUSH, custom)
	if err != nil {
		return result, err
	}
	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)
	if err != nil {
		return result, err
	}

	return TaskPush.New(response)
}

//群組推播
func (c *IOSClient) GroupPush(p Payload, filter string) (result *TaskPush.TaskPush, err error) {
	custom := Customized{
		Filter: filter,
	}
	params, err := c.getParams(p, Constants.GROUP_PUSH, custom)
	if err != nil {
		return result, err
	}
	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)
	if err != nil {
		return result, err
	}
	return TaskPush.New(response)

}

func (c *IOSClient) getParams(p Payload, pushType string, customized Customized) (map[string]string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return map[string]string{}, err
	}

	params := map[string]string{
		"appkey":          c.abstractNotification.appKey,
		"timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"type":            pushType,
		"device_tokens":   customized.DeviceTokens,
		"alias_type":      customized.AliasType,
		"alias":           customized.Alias,
		"file_id":         customized.FileId,
		"production_mode": c.abstractNotification.envMode,
		"description":     customized.Description,
		"filter":          customized.Filter,
		"payload":         string(b),
	}
	return params, err

}

//群組推播
func (c *IOSClient) PushStatus(taskId string) (result *StatusResponse.IOSStatusResponse, err error) {
	response, err := c.statusQuery(taskId)
	if err != nil {
		return result, err

	}
	return StatusResponse.NewIOSStatusResponse(response)

}

//檔案推播
func (c *IOSClient) FilePush(p Payload, fileIds []string) (result *TaskPush.TaskPush, err error) {
	ids := strings.Join(fileIds, ",")
	customized := Customized{
		FileId: ids,
	}
	params, err := c.getParams(p, Constants.FILE_PUSH, customized)
	if err != nil {
		return result, err
	}
	response, err := c.abstractNotification.sent(Constants.HOST_URL+Constants.PUSH_URI, params)
	if err != nil {
		return result, err
	}
	return TaskPush.New(response)

}
