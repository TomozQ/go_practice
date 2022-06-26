package main

import (
	"strconv"
	"fmt"
	"database/sql"
	// importされているパッケージと依存関係にあるパッケージを指定するときには_をつけて記述する。パッケージ内の関数を呼び出して使用するわけではないため、_なしで記述すると構文チェックが実行され使用していないパッケージと判断されてしまう。
	_"github.com/mattn/go-sqlite3" // github.com/mattn/go-sqlite3パッケージには、database/sqlを使う際にドライバとして使用されるプログラムが用意されている。このパッケージにある関数や構造体を直接利用するわけではなく、database/sqlにある機能を使う際に、依存関係にあるgithub.com/mattn/go-sqlite3パッケージの機能が呼び出される。
)

// sql.Open
// 変数1, 変数2 := sql.Open( ドライバ名, データベース名 )
// 引数はどちらもstring
// DBにアクセスし、「DB」構造体を返す（正確にはそのポインタ）

// DBのclose
// defer <<*DB>>.Close() DBへの接続を開放する

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

	// queryをstringで用意
	q := "select * from mydata"
	// 変数1, 変数2 := <<DB>>.Query( sqlクエリ )
	rs, er := con.Query(q) // rsに「Rows」という構造体が返却される
	if er != nil {
		panic(er)
	}

	// Rows
	// カーソルと呼ばれる機能を持っている。カーソル...現在設定されているレコードを示す情報 初期値は取得したレコードの冒頭に設定されている。
	// カーソルが指定されているレコードからデータを取り出す機能を持っている。
	
	// 順番にカーソルを移動させレコードを取得する
	for rs.Next() {	// <<Rows>>.Next()...カーソルを次のレコードに移動する
		var md Mydata
		// 現在のカーソルが当たっているレコードから値を取り出す。
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age) // 引数にはあらかじめ用意しておいた変数のポインタを指定する。レコードから各カラムの値を取り出し、それを引数に用意されている変数に代入する。用意する引数はレコードの値の数だけ揃える必要がある。数が足りないとエラーになる。
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
}