package Constants

//@see https://developer.umeng.com/docs/66632/detail/68343#h2-http-status-code-16
//请求参数没有appkey或为空值
const INVALID_APP_KEY = 1000

//请求参数没有payload或为非法json
const INVALID_PAYLOAD = 1001

//请求参数payload中, 没有body或为非法json
const CONTENT_TYPE_IS_NOT_JSON = 1002

//payload.display_type为message时, 请求参数payload.body中, 没有custom字段
const CUSTOM_PARAMS_IS_REQUIRED = 1003

//请求参数payload中, 没有display_type或为空值
const DISPLAY_TYPE_IS_REQUIRED = 1004

//请求参数payload.body中, img格式有误, 需以http或https开始
const INVALID_IMG_PARAMS = 1005

//payload.body.after_open为go_url时, 请求参数payload.body中, url格式有误, 需以http或https开始
const INVALID_GO_URL_PARAMS = 1007

//payload.display_type为notification时, 请求参数payload.body中, 没有ticker参数
const TICKER_IS_REQUIRED = 1008

//payload.display_type为notification时, 请求参数payload.body中, 没有title参数
const TITLE_IS_REQUIRED = 1009

//payload.display_type为notification时, 请求参数payload.body中, 没有text参数
const TEXT_IS_REQUIRED = 1010

//task_id对应任务没有找到
const NOT_FOUND_TASK_ID = 1014

//type为unicast或listcast时, 请求参数没有device_tokens或为空值
const DEVICE_TOKEN_IS_REQUIRED = 1015

//请求参数没有type或为空值
const TYPE_IS_REQUIRED = 1016

//请求参数payload中, display_type值非法
const INVALID_DISPLAY_TYPE = 1019

//应用组中尚未添加应用
const NO_HAVE_APPLICATION = 1020

//payload.body.after_open为go_url时, 请求参数payload.body中, 没有url参数或为空
const GO_URL_IS_REQUIRED = 1022

//payload.body.after_open为go_activity时, 请求参数payload.body中, 没有activity或为空值
const GO_ACTIVITY_IS_REQUIRED = 1024

//请求参数payload中, extra为非法json
const INVALID_EXTRA_PARAMS = 1026

//请请求参数payload中, policy为非法json
const INVALID_POLICY_PARAMS = 1027

//task_id对应任务无法撤销
const CAN_NOT_CANCEL_TASK_ID = 1028

//该应用已被禁用
const IS_DISABLE_APPLICATION = 2000

//请求参数policy中, start_time必须大于当前时间
const START_TIME_MUST_GREATER_THAN_CURRENT_TIME = 2002

//请求参数policy中, expire_time必须大于start_time和当前时间
const EXPIRE_TIME_MUST_GREATER_THANE_START_TIME = 2003

//IP白名单尚未添加, 请到网站后台添加您的服务器IP或关闭IP白名单功能
const NOT_YET_ADD_WHITE_IP_LIST = 2004

//Validation token不一致(PS: 此校验方法已废弃, 请采用sign进行校验)
const VALIDATION_TOKEN_FAIL = 2006

//未对请求进行签名
const SIGNATURE_IS_REQUIRED = 2007

//json解析错误
const JSON_FORMAT_ERROR = 2008

//type为customizedcast时, 请求参数没有alias、file_id或皆为空值
const ALIAS_PARAMS_OR_FILE_ID_IS_REQUIRED = 2009

//type为groupcast时, 请求参数没有filter或为非法json
const FILTER_IS_REQUIRED = 20016

//添加tag失败
const ADD_TAG_FAIL = 2017

//type为filecast时, 请求参数没有file_id或为空值
const FILE_ID_IS_REQUIRED = 2018

//type为filecast时, file_id对应的文件不存在
const FILE_ID_IS_NOT_FOUND = 2019

//appkey不存在
const APP_KEY_IS_NOT_FOUND = 2021

//payload长度过长
const PAYLOAD_LENGTH_IS_TOO_LONG = 2022

//文件上传失败, 请稍后重试
const UPLOAD_FAIL = 2023

//请求参数没有aps或为非法json
const APS_PARAMS_IS_REQUIRED = 2025

//签名不正确
const INVALID_SIGNATURE = 2027

//时间戳已过期
const TIMESTAMP_IS_EXPIRED = 2028

//请求参数没有content或为空值
const CONTENT_IS_REQUIRED = 2029

//filter格式不正确
const INVALID_FILTER = 2031

//未上传生产证书, 请到Web后台上传
const NOT_YET_UPLOAD_DISTRIBUTION_CERTIFICATE = 2032

//未上传开发证书, 请到Web后台上传
const NOT_YET_UPLOAD_DEVELOPMENT_CERTIFICATE = 2033

//证书已过期
const CERTIFICATE_IS_EXPIRED = 2034

//定时任务发送时, 证书已过期
const AT_DESIGNATED_TIME_CERTIFICATE_IS_EXPIRED = 2035

//时间戳格式错误
const TIMESTAMP_FORMAT_ERROR = 2036

//请求参数policy中, 时间格式必须是yyyy-MM-dd HH:mm:ss
const TIME_FORMAT_ERROR = 2039

//请求参数policy中, expire_time不能超过发送时间+7天
const EXPIRE_TIME_CAN_NOT_GREATER_THAN_SEVEN_DAY = 2040

//请求参数policy中, start_time不能超过当前时间+7天
const START_TIME_CAN_NOT_GRATER_THAN_SEVEN_DAY = 2046

//type为customizedcast时, 请求参数没有alias_type或为空值
const ALIAS_TYPE_IS_REQUIRED = 2047

//type值须为unicast、listcast、filecast、broadcast、groupcast、groupcast中的一种
const INVALID_TYPE = 2048

//type为customizedcast时, 请求参数alias、file_id只可二选一
const ALIAS_OR_FILE_ID_IS_REQUIRED = 2049

//发送频率超出限制
const PUSH_TIME_EXCEED_THE_LIMIT = 2050

//请求参数没有timestamp或为空值
const TIMESTAMP_IS_REQUIRED = 2052

//请求参数没有task_id或为空值
const TASK_ID_IS_REQUIRED = 2053

//IP不在白名单中, 请到网站后台添加您的服务器IP或关闭IP白名单功能
const IP_NOT_ON_THE_LIST = 2054

//证书解析bundle id失败, 请重新上传
const VALIDATION_BUNDLE_ID_FAIL = 5001

//请求参数payload中p、d为友盟保留字段
const P_AND_KEY_IS_RESERVED_WORD = 5002

//certificate_revoked错误
const CERTIFICATE_REVOKED_ERROR = 5007

//certificate_unkown错误
const CERTIFICATE_UNKNOWN_ERROR = 5008

//handshake_failure错误
const HANDSHAKE_FAILURE_ERROR = 5009

//配置使用Token Auth, 但未上传p8证书
const NOT_YET_UPLOAD_P8_CERTIFICATE = 5010

//该app未开通服务端tag接口
const APP_NOT_UNOPENED_TAG = 6001

//内部错误(iOS证书)
const INTERNAL_CERTIFICATE_ERROR = 6002

//内部错误(数据库)
const INTERNAL_DB_ERROR = 6003

//内部错误(TFS)
const INTERNAL_TFS_ERROR = 6004
