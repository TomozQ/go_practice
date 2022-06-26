package main

import (
	"bufio"
	"os"
	"fmt"
)

// 変数 := bufio.NewReaderSize( <<Reader>>, int値 )
// NewReaderSize...データを保管するバッファを持ったReaderを作成
// 第一引数...Reader（この場合はFile） 第二引数...バッファサイズをint型で指定

// 改行までデータを読み込む
// 変数1, 変数2, 変数3 := <<Reader>>.ReadLine()
// 戻り値...変数1 -> 読み込んだテキストがbyte配列で返される, 変数2 -> バッファにテキストを格納しきれない場合にはtrue, 変数3 -> error 

// 区切り文字を指定して読み込む
// 変数1, 変数2 := <<Reader>>.ReadString(byte値)
// 戻り値...変数1 -> 読み込んだテキストがstringで返される。変数2 -> error

func main(){
	// 値読み込み関数
	rt := func(f *os.File){ // f *os.FileにOpenFileで開いたファイルの構造体が渡される
		//バッファサイズ4096
		r := bufio.NewReaderSize(f, 4096)
		// 1からスタートし無限ループ
		for i := 1; true; i ++ {
			s, _, er := r.ReadLine() // ReadLineでは最後まで読み込み終わった後で更にReadLineされるとEOF(End Of File)というerrorが発生する。
			// ReadLineのEOFエラーでループを抜ける
			if er != nil {
				break
			}
			// iを増やしてナンバリングして出力
			fmt.Println(i, ":", string(s))
		}
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