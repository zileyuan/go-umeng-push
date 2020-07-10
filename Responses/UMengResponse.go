package Responses

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UMengResponse struct {
	response  *http.Response
	Content   string
	body      UmengResponseBody
	isConnect bool
}
type StatusResponseBody struct {
	Ret string `json:"ret"`
}
type UmengResponseBody struct {
	StatusResponseBody
	Data struct {
		ErrorMsg  string `json:"error_msg"`
		ErrorCode string `json:"error_code"`
	} `json:"data"`
}

func New(response *http.Response) UMengResponse {
	umeng := UMengResponse{
		response: response,
	}
	umeng.isConnect = (umeng.response.StatusCode == http.StatusOK) || (umeng.response.StatusCode == http.StatusBadRequest)
	if umeng.isConnect {
		content, err := umeng.readBody()
		if err != nil {
			log.Fatalln(err)
		} else {
			umeng.Content = content
			body := UmengResponseBody{}
			err := json.Unmarshal([]byte(umeng.Content), &body)
			if err != nil {
				log.Fatalln(err)

			} else {
				umeng.body = body

			}
		}
	}

	return umeng
}

//是否連線成功
func (umeng *UMengResponse) IsConnectSuccess() bool {
	return umeng.isConnect
}

func (umeng *UMengResponse) readBody() (string, error) {
	body, err := ioutil.ReadAll(umeng.response.Body)
	return string(body), err
}

//是否發生錯誤
func (umeng *UMengResponse) IsErrorOccur() bool {
	return umeng.response.StatusCode == http.StatusBadRequest
}

func (umeng *UMengResponse) GetHttpResponse() *http.Response {
	return umeng.response
}

//error Message
func (umeng *UMengResponse) GetErrorMessage() string {
	return umeng.body.Data.ErrorMsg
}

//error Code
func (umeng *UMengResponse) GetErrorCode() string {
	return umeng.body.Data.ErrorCode
}

func (umeng *UMengResponse) Close() {
	umeng.response.Body.Close()
}
