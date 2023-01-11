package controller

import (
	"fmt"
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetExplorerPageOrFile(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	path, _ = url.PathUnescape(path)

	fullPath := filepath.Join(common.ExplorerRootPath, path)
	if !strings.HasPrefix(fullPath, common.ExplorerRootPath) {
		// We may being attacked!
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": fmt.Sprintf("只能访问指定文件夹的子目录"),
			"option":  common.OptionMap,
		})
		return
	}
	root, err := os.Stat(fullPath)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": "处理路径时发生了错误，请确认路径正确",
			"option":  common.OptionMap,
		})
		return
	}
	if root.IsDir() {
		localFilesPtr, readmeFileLink, err := getData(path, fullPath)
		fmt.Println("ccc", localFilesPtr)
		fmt.Println("aaaaa", readmeFileLink)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"message": err.Error(),
				"option":  common.OptionMap,
			})
			return
		}

		c.HTML(http.StatusOK, "explorer.html", gin.H{
			"message":        "",
			"option":         common.OptionMap,
			"files":          localFilesPtr,
			"readmeFileLink": readmeFileLink,
		})
	} else {
		c.File(filepath.Join(common.ExplorerRootPath, path))
	}
}

func getDataFromFS(path string, fullPath string) (localFilesPtr *[]model.LocalFile, readmeFileLink string, err error) {
	var localFiles []model.LocalFile
	var tempFiles []model.LocalFile
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return
	}
	if path != "/" {
		parts := strings.Split(path, "/")
		// Add the special item: ".." which means parent dir
		if len(parts) > 0 {
			parts = parts[:len(parts)-1]
		}
		parentPath := strings.Join(parts, "/")
		parentFile := model.LocalFile{
			Name:         "..",
			Link:         "explorer?path=" + url.QueryEscape(parentPath),
			Size:         "",
			IsFolder:     true,
			ModifiedTime: "",
		}
		localFiles = append(localFiles, parentFile)
		path = strings.Trim(path, "/") + "/"
	} else {
		path = ""
	}
	for _, f := range files {
		link := "explorer?path=" + url.QueryEscape(path+f.Name())
		file := model.LocalFile{
			Name:         f.Name(),
			Link:         link,
			Size:         common.Bytes2Size(f.Size()),
			IsFolder:     f.Mode().IsDir(),
			ModifiedTime: f.ModTime().String()[:19],
		}
		if file.IsFolder {
			localFiles = append(localFiles, file)
		} else {
			tempFiles = append(tempFiles, file)
		}
		if f.Name() == "README.md" {
			readmeFileLink = link
		}
	}
	localFiles = append(localFiles, tempFiles...)
	localFilesPtr = &localFiles
	return
}

func getData(path string, fullPath string) (localFilesPtr *[]model.LocalFile, readmeFileLink string, err error) {
	return getDataFromFS(path, fullPath)
}
