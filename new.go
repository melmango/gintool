package main

import (
	"os"
	"strings"
	"path"
	"fmt"
)

var newDoc4Web = `
New application created, name %s

The files/directories structure as follow:
	|- main.go
	|- app.conf
	|- assets
	    |- js
	    |- css
	    |- img
	|- templates
	    |- index.tpl
`
var newDoc4Api = `
New application created, name %s

The files/directories structure as follow:
	|- main.go
	|- app.conf
`

var errorMsg = `
No enough parameters for creating a new application, format 'gintool new [app_name]
`

func NewApp(args []string) {
	if (len(args) < 2) {
		println(errorMsg)
	}
	name := args[2]
	createFiles(name)
}

func createFiles(appName string) {
	curpath, _ := os.Getwd()
	targetPath := path.Join(curpath, appName)
	if isExist(targetPath) {
		fmt.Printf("[ERRO] Path (%s) already exists\n", targetPath)
		fmt.Printf("[WARN] Do you want to overwrite it? [yes|no]")
		if !askForConfirmation() {
			os.Exit(2)
		}
	}
	fmt.Printf("[INFO] Creating application,path : %s\n", targetPath)
	os.MkdirAll(targetPath, 0755)
	writeTofile(path.Join(targetPath, "main.go"), mainFileContent)
	writeTofile(path.Join(targetPath, "buildAndRun.sh"), strings.Replace(shell, "{{ .AppName }}", appName, -1))
	writeTofile(path.Join(targetPath, "app.conf"), strings.Replace(defaultConf, "{{ .AppName }}", appName, -1))
	os.Mkdir(path.Join(targetPath, "assets"), 0755)
	os.Mkdir(path.Join(targetPath, "assets", "js"), 0755)
	os.Mkdir(path.Join(targetPath, "assets", "css"), 0755)
	os.Mkdir(path.Join(targetPath, "assets", "img"), 0755)
	os.Mkdir(path.Join(targetPath, "templates"), 0755)
	writeTofile(path.Join(targetPath, "templates", "index.tmpl"), strings.Replace(indexHtml, "{{ .title }}", appName, -1))
	fmt.Printf("[INFO] Done, new app built with name '%s'\n", appName)

}

var defaultConf = `APP_NAME = {{ .AppName }}
HTTP_HOST = 0.0.0.0
HTTP_PORT = 8080
DEBUG = true
`

var mainFileContent = `package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itkinside/itkconfig"
	"net/http"
	"time"
	"fmt"
)

var AppConfig *AppConfigEntity

func main() {
	AppConfig = initConf()
	fmt.Printf("config : %s\n", AppConfig.HTTP_PORT)

	router := gin.Default()

	router.Static("/assets", "./assets")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"welcome": "Welcome! Thanks for using gintool",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", AppConfig.HTTP_HOST, AppConfig.HTTP_PORT),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func initConf() *AppConfigEntity {
	AppConfig := &AppConfigEntity{

	}
	itkconfig.LoadConfig("app.conf", AppConfig)
	if AppConfig.DEBUG == "false" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	fmt.Printf("config is : %s\n", AppConfig.DEBUG)
	return AppConfig
}

type AppConfigEntity struct {
	HTTP_HOST string
	HTTP_PORT string
	DEBUG     string
	APP_NAME  string
}
`

var shell = `go clean
go build
./{{ .AppName }}
`

var indexHtml = `
<html>
    <head>
        <meta name="viewport" content="width=device-width">
        <title>{{ .title }}</title>
    </head>
    <h1>
        {{ .welcome }}
    </h1>
</html>
`

func writeTofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}


