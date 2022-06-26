package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

// OpenFie
// 変数 := os.OpenFile(ファイルパス, フラグ, <FileMode>)
// panic( <<error>> ) -> プログラムを強制終了する
// defer ...実行する処理... -> プログラムを終了する際、最後に必ず実行する処理

func main(){
	// 値読み込み関数
	rt := func(f *os.File){ // f *os.FileにOpenFileで開いたファイルの構造体が渡される
		// 読み取り値をsに格納する
		s, er := ioutil.ReadAll(f) // ReadAllで全てをまとめて取り出す。戻り値はbyte配列とerror

		// エラーがあった場合にはエラーを出力してファイルを閉じる
		if er != nil {
			panic(er)
		}

		fmt.Println(string(s))
	}

	// 書き出すファイル名を指定
	fn := "data.txt"
	// ファイル名にfnを指定
	// O_RDONRY...読み出し専用
	f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	// エラーがあった場合にはエラーを出力してファイルを閉じる
	if er != nil {
		panic(er)
	}

	// defer close
	defer f.Close()

	fmt.Println("<<start>>")
	rt(f)
	fmt.Println("<< end >>")
}