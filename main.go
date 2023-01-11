package main

import (
	"embed"
	"fmt"
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/defeng-hub/Go-Storage/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"os"
	"strconv"
)

//go:embed h5
var FS embed.FS

func loadTemplate() *template.Template {
	var funcMap = template.FuncMap{
		"unescape": common.UnescapeHTML,
	}
	var ShowImg = template.FuncMap{
		"ShowImg": func(filetype int) string {
			return model.ShowImg[filetype]
		},
	}
	t := template.Must(template.New("").Funcs(funcMap).Funcs(ShowImg).ParseFS(FS, "h5/*.html"))
	return t
}

func main() {
	// Initialize SQL Database
	db, err := model.InitDB()
	if err != nil {
		fmt.Printf("数据库连接失败...")
		os.Exit(0)
	}
	defer db.Close()

	// Initialize Redis
	err = common.InitRedisClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Initialize options
	model.InitOptionMap()

	// Initialize HTTP server
	server := gin.Default()
	server.Static("/static", "./h5")
	server.SetHTMLTemplate(loadTemplate())

	// Initialize session store
	if common.RedisEnabled {
		opt := common.ParseRedisOption()
		store, _ := redis.NewStore(opt.MinIdleConns, opt.Network, opt.Addr, opt.Password, []byte(common.SessionSecret))
		server.Use(sessions.Sessions("session", store))
	} else {
		store := cookie.NewStore([]byte(common.SessionSecret))
		server.Use(sessions.Sessions("session", store))
	}
	router.SetRouter(server)

	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*common.Port)
	}
	if *common.Host == "localhost" {
		ip := common.GetIp()
		if ip != "" {
			*common.Host = ip
		}
	}
	serverUrl := "http://" + *common.Host + ":" + realPort + "/"
	if !*common.NoBrowser {
		common.OpenBrowser(serverUrl)
	}

	err = server.Run(":" + realPort)
	if err != nil {
		fmt.Printf("gin启动失败...")
		os.Exit(0)
	}
}
