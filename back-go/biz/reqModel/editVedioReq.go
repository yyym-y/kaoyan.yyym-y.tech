package reqmodel

import "back/biz/model"

type EditVedio struct {
	VedioId      int              `json:"vedioId"`
	VedioName    string           `json:"vedioName"`
	TypeId       int              `json:"typeId"`
	Visable      bool             `json:"visable"`
	Episode      int              `json:"episode"`
	VedioCover   string           `json:"vedioCover"`
	Description  string           `json:"description"`
	EpisodeInfos []model.PlayInfo `json:"episodeInfos"`
}
