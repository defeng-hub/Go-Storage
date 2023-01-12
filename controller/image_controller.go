package controller

import (
	"context"
	"fmt"
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//type ImageDeleteRequest struct {
//	Filename string `json:"filename"`
//	//Token    string
//}

func UploadImage(c *gin.Context) {
	GetOssImpl()
	uploadPath := common.UploadPath
	description := c.PostForm("description")
	uploader := c.GetString("username")
	if uploader == "" {
		uploader = "匿名用户"
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["file"]
	createTextFile := false
	if files == nil && description != "" {
		createTextFile = true
		file := &multipart.FileHeader{
			Filename: "text.txt",
			Header:   nil,
			Size:     0,
		}
		files = append(files, file)
	}
	var filenames []string
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		link := common.AddFileName(filename)
		savePath := filepath.Join(uploadPath, link)
		if createTextFile {
			// Create a new text file and then write the description to it.
			filename = "文本分享"
			f, err := os.Create(savePath)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create file: %s", err.Error()))
				return
			}
			_, err = f.WriteString(description)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to write text to file: %s", err.Error()))
				return
			}
			descriptionRune := []rune(description)
			if len(descriptionRune) > common.AbstractTextLength {
				description = fmt.Sprintf("内容摘要：%s...", string(descriptionRune[:common.AbstractTextLength]))
			}
		} else {
			if err := c.SaveUploadedFile(file, savePath); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to save uploaded file: %s", err.Error()))
				return
			}
		}
		fileObj := &model.File{
			Description: description,
			Uploader:    uploader,
			Time:        currentTime,
			Link:        link,
			Filename:    filename,
			FileType:    model.GetFileType(filename),
		}

		resFile, CreateErr := yunImpl.CreateFile(context.Background(), fileObj)
		if CreateErr != nil {
			return
		}
		err = resFile.Insert()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "上传失败:" + err.Error(),
				"data":    "",
			})
		}
		filenames = append(filenames, resFile.Url)
	}
	fmt.Println("cccccccc", filenames)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    filenames,
	})
	return
	//------------------------

}
