package note

import (
	"back/biz/middle"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/pkg/app"
)

type moveType struct {
	Action int    `json:"action"`
	Origin string `json:"origin"`
	Final  string `josn:"final"`
}

func createDir(fullPath string) error {
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		return err
	}
	return nil
}
func createFile(fullPath string) error {
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
func deleteFile(fullPath string) error {
	err := os.Remove(fullPath)
	if err != nil {
		return err
	}
	return nil
}

func moveFile(sourcePath string, targetPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close() // 确保在函数结束时关闭源文件
	destinationFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close() // 确保在函数结束时关闭目标文件
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return deleteFile(sourcePath)
}

func renameFile(sourcePath string, targetPath string) error {
	err := os.Rename(sourcePath, targetPath)
	if err != nil {
		return err
	}
	return nil
}

func FileChangeHandler(ctx context.Context, c *app.RequestContext) {
	req := moveType{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, middle.FailWithMsg("Error bind"))
		return
	}
	fullPath := filepath.Join(NConf.NoteBasePath, req.Final)
	fmt.Println(fullPath)
	switch req.Action {
	case 1:
		if err := createFile(fullPath); err != nil {
			c.JSON(500, middle.FailWithMsg("Error creating File: "+err.Error()))
			return
		}
		c.JSON(http.StatusOK, middle.SuccessWithNil())
	case 2:
		if err := createDir(fullPath); err != nil {
			c.JSON(500, middle.FailWithMsg("Error creating Dir: "+err.Error()))
			return
		}
		c.JSON(http.StatusOK, middle.SuccessWithNil())
	case 3:
		temSavePath := filepath.Join(TSconf.temSaveBaseUrl, req.Origin)
		if err := moveFile(temSavePath, fullPath); err != nil {
			c.JSON(500, middle.FailWithMsg("Error upload/[move] File: "+err.Error()))
			return
		}
		c.JSON(http.StatusOK, middle.SuccessWithNil())
	case 4:
		oriPath := filepath.Join(NConf.NoteBasePath, req.Origin)
		fmt.Println(oriPath)
		if err := renameFile(oriPath, fullPath); err != nil {
			c.JSON(500, middle.FailWithMsg("Error rename File: "+err.Error()))
			return
		}
		c.JSON(http.StatusOK, middle.SuccessWithNil())
	case 5:
		if err := deleteFile(fullPath); err != nil {
			c.JSON(500, middle.FailWithMsg("Error delete File: "+err.Error()))
			return
		}
		c.JSON(http.StatusOK, middle.SuccessWithNil())
	}
	NConf.Update()
}
