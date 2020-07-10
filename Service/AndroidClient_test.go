package Service

import (
	"fmt"
	"go-umeng-push/Constants"
)

func ExampleAndroid_Broadcast() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}
	push, _ := anClient.Broadcast(anPayload, anPolicy, anOption)

	defer push.Close()

	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())

	// Output:
	// true
	// true
	// 2021
	// appkey不存在
}
func ExampleAndroid_UniCast() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}
	push, _ := anClient.UniCast(anPayload, anPolicy, anOption)
	defer push.Close()

	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())

	// Output:
	// true
	// true
	// 2021
	// appkey不存在
}
func ExampleAndroid_ListPush() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}
	push, _ := anClient.ListPush(anPayload, anPolicy, anOption)
	defer push.Close()

	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetMessageId())

	// Output:
	// true
	// true
	// appkey不存在
	// 2021
}

func ExampleAndroid_UniCast2() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}
	push, _ := anClient.UniCast(anPayload, anPolicy, anOption)
	defer push.Close()

	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetMessageId())

	// Output:
	// true
	// true
	// appkey不存在
	// 2021
}
func ExampleAndroid_FilePush() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}
	fileId := []string{
		"fileId1",
		"fileId2",
	}
	push, _ := anClient.FilePush(anPayload, anPolicy, anOption, fileId)
	defer push.Close()

	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorMessage())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetTaskId())

	// Output:
	// true
	// true
	// appkey不存在
	// 2021
}
func ExampleAbstractNotification_SetApp() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)
	anClient.SetApp("set your new app", "your new secret")
	fmt.Println(anClient)
	// Output:
	//&{{set your new app your new secret test {<nil> <nil> <nil> 0}}}
}

func ExampleAndroid_CustomizedPush() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)
	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}

	var aliasType, alias, fileIds string
	push, _ := anClient.CustomizedPush(anPayload, anPolicy, anOption, aliasType, alias, fileIds)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())
	// Output:
	//true
	//true
	//2021
	//appkey不存在
}

func ExampleAndroid_GroupPush() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)
	anBody := Body{
		Ticker: "title",
		Title:  "subTitle",
		Text:   "Body",
	}
	anPayload := AnPayload{
		DisplayType: "message",
		Body:        anBody,
	}
	anPolicy := Policy{}
	anOption := Option{}

	var filter string
	push, _ := anClient.GroupPush(anPayload, anPolicy, anOption, filter)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())
	// Output:
	//true
	//true
	//2021
	//appkey不存在
}

func ExampleAbstractNotification_Upload() {
	anClient := NewAndroidClient("your app key ", "your secret", Constants.TEST)

	deviceToken := []string{
		"devcetoke1",
		"devcetoke2",
	}

	push, _ := anClient.Upload(deviceToken)
	defer push.Close()
	fmt.Println(push.IsConnectSuccess())
	fmt.Println(push.IsErrorOccur())
	fmt.Println(push.GetErrorCode())
	fmt.Println(push.GetErrorMessage())
	// Output:
	//true
	//true
	//2021
	//appkey不存在
}
