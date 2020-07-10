package Constants

//广播 向安装该App的所有设备发送消息
const BROADCAST = "broadcast"

//单播 向指定的设备发送消息
const UNICAST = "unicast"

//清單推播：向指定的一批设备发送消息。
const LISTS_PUSH = "listcast"

//群組推播：向满足特定条件的设备集合发送消息，例如: “特定版本”、”特定地域”等。
const GROUP_PUSH = "groupcast"

//文件推播：开发者将批量的device_token或者alias存放到文件，通过文件ID进行消息发送。
const FILE_PUSH = "filecast"

//自定义播(customizedcast)：开发者通过自有的alias进行推送，可以针对单个或者一批alias进行推送，也可以将alias存放到文件进行发送。
const CUSTOMIZED_PUSH = "customizedcast"
