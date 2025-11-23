package consts

const (
	// Header 请求头
	Header    = 0x7E7E
	Footer    = 0xF7F7
	MinLength = 24
)

const (
	HeaderLen   = 4
	SerialNoLen = 4
	LengthLen   = 4
	TypeLen     = 4
	CodeLen     = 4
	FootLen     = 4
	StructLen   = 24
	pkgMinLen   = 24
)

// MessageType 消息类型
type MessageType uint32

// 获取消息类型描述
func (mt MessageType) String(value string) string {
	switch mt {
	case Undefined:
	default:
		value = "未知协议"
	case Login:
		value = "登录"
	case LoginAck:
		value = "登录响应"
	case Logout:
		value = "登出"
	case LogoutAck:
		value = "登出响应"
	case SubDataMode:
		value = "数据订阅"
	case SubDataModeAck:
		value = "数据订阅响应"
	case ControlCmd:
		value = "控制命令"
	case ControlCmdAck:
		value = "控制命令响应"
	case Heartbeat:
		value = "心跳"
	case HeartbeatAck:
		value = "心跳响应"
	case TimeCheck:
		value = "时间同步"
	case TimeCheckAck:
		value = "时间同步响应"
	}

	return value
}

const (
	Undefined MessageType = iota

	Login    MessageType = 101
	LoginAck             = 102

	Logout    = 103
	LogoutAck = 104

	SubDataMode    = 201
	SubDataModeAck = 202

	ControlCmd    = 301
	ControlCmdAck = 302

	Heartbeat    = 401
	HeartbeatAck = 402

	TimeCheck    = 501
	TimeCheckAck = 502
)
