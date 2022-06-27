package main

import (
	"strconv"
	"fmt"
	"database/sql"
	// importされているパッケージと依存関係にあるパッケージを指定するときには_をつけて記述する。パッケージ内の関数を呼び出して使用するわけではないため、_なしで記述すると構文チェックが実行され使用していないパッケージと判断されてしまう。
	_"github.com/mattn/go-sqlite3" // github.com/mattn/go-sqlite3パッケージには、database/sqlを使う際にドライバとして使用されるプログラムが用意されている。このパッケージにある関数や構造体を直接利用するわけではなく、database/sqlにある機能を使う際に、依存関係にあるgithub.com/mattn/go-sqlite3パッケージの機能が呼び出される。
	"hello"
)

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

	nm := hello.Input("name")
	ml := hello.Input("mail")
	age := hello.Input("age")
	ag, _ := strconv.Atoi(age)
	
	qry := "insert into mydata (name, mail, age) values (?, ?, ?)"
	con.Exec(qry, nm, ml, ag)
	showRecord(con)
}

// Print all record
func showRecord(con *sql.DB){
	qry := "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next(){
		fmt.Println(mydatafmRws(rs).Str())
	}
}

// get Mydata from Rows
func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age )
	if er != nil {
		panic(er)
	}

	return &md
}

func mydatafmRw(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	
	return &md
}