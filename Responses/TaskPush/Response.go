package TaskPush

import (
	"encoding/json"
	"github.com/zileyuan/go-umeng-push/Responses"
	"net/http"
)

type TaskPush struct {
	Responses.UMengResponse
	body body
}

type body struct {
	Responses.StatusResponseBody
	Data struct {
		TaskID    string `json:"task_id"`
		ErrorMsg  string `json:"error_msg"`
		ErrorCode string `json:"error_code"`
	} `json:"data"`
}

func New(response *http.Response) (*TaskPush, error) {

	task := TaskPush{}
	task.UMengResponse = Responses.New(response)
	body := body{}
	err := json.Unmarshal([]byte(task.UMengResponse.Content), &body)
	if err == nil {
		task.body = body
	}
	return &task, err
}

//任務ID
func (task *TaskPush) GetTaskId() string {
	return task.body.Data.TaskID
}

// Get All
func (task *TaskPush) All() map[string]string {

	return map[string]string{
		"errorCode":    task.GetErrorCode(),
		"errorMessage": task.GetErrorMessage(),
		"taskId":       task.GetTaskId(),
	}
}
