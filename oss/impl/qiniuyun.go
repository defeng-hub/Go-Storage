package impl

import (
	"context"
	"fmt"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/defeng-hub/Go-Storage/oss"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
)

var _ oss.Service = (*QiniuYunImpl)(nil)
var (
	AccessKey string = "xxx"
	SecretKey string = "xxx"
	Bucket    string = "defenga"
	Config           = storage.Config{}
	Address   string = "https://pic.bythedu.com"
	mac       *qbox.Mac
)

type QiniuYunImpl struct {
}

func init() {

	// 空间对应的机房
	Config.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	Config.UseHTTPS = true
	// 上传是否使用CDN上传加速
	Config.UseCdnDomains = false
	mac = qbox.NewMac(AccessKey, SecretKey)

}

func (q QiniuYunImpl) CreateFile(ctx context.Context, file *model.File) (*model.File, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	//文件路径
	localFile := currentDir + "/upload/" + file.Link
	key := "gofile/" + file.Link

	putPolicy := storage.PutPolicy{Scope: Bucket}

	upToken := putPolicy.UploadToken(mac)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&Config)
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(ret.Key, ret.Hash)
	file.Url = Address + "/" + ret.Key
	return file, nil
}

func (q QiniuYunImpl) DeleteFile(ctx context.Context, file *model.File) (*model.File, error) {
	bucketManager := storage.NewBucketManager(mac, &Config)
	key := "gofile/" + file.Link
	err := bucketManager.Delete(Bucket, key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}
