package util

type UserInfo struct{
	Uid int64 // uid
	Uname string // 昵称
	Avatar string // 头像
	Pwd string // 密码
}

/**
* 客户端与服务段交互消息结构
 */
type Msg struct{
	Type int32
	Data string
}

/**
* 用户发送消息结构
 */
type UserMsg struct{
	MsgType int32
	Data string
}

/**
* header
 */
type Header struct{
	Cookie map[string]string
	ContentType string
}
