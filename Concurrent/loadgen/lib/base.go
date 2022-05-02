package lib

import "time"

const ( // 状态量
	STATUS_ORIGINAL uint32 = iota
	STATUS_STARTING
	STATUS_STARTED
	STATUS_STOPPING
	STATUS_STOPEED
)

type CallResult struct {
	ID     int64
	Req    RawReq
	Resp   RawResp
	Code   RetCode
	Msg    string
	Elapse time.Duration
}

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type RetCode int // 返回代码的类型为 int

const (
	RET_CODE_SUCCESS              RetCode = 0    // 成功。
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001 // 调用超时警告。
	RET_CODE_ERROR_CALL                   = 2001 // 调用错误。
	RET_CODE_ERROR_RESPONSE               = 2002 // 响应内容错误。
	RET_CODE_ERROR_CALEE                  = 2003 // 被调用方（被测软件）的内部错误。
	RET_CODE_FATAL_CALL                   = 3001 // 调用过程中发生了致命错误！
)

// GetRetCodePlain 会依据结果代码返回相应的文字解释。
func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "Success"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "Call Timeout Warning"
	case RET_CODE_ERROR_CALL:
		codePlain = "Call Error"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "Response Error"
	case RET_CODE_ERROR_CALEE:
		codePlain = "Callee Error"
	case RET_CODE_FATAL_CALL:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown result code"
	}
	return codePlain
}

type Generator interface {
	Start() bool
	Stop() bool
	Status() bool
	CallCount() bool
}
