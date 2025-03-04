package note

import (
	"back/biz/config"
	"back/biz/middle"
	"context"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

type fileNode struct {
	Title    string     `json:"title"`
	Id       string     `json:"id"`
	FileType string     `json:"fileType"`
	Children []fileNode `json:"children"`
}

type NoteConf struct {
	NoteBasePath string
	FileTreeInfo []fileNode
}

var NConf NoteConf

func init() {
	NConf = NoteConf{}
	config.Add(&NConf)
}

func (cof *NoteConf) Update() {
	cof.NoteBasePath = config.Conf.NotePath
	cof.FileTreeInfo = make([]fileNode, 1)
	cof.FileTreeInfo[0] = fileNode{}
	cof.getFilePathInfo(cof.NoteBasePath, "", &cof.FileTreeInfo[0])
}

func (cof *NoteConf) getFilePathInfo(absolutePath string, relativePath string, node *fileNode) {
	files, err := ioutil.ReadDir(absolutePath)
	if err != nil {
		return
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		fullPath := filepath.Join(absolutePath, file.Name())
		temRelativePath := filepath.Join(relativePath, file.Name())
		childFileNode := fileNode{}
		childFileNode.Id = temRelativePath
		childFileNode.Title = file.Name()
		if file.IsDir() {
			childFileNode.FileType = "folder"
			cof.getFilePathInfo(fullPath, temRelativePath, &childFileNode)
			node.Children = append(node.Children, childFileNode)
		} else {
			temName := file.Name()
			if strings.HasSuffix(file.Name(), ".md") {
				temName = strings.Replace(file.Name(), ".md", "", -1)
				childFileNode.FileType = "md"
			} else if strings.HasSuffix(file.Name(), ".pdf") {
				temName = strings.Replace(file.Name(), ".pdf", "", -1)
				childFileNode.FileType = "pdf"
			} else {
				childFileNode.FileType = "unknow"
			}
			childFileNode.Title = temName
			node.Children = append(node.Children, childFileNode)
		}
	}
	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].FileType == "folder"
	})
}

func GetNotePathHandler(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, middle.SuccessWithDate(NConf.FileTreeInfo[0].Children))
}
