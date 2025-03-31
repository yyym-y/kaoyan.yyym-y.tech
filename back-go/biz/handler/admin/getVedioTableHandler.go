package admin

import (
	"back/biz/middle"
	"back/biz/model"
	respmodel "back/biz/respModel"
	"context"
	"net/http"
	"sort"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetVedioTableHandler(ctx context.Context, c *app.RequestContext) {
	var wg sync.WaitGroup
	wg.Add(2)

	vedios := make([]model.Vedio, 0)
	infos := make([]model.PlayInfo, 0)

	go func(vedios *[]model.Vedio) {
		defer wg.Done()
		model.DB.Find(vedios)
	}(&vedios)
	go func(infos *[]model.PlayInfo) {
		defer wg.Done()
		model.DB.Find(infos)
	}(&infos)

	wg.Wait()

	result := make([]respmodel.VedioTableItemResq, 0)
	mp := make(map[int][]model.PlayInfo)
	for _, info := range infos {
		if mp[info.VedioId] == nil {
			mp[info.VedioId] = make([]model.PlayInfo, 0)
		}
		mp[info.VedioId] = append(mp[info.VedioId], info)
	}
	for _, vedio := range vedios {
		item := respmodel.VedioTableItemResq{}
		item.VedioId = vedio.VedioId
		item.VedioName = vedio.VedioName
		item.VedioCover = vedio.VedioCover
		item.Visable = vedio.Visable
		item.Description = vedio.Description
		item.Episode = len(mp[item.VedioId])
		item.EpisodeInfos = mp[item.VedioId]
		item.TypeId = vedio.TypeId
		if item.EpisodeInfos == nil {
			item.EpisodeInfos = make([]model.PlayInfo, 0)
		}
		sort.Slice(item.EpisodeInfos, func(i, j int) bool {
			return item.EpisodeInfos[i].Num < item.EpisodeInfos[j].Num
		})
		result = append(result, item)
	}
	c.JSON(http.StatusOK, middle.SuccessWithDate(result))
}
