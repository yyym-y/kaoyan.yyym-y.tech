package reqmodel

type CreateVedioReq struct {
	VedioName   string `json:"vedioName"`
	VedioCover  string `json:"vedioCover"`
	Description string `json:"description"`
	Visable     bool   `json:"visable"`
	TypeId      int    `json:"typeId"`
}
