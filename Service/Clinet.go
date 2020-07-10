package Service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"go-umeng-push/Constants"
	"go-umeng-push/Responses/TaskPush"
	"go-umeng-push/Responses/Upload"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
	"time"
)

type abstractNotification struct {
	appKey     string
	appSecret  string
	envMode    string // 可选，正式/测试模式。默认为true @EnvModeConstants
	httpClient http.Client
}

func newNotification(appKey, appSecret, envMode string) *abstractNotification {
	if envMode == "" {
		envMode = Constants.TEST
	}

	result := abstractNotification{}
	result.appKey = appKey
	result.appSecret = appSecret
	result.envMode = envMode
	return &result
}

func (n *abstractNotification) SetApp(appkey, appSecret string) *abstractNotification {
	n.appKey = appkey
	n.appSecret = appSecret
	return n
}
func (n *abstractNotification) sent(sentUrl string, params map[string]string) (*http.Response, error) {
	var resp *http.Response
	var err error

	url, err := url2.Parse(sentUrl)
	if err != nil {
		return resp, err
	}

	q := url.Query()
	q.Set("sign", n.getSignature(sentUrl, params))
	url.RawQuery = q.Encode()
	body, err := json.Marshal(params)
	resp, err = n.httpClient.Post(url.String(), "application/json", strings.NewReader((string(body))))
	if err != nil {
		return resp, err
	}
	return resp, err

}
func (n *abstractNotification) getSignature(url string, params map[string]string) string {

	jsonParams, _ := json.Marshal(params)
	md5 := md5.New()
	data := http.MethodPost + url + string(jsonParams) + n.appSecret
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}

func (n *abstractNotification) statusQuery(taskId string) (*http.Response, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := map[string]string{
		"appkey":    n.appKey,
		"timestamp": timestamp,
		"task_id":   taskId,
	}

	return n.sent(Constants.HOST_URL+Constants.STATUS_URI, params)
}

// 任务类消息取消
//@see https://developer.umeng.com/docs/66632/detail/68343#h2-u4EFBu52A1u7C7Bu6D88u606Fu53D6u6D886
func (n *abstractNotification) ChancelPush(taskId string) (result *TaskPush.TaskPush, err error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := map[string]string{
		"appkey":    n.appKey,
		"timestamp": timestamp,
		"task_id":   taskId,
	}
	response, err := n.sent(Constants.HOST_URL+Constants.CANCEL_URI, params)
	if err != nil {
		return result, err
	}
	return TaskPush.New(response)

}

/**
 * 文件上传
 * 功能说明
 * 文件上传接口支持两种应用场景：
 * 发送类型为”filecast”的时候, 开发者批量上传device_token;
 * 发送类型为”customizedcast”时, 开发者批量上传alias。
 * 上传文件后获取file_id, 从而可以实现通过文件id来进行消息批量推送的目的。
 * 文件自创建起，服务器会保存两个月。开发者可以在有效期内重复使用该file-id进行消息发送。
 * @param $deviceToken 可丟入 device token or alias
 * @return Upload
 * @see https://developer.umeng.com/docs/66632/detail/68343#h2-u6587u4EF6u4E0Au4F207
 */
func (n *abstractNotification) Upload(deviceTokens []string) (*Upload.Upload, error) {
	tokens := strings.Join(deviceTokens, "\r\n")
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := map[string]string{
		"appkey":    n.appKey,
		"timestamp": timestamp,
		"content":   tokens,
	}
	response, err := n.sent(Constants.HOST_URL+Constants.UPLOAD_URI, params)
	if err != nil {
		log.Fatalln(err)

	}
	return Upload.New(response)

}
