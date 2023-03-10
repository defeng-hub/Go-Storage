package common

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var StartTime = time.Now()
var Version = "v3.0.0"
var OptionMap map[string]string
var ItemsPerPage = 12       // 首页每页显示12个
var AbstractTextLength = 40 //文本摘要长度
var StatIPNum = 20
var StatURLNum = 20

const (
	RoleGuestUser  = 0  //游客权限
	RoleCommonUser = 1  // 用户权限
	RoleAdminUser  = 10 //管理员权限
)

const (
	LocalStorage = 0 // 本地存储
	QiniuYun     = 1 // 七牛云
	AliYun       = 2 // 阿里云
	TxYun        = 3 // 腾讯云
)

var (
	// FileUploadPermission 上传文件的最低权限
	FileUploadPermission = RoleGuestUser

	// ImageUploadPermission 图库上传的最小权限
	ImageUploadPermission = RoleGuestUser

	// IndexPermission 首页访问最低权限
	IndexPermission = RoleCommonUser
)

const (
	UserStatusEnabled  = 1 // 用户运行
	UserStatusDisabled = 2 // 用户禁用don't use 0
)

// 解析入参  需要带上--,例如 --port
var (
	SqlDsn        = flag.String("Sql_Dsn", "", "数据库DSN")
	RunUrl        = flag.String("Run_Url", "http://127.0.0.1:3000", "当设置为本地存储时, 必须设置访问域名 会将其拼接成为 图片存储的url地址")
	Port          = flag.Int("Port", 3000, "specify the server listening port")
	Host          = flag.String("Host", "localhost", "the server's ip address or domain")
	IsOpenBrowser = flag.Bool("Is_Open_Browser", true, "open browser or not")
	OssType       = flag.Int("Oss_Type", LocalStorage, "存储引擎(默认为本地存储)")
	QiniuAK       = flag.String("Qiniu_AK", "", "七牛云AK")
	QiniuSK       = flag.String("Qiniu_SK", "", "七牛云SK")
	AliAK         = flag.String("Ali_AK", "", "阿里云AK")
	AliSK         = flag.String("Ali_SK", "", "阿里云SK")
	TxAK          = flag.String("Tx_AK", "", "腾讯云AK")
	TxSK          = flag.String("Tx_SK", "", "腾讯云SK")
)

var UploadPath = "upload"
var ExplorerRootPath = "upload"
var ImageUploadPath = "upload"
var VideoServePath = "upload"
var SessionSecret = uuid.New().String()

func init() {
	flag.Parse()
	fmt.Printf("cccc,%v\n", *RunUrl)

	// 获取 七牛云，阿里云，腾讯云的AK和SK
	getOssAKAndSK()

	//读取环境变量
	if os.Getenv("Oss_Type") != "" {
		atoi, _ := strconv.Atoi(os.Getenv("Oss_Type"))
		OssType = &atoi
	}
	if os.Getenv("Run_Url") != "" {
		atoi := os.Getenv("Run_Url")
		*RunUrl = atoi
	}
	if os.Getenv("SESSION_SECRET") != "" {
		SessionSecret = os.Getenv("SESSION_SECRET")
	}
	if os.Getenv("UPLOAD_PATH") != "" {
		UploadPath = os.Getenv("UPLOAD_PATH")
		ExplorerRootPath = UploadPath
		ImageUploadPath = UploadPath
		VideoServePath = UploadPath
	}
	ExplorerRootPath, _ = filepath.Abs(ExplorerRootPath)
	VideoServePath, _ = filepath.Abs(VideoServePath)
	ImageUploadPath, _ = filepath.Abs(ImageUploadPath)
	makeDir()
}

// 获取 七牛云，阿里云，腾讯云的AK和SK
func getOssAKAndSK() {
	// 七牛云获取ak sk
	if tmp := os.Getenv("Qiniu_AK"); tmp != "" {
		*QiniuAK = tmp
	}
	if tmp := os.Getenv("Qiniu_SK"); tmp != "" {
		*QiniuSK = tmp
	}

	// 阿里云获取ak sk
	if tmp := os.Getenv("Ali_AK"); tmp != "" {
		*AliAK = tmp
	}
	if tmp := os.Getenv("Ali_SK"); tmp != "" {
		*AliSK = tmp
	}

	// 腾讯云获取ak sk
	if tmp := os.Getenv("Tx_AK"); tmp != "" {
		*TxAK = tmp
	}
	if tmp := os.Getenv("Tx_SK"); tmp != "" {
		*TxSK = tmp
	}

}

// 创建目录，防止没有
func makeDir() {
	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(UploadPath, 0777)
	}
	if _, err := os.Stat(ImageUploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
	if _, err := os.Stat(VideoServePath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
	if _, err := os.Stat(ExplorerRootPath); os.IsNotExist(err) {
		_ = os.Mkdir(ExplorerRootPath, 0777)
	}
}
