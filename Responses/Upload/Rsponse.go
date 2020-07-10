package Upload

import (
	"encoding/json"
	"go-umeng-push/Responses"
	"net/http"
)

type Upload struct {
	Responses.UMengResponse
	body responseBody
}

type responseBody struct {
	Ret  string `json:"ret"`
	Data struct {
		FileID    string `json:"file_id"`
		ErrorMsg  string `json:"error_msg"`
		ErrorCode string `json:"error_code"`
	} `json:"data"`
}

func New(response *http.Response) (*Upload, error) {

	upload := Upload{}
	upload.UMengResponse = Responses.New(response)
	body := responseBody{}
	err := json.Unmarshal([]byte(upload.UMengResponse.Content), &body)
	if err == nil {
		upload.body = body
	}
	return &upload, err
}

func (upload *Upload) GetFileId() string {
	return upload.body.Data.FileID
}

func (upload *Upload) All() map[string]string {
	return map[string]string{
		"errorCode":    upload.GetErrorCode(),
		"errorMessage": upload.GetErrorMessage(),
		"FileId":       upload.GetFileId(),
	}
}
