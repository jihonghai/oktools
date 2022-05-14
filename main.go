package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

//go:embed templates
var tmpl embed.FS

//go:embed static
var assets embed.FS

func main() {
	gin.SetMode(Conf.App.Mode)

	if gin.Mode() == gin.ReleaseMode {
		gin.DisableConsoleColor()

		logfile := Conf.App.LogFile
		if logfile == "" {
			log.Fatalln("Please set the log file path!")
		}

		file, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			file, err = os.Create(logfile)
			if file == nil {
				log.Fatalln(err)
			}
		}

		defer func() {
			err := file.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()

		gin.DefaultWriter = file
		log.SetOutput(file)
	}

	r := initRouter()
	if err := r.Run(":" + Conf.Http.Port); err != nil {
		log.Fatalln("Something terrible happened:", err)
	}
}

func initRouter() *gin.Engine {
	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(tmpl, "templates/*.html")))
	r.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(assets))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Cache-Control", "max-age=3153600")
		c.File("./static/favicon.ico")
	})

	r.GET("/", Index)
	r.GET("/uptime", Uptime)
	r.GET("/ph/:size", PlaceHolder)
	r.GET("/base64", Base64)
	r.GET("/image2base64", Image2Base64)
	r.GET("/tinyimg", TinyImage)
	r.GET("/hash", Hash)
	r.GET("/file-hash", FileHash)
	r.GET("/json", JSONView)
	r.GET("/number", Number)
	r.GET("/placeholder", Placeholder)
	r.GET("/qrcode", QRCode)
	r.GET("/regex", Regex)
	r.GET("/timestamp", Timestamp)
	r.GET("/color", Color)
	r.GET("/aes", AES)
	r.GET("/des", DES)
	r.GET("/rsa", RSA)
	r.GET("/morse", Morse)
	r.GET("/url", URL)
	r.GET("/unicode", Unicode)
	r.GET("/json2go", JSON2GO)
	r.GET("/json2xml", JSON2XML)
	r.GET("/json2yaml", JSON2YAML)
	r.GET("/pdf2img", PDF2IMG)
	r.GET("/websocket", WebSocket)
	return r
}
