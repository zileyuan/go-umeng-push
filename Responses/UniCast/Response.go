package UniCast

import (
	"encoding/json"
	"go-umeng-push/Responses"
	"net/http"
)

type UniCast struct {
	Responses.UMengResponse
	body body
}

type body struct {
	Responses.StatusResponseBody
	Data struct {
		MessageId string `json:"msg_id"`
		ErrorMsg  string `json:"error_msg"`
		ErrorCode string `json:"error_code"`
	} `json:"data"`
}

func New(response *http.Response) (*UniCast, error) {

	task := UniCast{}
	task.UMengResponse = Responses.New(response)
	body := body{}
	err := json.Unmarshal([]byte(task.UMengResponse.Content), &body)
	if err == nil {
		task.body = body
	}
	return &task, err
}

/** 訊息ID */
func (task *UniCast) GetMessageId() string {
	return task.body.Data.MessageId
}

// Get All
func (u *UniCast) All() map[string]string {

	return map[string]string{
		"errorCode":    u.GetErrorCode(),
		"errorMessage": u.GetErrorMessage(),
		"messageId":    u.GetMessageId(),
	}
}
