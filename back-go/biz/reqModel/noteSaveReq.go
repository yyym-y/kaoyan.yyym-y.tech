package reqmodel

type NoteSaveReq struct {
	NotePath string `json:"notePath"`
	Content  string `json:"content"`
}
