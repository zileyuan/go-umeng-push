package StatusResponse

import (
	"encoding/json"
	"github.com/zileyuan/go-umeng-push/Responses"
	"net/http"
	"strconv"
)

type abstractStatusResponse struct {
	Responses.UMengResponse
	body abstractStatusBodyResponse
}
type abstractStatusBodyResponse struct {
	Ret  string `json:"ret"`
	Data struct {
		TaskID    string `json:"task_id"`
		OpenCount int    `json:"open_count"`
		SentCount int    `json:"sent_count"`
		Status    int    `json:"status"`
	} `json:"data"`
}

func newStatusResponse(response *http.Response) (abstractStatusResponse, error) {
	status := abstractStatusResponse{}
	status.UMengResponse = Responses.New(response)
	body := abstractStatusBodyResponse{}
	err := json.Unmarshal([]byte(status.UMengResponse.Content), &body)
	if err == nil {
		status.body = body
	}
	return status, err
}
func (r *abstractStatusResponse) GetSentCount() int {
	return r.body.Data.SentCount
}

func (r *abstractStatusResponse) GetMessageStatus() int {
	return r.body.Data.Status
}

func (r *abstractStatusResponse) GetOpenCount() int {
	return r.body.Data.OpenCount
}

type IOSStatusResponse struct {
	abstractStatusResponse
	body iosStatusBodyResponse
}
type iosStatusBodyResponse struct {
	Ret  string `json:"ret"`
	Data struct {
		TotalCount int    `json:"total_count"`
		TaskID     string `json:"task_id"`
		OpenCount  int    `json:"open_count"`
		SentCount  int    `json:"sent_count"`
		Status     int    `json:"status"`
	} `json:"data"`
}

func NewIOSStatusResponse(response *http.Response) (*IOSStatusResponse, error) {
	ios := IOSStatusResponse{}
	statusResponse, e := newStatusResponse(response)
	if e != nil {
		return &ios, e
	}
	ios.abstractStatusResponse = statusResponse
	body := iosStatusBodyResponse{}
	err := json.Unmarshal([]byte(ios.UMengResponse.Content), &body)
	if err == nil {
		ios.body = body
	}
	return &ios, err
}

// Get total
func (status *IOSStatusResponse) GetTotalCount() int {

	return status.body.Data.TotalCount
}

// Get All
func (status *IOSStatusResponse) All() map[string]string {

	return map[string]string{
		"errorCode":     status.GetErrorCode(),
		"errorMessage":  status.GetErrorMessage(),
		"messageStatus": strconv.Itoa(status.GetMessageStatus()),
		"openCount":     strconv.Itoa(status.GetOpenCount()),
		"sentCount":     strconv.Itoa(status.GetSentCount()),
	}
}

type AndroidStatusResponse struct {
	abstractStatusResponse
	body androidStatusBodyResponse
}

type androidStatusBodyResponse struct {
	Ret  string `json:"ret"`
	Data struct {
		DismissCount int    `json:"dismiss_count"`
		TaskID       string `json:"task_id"`
		OpenCount    int    `json:"open_count"`
		SentCount    int    `json:"sent_count"`
		Status       int    `json:"status"`
	} `json:"data"`
}

func NewAndroidStatusResponse(response *http.Response) (*AndroidStatusResponse, error) {
	android := AndroidStatusResponse{}
	a, e := newStatusResponse(response)
	if e != nil {
		return &android, e
	}
	android.abstractStatusResponse = a
	body := androidStatusBodyResponse{}
	err := json.Unmarshal([]byte(android.UMengResponse.Content), &body)
	if err == nil {
		android.body = body
	}
	return &android, err
}

// 忽略數
func (status *AndroidStatusResponse) GetDisMissCount() int {

	return status.body.Data.DismissCount
}

// Get All
func (status *AndroidStatusResponse) All() map[string]string {

	return map[string]string{
		"errorCode":     status.GetErrorCode(),
		"errorMessage":  status.GetErrorMessage(),
		"messageStatus": strconv.Itoa(status.GetOpenCount()),
		"openCount":     strconv.Itoa(status.GetOpenCount()),
		"sentCount":     strconv.Itoa(status.GetSentCount()),
		"disMissCount":  strconv.Itoa(status.GetSentCount()),
	}
}
