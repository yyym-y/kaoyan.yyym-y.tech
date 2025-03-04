package respmodel

import "back/biz/model"

type VedioTableItemResq struct {
	VedioId      int              `json:"vedioId"`
	VedioName    string           `json:"vedioName"`
	Visable      bool             `json:"visable"`
	Episode      int              `json:"episode"`
	TypeId       int              `json:"typeId"`
	VedioCover   string           `json:"vedioCover"`
	Description  string           `json:"description"`
	EpisodeInfos []model.PlayInfo `json:"episodeInfos"`
}
