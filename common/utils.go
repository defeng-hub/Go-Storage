package common

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func OpenBrowser(url string) {
	switch runtime.GOOS {
	case "linux":
		_ = exec.Command("xdg-open", url).Start()
	case "windows":
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		_ = exec.Command("open", url).Start()
	}
}

func GetIp() (ip string) {
	ips, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return ip
	}

	for _, a := range ips {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				if strings.HasPrefix(ip, "10") {
					return
				}
				if strings.HasPrefix(ip, "172") {
					return
				}
				if strings.HasPrefix(ip, "192.168") {
					return
				}
				ip = ""
			}
		}
	}
	return
}

var sizeKB = 1024
var sizeMB = sizeKB * 1024
var sizeGB = sizeMB * 1024

func Bytes2Size(num int64) string {
	numStr := ""
	unit := "B"
	if num/int64(sizeGB) > 1 {
		numStr = fmt.Sprintf("%.2f", float64(num)/float64(sizeGB))
		unit = "GB"
	} else if num/int64(sizeMB) > 1 {
		numStr = fmt.Sprintf("%d", int(float64(num)/float64(sizeMB)))
		unit = "MB"
	} else if num/int64(sizeKB) > 1 {
		numStr = fmt.Sprintf("%d", int(float64(num)/float64(sizeKB)))
		unit = "KB"
	} else {
		numStr = fmt.Sprintf("%d", num)
	}
	return numStr + " " + unit
}

func Seconds2Time(num int) (time string) {
	if num/31104000 > 0 {
		time += strconv.Itoa(num/31104000) + " 年 "
		num %= 31104000
	}
	if num/2592000 > 0 {
		time += strconv.Itoa(num/2592000) + " 个月 "
		num %= 2592000
	}
	if num/86400 > 0 {
		time += strconv.Itoa(num/86400) + " 天 "
		num %= 86400
	}
	if num/3600 > 0 {
		time += strconv.Itoa(num/3600) + " 小时 "
		num %= 3600
	}
	if num/60 > 0 {
		time += strconv.Itoa(num/60) + " 分钟 "
		num %= 60
	}
	time += strconv.Itoa(num) + " 秒"
	return
}

func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	case int:
		return fmt.Sprintf("%d", inter.(int))
	case float64:
		return fmt.Sprintf("%f", inter.(float64))
	}
	return "Not Implemented"
}

func UnescapeHTML(x string) interface{} {
	return template.HTML(x)
}

func IntMax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// AddFileName 给文件名加上日期后缀
func AddFileName(filename string) string {
	link := filename
	t := time.Now()
	timestamp := t.Format("_2006-01-02_15-04-05")
	ext := filepath.Ext(filename)
	if ext == "" {
		link += timestamp
	} else {
		link = filename[:len(filename)-len(ext)] + timestamp + ext
	}
	return link
}
