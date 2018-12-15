package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
)

//数据库配置
const (


	userName = "root"
	password = "root"
	ip = "192.168.1.214"
	port = "3306"
	dbName = "foreign-project"
)
//Db数据库连接池
var DB *sql.DB


func main() {
	//其实是可以直接作为接口 使用的  可以直接 连接 sql  直接 获取数据  test
	var sdds  = "taatta"
	_= sdds

	e := echo.New()
	//返回字符串
	e.GET("/", func(c echo.Context) error {
		println ("test")
		return c.String(http.StatusOK, "Hello, World!")
	})

	//读取内存
	e.GET("/aaa", func(c echo.Context) error {
		return c.String(http.StatusOK,sdds)
	})

	e.GET("/sql",func(c echo.Context) error {
		var sql = "dasa"
		println("sql get")
		return c.String(http.StatusOK,sql)
	})

	e.GET("/sql1", func(c echo.Context) error {
		return c.String(http.StatusOK,sdds+"ddd")
	})
	//cao  最后这句话  必须放在最后一行  要不不执行
	e.Logger.Fatal(e.Start(":1323"))


	////构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//
	////打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	//DB, _ = sql.Open("mysql", path)
	////设置数据库最大连接数
	//DB.SetConnMaxLifetime(100)
	////设置上数据库最大闲置连接数
	//DB.SetMaxIdleConns(10)
	////验证连接
	//if err := DB.Ping(); err != nil{
	//	fmt.Println("opon database fail")
	//	return
	//}
	//fmt.Println("connnect success")

	//Sql = "SELECT Id,English,Chinese,FilePth from ccbTranslate WHERE Id IS NOT NULL and Id != 0"

}
