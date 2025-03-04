package note

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
)

func PdfContentHandler(ctx context.Context, c *app.RequestContext) {
	path := c.Query("pdfPath")
	fullPath := filepath.Join(NConf.NoteBasePath, path)
	fmt.Println(fullPath)
	c.Header("Content-Type", "application/pdf")
	c.File(fullPath)
}
