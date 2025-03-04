package respmodel

import "back/biz/model"

type VedioTypeResp struct {
	TypeInfos []model.VedioType `json:"typeInfos"`
}
