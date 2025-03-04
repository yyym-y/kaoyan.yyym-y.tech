package note

import (
	"back/biz/middle"
	reqmodel "back/biz/reqModel"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
)

func SaveNoteContentHandler(ctx context.Context, c *app.RequestContext) {
	req := reqmodel.NoteSaveReq{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("Error bind"))
		return
	}
	// 构建文件的完整路径
	fullPath := filepath.Join(NConf.NoteBasePath, req.NotePath)
	fmt.Println(req.Content, fullPath)

	// 创建或打开文件
	file, err := os.Create(fullPath) // 使用 os.Create 创建文件，如果文件已存在则会清空内容
	if err != nil {
		c.JSON(500, middle.FailWithMsg("Error creating file: "+err.Error()))
		return
	}
	defer file.Close() // 确保在函数结束时关闭文件
	// 写入内容到文件
	if _, err := file.WriteString(req.Content); err != nil {
		c.JSON(500, middle.FailWithMsg("Error writing to file: "+err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(200, middle.SuccessWithNil())
}
