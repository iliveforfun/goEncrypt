package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"goEncrypt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 主程序
	db, err := sql.Open("mysql", "root:XC95098593@tcp(127.0.0.1:3306)/testforiot")
	// 打开连接  方法是 sql.Open 第一个参数是 数据库类型. 第二个是 用户名:密码@网络协议(ip:port)/需要查询的数据库名
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from test")
	if err != nil {
		log.Fatal(err)
	}
	//rows 查询 表里面所有的数据 结果应该是一个数组 方式db.Query
	var (
		id   string
		name string
	)
	for rows.Next() {
		rows.Scan(&id, &name)
	}
	var plaintext string = id + name
	fmt.Println("明文是:", plaintext, "\n")
	cryptText, err1 := goEncrypt.DesCbcEncrypt([]byte(plaintext), []byte("zhuhao12"))
	if err1 != nil {
		return
	}
	fmt.Println("密文是:", base64.StdEncoding.EncodeToString(cryptText), "\n")
	ans, err2 := goEncrypt.DesCbcDecrypt(cryptText, []byte("zhuhao12"))
	if err2 != nil {
		return
	}
	fmt.Println("明文是:", string(ans), "\n")
	dbinsert, err := db.Exec("insert into test(id,name) values('20190910','haha')")
	fmt.Println(dbinsert)
}
