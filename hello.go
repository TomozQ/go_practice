package main

import (
	"os"
	"fmt"
	"hello"
)

// OpenFie
// 変数 := os.OpenFile(ファイルパス, フラグ, <FileMode>)

func main(){
	// 値書き出し関数
	wt := func(f *os.File, s string){ // f *os.FileにOpenFileで開いたファイルの構造体が渡される s strignが書き出すテキスト
		// 引数に指定されたstringをファイルに書きだす。
		_, er := f.WriteString(s + "\n")

		// エラーがあった場合にはエラーを出力してファイルを閉じる
		if er != nil {
			fmt.Println(er)
			f.Close()
			return
		}
	}

	// 書き出すファイル名を指定
	fn := "data.txt"
	// ファイル名にfnを指定
	// O_APPEND...既にある場合には最後に追記する
	// O_CREATE...ファイルがない場合には新たに作成
	// O_WRONLY...書き出し専用モード
	f, er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	// エラーがあった場合にはエラーを出力してファイルを閉じる
	if er != nil {
		fmt.Println(er)
		return
	}

	fmt.Println("***start***")
	// wt関数にOpenFileで開いたファイルの構造体とstringを渡す
	wt(f, "***start***")

	for {
		// helloパッケージのInput関数を実行
		s := hello.Input("type message")
		// 空白でエンターが入力されたら処理を終了する
		if s == "" {
			break
		}
		// 開いているファイルと入力されたテキストを渡す
		wt(f, s)
	}

	// 上のforがbreakされたらendをファイルに書き出し、コンソールに出力する
	wt(f, "*** end ***\n\n")
	fmt.Println("*** end ***")
	// ファイルを閉じる
	er = f.Close()
	// エラーが出た場合にはエラーを出力する。
	if er != nil {
		fmt.Println(er)
	}
}