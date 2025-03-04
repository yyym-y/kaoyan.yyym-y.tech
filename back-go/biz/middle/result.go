package middle

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func FailWithMsg(msg string) *Result {
	return &Result{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
}

func SuccessWithDate(data interface{}) *Result {
	return &Result{
		Code: 1,
		Msg:  "Success",
		Data: data,
	}
}

func SuccessWithNil() *Result {
	return &Result{
		Code: 1,
		Msg:  "Success",
		Data: nil,
	}
}
