package reqmodel

type LoginReq struct {
	Key      string `json:"key"`
	Password string `json:"password"`
}
