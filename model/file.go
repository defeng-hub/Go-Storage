package model

import (
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/jinzhu/gorm"
	"os"
	"path"
	"strings"
)

const (
	FileType_def = iota + 1
	FileType_img
	FileType_mp4
	FileType_txt
	FileType_pdf
	FileType_yasuo
)

var ShowImg = map[int]string{
	FileType_def:   "http://mms0.baidu.com/it/u=1782245366,4047385083&fm=253&app=138&f=JPEG&fmt=auto&q=75?w=300&h=300",
	FileType_img:   "",
	FileType_mp4:   "https://is2-ssl.mzstatic.com/image/thumb/Purple49/v4/c4/03/02/c40302b1-74c2-de7e-6f91-100ec644cee7/source/1024x1024bb.jpg",
	FileType_txt:   "/upload/txt.png",
	FileType_pdf:   "https://www.dugoogle.com/d/file/p/2021/03-15/052f3f1693520ec047b5413b50e33b37.png",
	FileType_yasuo: "https://bpic.588ku.com/element_origin_min_pic/00/15/97/3856aecc11b3d5b.jpg",
}

type File struct {
	Id              int    `json:"id"`
	Filename        string `json:"filename" gorm:"comment:'原始名称';"`
	Link            string `json:"link" gorm:"comment:'加长名称';"`
	Description     string `json:"description" gorm:"comment:'描述';"`
	Uploader        string `json:"uploader" gorm:"comment:'所属用户';"`
	Time            string `json:"time"`
	Url             string `json:"url" gorm:"comment:'图片地址';"`
	FileType        int    `json:"文件类型" gorm:"comment:'文件类型';default:1;"`
	DownloadCounter int    `json:"download_counter" gorm:"comment:'下载次数';"`
}

type LocalFile struct {
	Name         string
	Link         string
	Size         string
	IsFolder     bool
	ModifiedTime string
}

// GetFileType 获取文件类型
func GetFileType(name string) int {
	ext := path.Ext(name)
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif":
		return FileType_img
	case ".mp4", ".avi":
		return FileType_mp4
	case ".txt":
		return FileType_txt
	case ".pdf":
		return FileType_pdf
	case ".zip", ".7z", ".rar":
		return FileType_yasuo
	default:
		return FileType_def

	}
}

func AllFiles() ([]*File, error) {
	var files []*File
	var err error
	err = DB.Find(&files).Error
	return files, err
}

// QueryFiles 首页获取文件列表
func QueryFiles(query string, startIdx int) ([]*File, error) {
	var files []*File
	var err error
	query = strings.ToLower(query)
	err = DB.Limit(common.ItemsPerPage).Offset(startIdx).Where("filename LIKE ? or description LIKE ? or uploader LIKE ? or time LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Order("id desc").Find(&files).Error
	return files, err
}

func (file *File) Insert() error {
	var err error
	err = DB.Create(file).Error
	return err
}

// Delete Make sure link is valid! Because we will use os.Remove to delete it!
func (file *File) Delete() error {
	var err error
	err = DB.Delete(file).Error
	err = os.Remove(path.Join(common.UploadPath, file.Link))
	return err
}

func UpdateDownloadCounter(link string) {
	DB.Model(&File{}).Where("link = ?", link).UpdateColumn("download_counter", gorm.Expr("download_counter + 1"))
}
