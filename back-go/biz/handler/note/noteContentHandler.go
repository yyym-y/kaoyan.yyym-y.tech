package note

import (
	"back/biz/middle"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetNoteContentHandler(ctx context.Context, c *app.RequestContext) {
	path := c.Query("notePath")
	fullPath := filepath.Join(NConf.NoteBasePath, path)
	fmt.Println(fullPath)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		c.JSON(400, middle.FailWithMsg("读取文件失败"))
		return
	}
	fmt.Println(string(content))
	c.JSON(200, middle.SuccessWithDate(string(content)))
}
