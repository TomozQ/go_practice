package main

import (
	"io/ioutil"
	"fmt"
)

// 変数1, 変数2 := os.Stat(パス)
// 戻り値...変数1 -> FileInfo, 変数2 -> error

// Fileから直接FileInfoを取り出したい場合
// 変数1, 変数2 := <<Fule>>.Stat()
// 戻り値...取得できた場合 -> 変数1にFileInfo, 取得できなかった場合 -> 変数2にFileInfo

// ディレクトリ内にあるファイルを調べる
// 変数 := ioutil.ReadDir(パス)
// 戻り値...直下にあるファイル/フォルダ類のFileInfo配列が得られる。

func main(){
	// プログラムがある場所のパスを指定
	fs, er := ioutil.ReadDir(".")

	if er != nil {
		panic(er)
	}

	// fsに対してループを回す
	for _, f := range(fs) {
		// ファイル名とbyte数を出力
		fmt.Println(f.Name(), "(", f.Size(), ")")
	}
}