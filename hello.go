// 外部のwebサイトにアクセスする

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

// 変数1, 変数2 := http.Get( アドレス )
// 引数...webサイトのアドレス(URL)をstringで渡す。
// 戻り値...変数1 -> 指定アドレスにGETアクセスし取得した結果, 変数2 -> error
// ※POSTアクセスする場合にはhttp.Post()もある。
// 変数1に返される値はResponse構造体

// ReadCloser
// 変数1, 変数2 := ioutil.ReadAll( <<ReadCloser>> )
// Response構造体のBodyを引数に指定して全テキストをbyte配列で取得する。

// ReadCloserの開放
// <<ReadCloser>>.Close()


// WEBサイトのコンテンツを取得する
func main(){
	p := "https://golang.org"
	// 指定のアドレスにアクセスする
	re, er := http.Get(p)

	if er != nil {
		panic(er)
	}

	defer re.Body.Close()

	// リクエストからコンテンツのテキストを取り出す
	s, er := ioutil.ReadAll(re.Body)

	if er != nil {
		panic(er)
	}

	fmt.Println(string(s))
}