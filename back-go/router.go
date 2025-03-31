package main

import (
	adminHandler "back/biz/handler/admin"
	indexHandler "back/biz/handler/index"
	"back/biz/handler/note"
	"back/biz/handler/orc"
	vedioHandler "back/biz/handler/vedio"

	"back/biz/middle"
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func customizedRegister(r *server.Hertz) {

	auth, err := middle.GetJwt()
	if err != nil {
		panic(err)
	}
	errInit := auth.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	// test ping pong
	r.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	// r.GET("/test", noteHandler.GetNotePathHandler)

	// 公共访问路由
	r.GET("/showVedio", indexHandler.GetVedioHandler)
	r.POST("/login", auth.LoginHandler)
	r.GET("/image/:filename", indexHandler.ImageHandler)
	r.GET("/vedioType", indexHandler.GetVedioTypeHandler)
	r.GET("/showIndex", indexHandler.GetIndexHandler)
	r.GET("/showCarousel", indexHandler.GetCarouselHandler)

	// 视频播放路由
	vedio := r.Group("vedio")
	vedio.Use(auth.MiddlewareFunc())
	vedio.GET("/vedioPlay/:VedioId/:Episo/:filename", vedioHandler.VedioPlayHandler)
	vedio.POST("/getPlayList", vedioHandler.GetPlayListHandler)
	vedio.GET("/getVedioById", vedioHandler.GetVedioByIdHandler)

	// 管理员路由
	admin := r.Group("admin")
	admin.Use(auth.MiddlewareFunc())
	// 管理员通用路由
	admin.POST("/uploadImg", adminHandler.UploadImgHandler)
	// 管理员日志路由
	admin.GET("/developLog", adminHandler.GetDevelopContentHandler)
	admin.POST("/saveDevelopLog", adminHandler.SaveDevelopContentHandler)
	// 管理员影视管理路由
	// admin.POST("/uploadVedio", adminHandler.UploadVedioHandler)
	admin.POST("/createVedio", adminHandler.CreateVedioHandler)
	admin.GET("/getVedioTable", adminHandler.GetVedioTableHandler)
	// admin.POST("/editVedio", adminHandler.EditVedioHandler)
	// admin.POST("/createCarousel", adminHandler.CreateCarouselHandler)
	// admin.GET("/showVedio", adminHandler.ShowVedioHandler)
	// admin.GET("/vedioType", adminHandler.VedioTypeHandler)

	// 笔记路由
	noteGroup := r.Group("note")
	noteGroup.GET("/notePath", note.GetNotePathHandler)
	noteGroup.GET("/noteContent", note.GetNoteContentHandler)
	noteGroup.GET("/pdfContent", note.PdfContentHandler)
	// 笔记管理员操作路由
	noteAdminGroup := r.Group("note-admin")
	noteAdminGroup.Use(auth.MiddlewareFunc())
	noteAdminGroup.POST("/uploadNote", note.UploadNoteHandler)
	noteAdminGroup.POST("/saveContent", note.SaveNoteContentHandler)
	noteAdminGroup.POST("/fileChange", note.FileChangeHandler)

	orcGroup := r.Group("orc")
	orcGroup.Use(auth.MiddlewareFunc())
	orcGroup.POST("/uploadImg", orc.UploadImgHandler)
	orcGroup.POST("/latexOrc", orc.LatexOrcController)
}
