package common

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path"
	"path/filepath"
	"time"
)

var StartTime = time.Now()
var Version = "v1.3.0"
var OptionMap map[string]string

// ItemsPerPage 首页每页显示12个
var ItemsPerPage = 12

// AbstractTextLength 文本摘要长度
var AbstractTextLength = 40

var ExplorerCacheEnabled = false // After my test, enable this will make the server slower...

var StatEnabled = false // 本来是true
var StatIPNum = 20
var StatURLNum = 20

const (
	RoleGuestUser  = 0  //游客权限
	RoleCommonUser = 1  // 用户权限
	RoleAdminUser  = 10 //管理员权限
)

var (
	FileUploadPermission    = RoleGuestUser
	FileDownloadPermission  = RoleGuestUser
	ImageUploadPermission   = RoleGuestUser
	ImageDownloadPermission = RoleGuestUser
)

var (
	GlobalApiRateLimit = 20
	GlobalWebRateLimit = 60
	DownloadRateLimit  = 10
	CriticalRateLimit  = 3
)

const (
	UserStatusEnabled  = 1 // 用户运行
	UserStatusDisabled = 2 // 用户禁用don't use 0
)

var (
	Port         = flag.Int("port", 3000, "specify the server listening port")
	Host         = flag.String("host", "localhost", "the server's ip address or domain")
	Path         = flag.String("path", "", "specify a local path to public")
	NoBrowser    = flag.Bool("no-browser", false, "open browser or not")
	PrintVersion = flag.Bool("version", false, "print version")
)

var UploadPath = "upload"
var ExplorerRootPath = "upload"
var ImageUploadPath = "upload"
var VideoServePath = "upload"

var SessionSecret = uuid.New().String()

func init() {
	flag.Parse()

	if *PrintVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	if os.Getenv("SESSION_SECRET") != "" {
		SessionSecret = os.Getenv("SESSION_SECRET")
	}

	if os.Getenv("UPLOAD_PATH") != "" {
		UploadPath = os.Getenv("UPLOAD_PATH")
		ExplorerRootPath = UploadPath
		ImageUploadPath = path.Join(UploadPath, "images")
		VideoServePath = UploadPath
	}
	if *Path != "" {
		ExplorerRootPath = *Path
	}

	ExplorerRootPath, _ = filepath.Abs(ExplorerRootPath)
	VideoServePath, _ = filepath.Abs(VideoServePath)
	ImageUploadPath, _ = filepath.Abs(ImageUploadPath)

	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(UploadPath, 0777)
	}
	if _, err := os.Stat(ImageUploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
}
