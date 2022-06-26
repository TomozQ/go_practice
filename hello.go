package main

import (
	"strconv"
	"fmt"
	"database/sql"
	// importされているパッケージと依存関係にあるパッケージを指定するときには_をつけて記述する。パッケージ内の関数を呼び出して使用するわけではないため、_なしで記述すると構文チェックが実行され使用していないパッケージと判断されてしまう。
	_"github.com/mattn/go-sqlite3" // github.com/mattn/go-sqlite3パッケージには、database/sqlを使う際にドライバとして使用されるプログラムが用意されている。このパッケージにある関数や構造体を直接利用するわけではなく、database/sqlにある機能を使う際に、依存関係にあるgithub.com/mattn/go-sqlite3パッケージの機能が呼び出される。
	"hello"
)

var query string = "select * from mydata where id = ?" // ? -> プレースホルダ...あとから値をはめ込む場所を指定

// Mydata is json structure
type Mydata struct {
	ID int
	Name string
	Mail string
	Age int
}

// Str get string value
func (m *Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

// FirebaseからJSONデータを取得する
func main(){
	// DBにアクセス 変数conにDBのポインタが代入される
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for true {
		s := hello.Input("id")
		if s == "" {
			break
		}
		
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		
		rs := con.QueryRow(query, n) // queryの ? に n がはめ込まれる
		if er != nil {
			panic(er)
		}
		
		var md Mydata
		er2 := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er2 != nil {
			panic(er2)
		}
		fmt.Println(md.Str())
	}
	fmt.Println("***end***")
}