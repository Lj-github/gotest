package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"log"
	"os"
	"time"
	"io"
	"html/template"
	"server/uitl"
	"encoding/json"
)

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "192.168.1.214"
	port     = "3306"
	dbName   = "foreign-project"

	txtFIle = "output/txt/"
	logFIle = "output/log/"
)

//Db数据库连接池
var DB *sql.DB

type TemplateRenderer struct {
	templates *template.Template
}

type PostData struct {
	Status  string `json:"status" xml:"status"`
	Email string `json:"email" xml:"email"`
}



// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

type PostContext struct {
	echo.Context
	username string
	content  string
}

func main() {

	logFile, _ := os.Create(logFIle + time.Now().Format("20060102") + ".log");

	//创建一个Logger
	//参数1：日志写入目的地
	//参数2：每条日志的前缀
	//参数3：日志属性
	loger := log.New(logFile, "log_", log.Ldate|log.Ltime|log.Lshortfile);
	//设置输出前缀
	loger.SetPrefix("log_");
	//输出一条日志
	loger.Output(2, "打印一条日志信息");

	//格式化输出日志
	loger.Printf("第%d行 内容:%s", 11, "aa");
	////等价于print();os.Exit(1);
	//loger.Fatal("我是错误");
	////等价于print();panic();
	//loger.Panic("我是错误");

	//其实是可以直接作为接口 使用的  可以直接 连接 sql  直接 获取数据  test
	var sdds = "taatta"
	_ = sdds
	e := echo.New()

	e.Static("/", "web/assets")

	renderer := &TemplateRenderer{
		//template 路径 需要写清楚 package 类似java  不能写绝对路径
		templates: template.Must(template.ParseGlob("web/template/*.html")),
	}
	e.Renderer = renderer
	//返回字符串
	e.GET("/", func(c echo.Context) error {
		loger.Output(2, "test")
		return c.String(http.StatusOK, "Hello, World!")
	})

	//读取内存
	e.GET("/aaa", func(c echo.Context) error {
		return c.String(http.StatusOK, sdds)
	})

	e.GET("/sql", func(c echo.Context) error {
		content := c.FormValue("username")
		loger.Output(2, content)
		return c.String(http.StatusOK, content)
	})

	e.GET("/sql1", func(c echo.Context) error {
		return c.String(http.StatusOK, sdds+"ddd")
	})

	// Named route "foobar"
	e.GET("/index", func(c echo.Context) error {

		loger.Output(2, "index")
		// rander  里面的HTML 需要 写 后缀
		return c.Render(http.StatusOK, "d3.html", "World")
	})

	//uitl.SaveFile(txtFIle+"dsa.txt", "dsdsds")

	e.POST("/post", func(context echo.Context) (err error) {
		loger.Output(4, "test")
		id := context.FormValue("username")
		content := context.FormValue("content")
		u := &PostData{
			Status:  "success",
			Email: "jon@labstack.com",
		}
		context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		context.Response().WriteHeader(http.StatusOK)

		uitl.SaveFile(txtFIle+id+".txt", content)
		return  json.NewEncoder(context.Response()).Encode(u)
	})

	//cao  最后这句话  必须放在最后一行  要不不执行
	e.Logger.Fatal(e.Start(":1323"))

	// d3  js  数据 可视化  展示
	// html  d3 demo
	e.GET("/d3demo", func(c echo.Context) error {
		return c.Render(http.StatusOK, "d3.html", "d3demo")
	})




}
