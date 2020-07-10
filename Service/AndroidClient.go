package Service

import (
	"encoding/json"
	"github.com/zileyuan/go-umeng-push/Constants"
	"github.com/zileyuan/go-umeng-push/Responses/Status"
	"github.com/zileyuan/go-umeng-push/Responses/TaskPush"
	"github.com/zileyuan/go-umeng-push/Responses/UniCast"
	"strconv"
	"strings"
	"time"
)

type Android struct {
	abstractNotification
}

type Policy struct {
	StartTime  string `json:"start_time"`   // 可选，定时发送时，若不填写表示立即发送。 定时发送时间不能小于当前时间 格式: "yyyy-MM-dd HH:mm:ss"。 注意，start_time只对任务类消息生效。
	ExpireTime string `json:"expire_time"`  // 可选，消息过期时间，其值不可小于发送时间或者 start_time(如果填写了的话) 如果不填写此参数，默认为3天后过期。格式同start_time
	OutBizNo   string `json:"out_biz_no"`   // 可选，开发者对消息的唯一标识，服务器会根据这个标识避免重复发送。
	MaxSendNum string `json:"max_send_num"` // 可选，发送限速，每秒发送的最大条数。最小值1000
}
type Body struct {
	Ticker      string `json:"ticker"`       //必填，通知栏提示文字
	Title       string `json:"title"`        // 必填，通知标题
	Text        string `json:"text"`         // 必填，通知文字描述
	Icon        string `json:"icon"`         // 可选，状态栏图标ID，R.drawable.[smallIcon]，如果没有，默认使用应用图标。图片要求为24*24dp的图标，或24*24px放在drawable-mdpi下。
	LargeIcon   string `json:"largeIcon"`    /// 可选，通知栏拉开后左侧图标ID，R.drawable.[largeIcon]，图片要求为64*64dp的图标，
	ImgUrl      string `json:"img"`          // 可选，通知栏大图标的URL链接。该字段的优先级大于largeIcon。
	Sound       string `json:"sound"`        // 可选，通知声音，R.raw.[sound]。 如果该字段为空，采用SDK默认的声音，即res/raw/下的 umeng_push_notification_default_sound声音文件。
	BuilderId   string `json:"builder_id"`   // 可选，默认为0，用于标识该通知采用的样式。使用该参数时，开发者必须在SDK里面实现自定义通知栏样式。
	PlayVibrate string `json:"play_vibrate"` // 可选，收到通知是否震动，默认为"true" @see https://github.com/3rdpay/xc-golang-umeng-push/blob/master/src/Constants/Status/Vibrate/VibrateStatusConstants.go
	PlayLights  string `json:"play_lights"`  // 可选，收到通知是否闪灯，默认为"true" @see https://github.com/3rdpay/xc-golang-umeng-push/blob/master/src/Constants/Status/Light/LightStatusConstants.go
	PlaySound   string `json:"play_sound"`   // 可选，收到通知是否发出声音，默认为"true" @see https://github.com/3rdpay/xc-golang-umeng-push/blob/master/src/Constants/Status/SoundStatusConstants.go
	AfterOpen   string `json:"after_open"`   // 可选，默认为"go_app" @see https://github.com/3rdpay/xc-golang-umeng-push/blob/master/src/Constants/NotifcationActiveConstants.go
	Url         string `json:"url"`          // 当after_open=go_url时，必填。  通知栏点击后跳转的URL，要求以http或者https开头
	Activity    string `json:"activity"`     // 当after_open=go_activity时，必填。
	Custom      string `json:"custom"`       // 当display_type=message时, 必填
}

type AnPayload struct {
	DisplayType string `json:"display_type"`
	Body        Body   `json:"body"`
	Extra       string `json:"extra"`
}

type AnCustomized struct {
	PushType     string   // 必填 @see https://github.com/3rdpay/xc-golang-umeng-push/blob/master/src/Constants/PushTypeCostants.go
	DeviceTokens []string // 当type=unicast时, 必填, 表示指定的单个设备  当type=listcast时, 必填, 要求不超过500个, 以英文逗号分隔
	AliasType    string   // 当type=customizedcast时, 必填
	Alias        string   // 当type=customizedcast时, 选填(此参数和file_id二选一)  开发者填写自己的alias, 要求不超过500个alias, 多个alias以英文逗号间隔
	FileId       string   // 当type=filecast时，必填，file内容为多条device_token，以回车符分割 当type=customizedcast时，选填(此参数和alias二选一)
	Filter       string   // 当type=groupcast时，必填，用户筛选条件，如用户标签、渠道等 @see https://developer.umeng.com/docs/66632/detail/68343#h2--g-14
}
type Option struct {
	Description string // 可选，发送消息描述，建议填写。
	MiPush      string // 可选，默认为false。当为true时，表示MIUI、EMUI、Flyme系统设备离线转为系统下发
	MiActivity  string // 可选，mipush值为true时生效，表示走系统通道时打开指定页面acitivity的完整包路径。
}

//New Android Client
func NewAndroidClient(appKey, appSecret, envMode string) *Android {
	notification := newNotification(appKey, appSecret, envMode)
	android := Android{*notification}
	return &android
}

//Broadcast廣播
func (a *Android) Broadcast(payload AnPayload, policy Policy, option Option) (response *TaskPush.TaskPush, err error) {
	params, err := a.getParams(payload, policy, AnCustomized{PushType: Constants.BROADCAST}, option)
	if err != nil {
		return response, err
	}
	httpResponse, err := a.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return TaskPush.New(httpResponse)
}

//單一推播
func (a *Android) UniCast(payload AnPayload, policy Policy, option Option) (response *UniCast.UniCast, err error) {
	params, err := a.getParams(payload, policy, AnCustomized{PushType: Constants.UNICAST}, option)
	if err != nil {
		return response, err
	}
	httpResponse, err := a.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return UniCast.New(httpResponse)
}

//清單推播
func (a *Android) ListPush(payload AnPayload, policy Policy, option Option) (response *UniCast.UniCast, err error) {
	params, err := a.getParams(payload, policy, AnCustomized{PushType: Constants.LISTS_PUSH}, option)
	if err != nil {
		return response, err
	}
	httpResponse, err := a.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return UniCast.New(httpResponse)
}

//檔案推播
func (a *Android) FilePush(payload AnPayload, policy Policy, option Option, fileIds []string) (response *TaskPush.TaskPush, err error) {
	ids := strings.Join(fileIds, ",")
	customized := AnCustomized{
		PushType: Constants.FILE_PUSH,
		FileId:   ids,
	}
	params, err := a.getParams(payload, policy, customized, option)
	if err != nil {
		return response, err
	}
	httpResponse, err := a.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return TaskPush.New(httpResponse)
}

//群組推播
func (a *Android) GroupPush(payload AnPayload, policy Policy, option Option, filter string) (response *TaskPush.TaskPush, err error) {
	customized := AnCustomized{
		PushType: Constants.CUSTOMIZED_PUSH,
		Filter:   filter,
	}
	params, err := a.getParams(payload, policy, customized, option)
	if err != nil {
		return response, err
	}
	httpResponse, err := a.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return TaskPush.New(httpResponse)
}

//客製推播
func (a *Android) CustomizedPush(payload AnPayload, policy Policy, option Option, aliasType, alias, fileIds string) (response *TaskPush.TaskPush, err error) {
	customized := AnCustomized{
		PushType:  Constants.CUSTOMIZED_PUSH,
		AliasType: aliasType,
		Alias:     alias,
		FileId:    fileIds,
	}
	params, err := a.getParams(payload, policy, customized, option)
	if err != nil {
		return response, err
	}
	httpResponse, err := a.sent(Constants.HOST_URL+Constants.PUSH_URI, params)

	return TaskPush.New(httpResponse)
}

//推播狀態查詢
func (a *Android) PushStatus(taskId string) (response *StatusResponse.AndroidStatusResponse, err error) {
	httpResponse, err := a.statusQuery(taskId)
	if err != nil {
		return response, err

	}
	return StatusResponse.NewAndroidStatusResponse(httpResponse)

}

//get params
func (a *Android) getParams(payload AnPayload, policy Policy, customized AnCustomized, option Option) (result map[string]string, err error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return result, err
	}
	policyByte, err := json.Marshal(policy)
	if err != nil {
		return result, err
	}

	result = map[string]string{
		"appkey":          a.abstractNotification.appKey,
		"timestamp":       strconv.FormatInt(time.Now().Unix(), 10),
		"type":            customized.PushType,
		"device_tokens":   strings.Join(customized.DeviceTokens, ","),
		"alias_type":      customized.AliasType,
		"alias":           customized.Alias,
		"file_id":         customized.FileId,
		"filter":          customized.Filter,
		"payload":         string(p),
		"policy":          string(policyByte),
		"production_mode": a.abstractNotification.envMode,
		"description":     option.Description,
		"mipush":          option.MiPush,
		"mi_activity":     option.MiActivity,
	}
	return result, err
}
