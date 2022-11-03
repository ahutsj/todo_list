package e

var MsgFlags = map[uint]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "params error",
}

func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
