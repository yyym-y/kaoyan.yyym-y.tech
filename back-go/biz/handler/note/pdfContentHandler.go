package note

import (
	"back/biz/middle"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// PdfContentHandler handles the request to extract a specific page from a PDF file.
func PdfContentHandler(ctx context.Context, c *app.RequestContext) {
	// 获取查询参数
	path := c.Query("pdfPath")
	page := c.Query("page")
	fullPath := filepath.Join(NConf.NoteBasePath, path)

	// 验证文件存在性
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "PDF file not found")
		return
	}
	// 解析页面号
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum <= 0 {
		c.String(http.StatusBadRequest, "Invalid page number")
		return
	}
	// 提取指定页面
	extractedPage, err := ExtractPDFPage(fullPath, pageNum)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error extracting page")
		return
	}

	// 设置响应头并返回提取的页面
	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", extractedPage)
}

func ExtractPDFPage(pdfPath string, pageNum int) ([]byte, error) {
	// 创建临时目录以存储提取的页面
	tmpDir, err := os.MkdirTemp("", "extracted_page_")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %w", err)
	}
	// defer os.RemoveAll(tmpDir) // 确保在函数结束时删除临时目录
	tempName := middle.NewUUID() + filepath.Base(pdfPath)
	tempName = strings.Replace(tempName, ".pdf", "", -1)
	fmt.Println(tmpDir)

	f, err := os.Open(pdfPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = api.ExtractPages(f, tmpDir, tempName, []string{fmt.Sprintf("%d", pageNum)}, nil)
	if err != nil {
		return nil, fmt.Errorf("error extracting pages: %s", err)
	}
	fmt.Println(filepath.Join(tmpDir, fmt.Sprintf("%s_page_%d.pdf", tempName, pageNum)))
	extractedPage, err := os.ReadFile(filepath.Join(tmpDir, fmt.Sprintf("%s_page_%d.pdf", tempName, pageNum)))
	if err != nil {
		return nil, fmt.Errorf("failed to read extracted page: %w", err)
	}
	return extractedPage, nil
}
