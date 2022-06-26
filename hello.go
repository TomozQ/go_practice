// Webスクレーピング

package main

import (
	"github.com/PuerkitoBio/goquery"
)

// Webスクレーピング...取得したコンテンツから有用な情報を探して取り出して処理する（Webデータ抽出）
// goquery...HTMLをパース処理し、必要な情報を簡単に取り出せるようにしてくれるパッケージ

// Documentの作成
// HTMLから<<Document>>という構造体を作成

// URLアドレスから作成
// 変数 := goquery.NewDocument( アドレス )

// Responseから作成
// 変数 := goquery.NewDocumentFromResponse( <<Response>> )

// Readerから作成
// 変数 := goquery.NewDocumentFromReader( <<Reader>> )

// Findによる検索処理
// Find...ドキュメントないにあるHTML要素を検索するためのもの
// 変数 := <<Document>>.Find( 検索対象 )
// 引数...検索対象をstringで指定
// 戻り値...「Selection」という構造体

// Selectionからテキストと属性を取得
// テキストを取得
// 変数 := <<Selection>>.Text()

// 要素にある属性の値を取得
// 変数1, 変数2 := <<Selection>>.Attr( 名前 )
// 問題が発生した場合には変数2にerrorが格納される

// 要素のHTMLを取得
// 変数1, 変数2 := <<Selection>>.Html()
// 戻り値...変数1 -> string, 変数2 -> error

// 要素の前後の要素を取得
// 変数 := <<Selection>>.BeforeSelection()
// 変数 := <<Selection>>.AfterSelection()

// 要素の親要素を取得
// 変数 := <<Selection>>.Parent()

// 要素に組み込まれている子要素を取得
// 変数 := <<Selection>>.Children()

// aタグのリンクを出力する
func main(){
	p := "https://golang.org"
	doc, er := goquery.NewDocument(p) // goqueryではhttp.GetのようにCloseでの開放を考える必要がない。

	if er != nil {
		panic(er)
	}

	// Findの引数にcssセレクタでaタグを指定
	doc.Find("a").Each(func(n int, sel *goquery.Selection) {
	// Eachの引数 -> func ( <<int>>, <<Selection>> ){....}
	// 第一引数には、取得したSelectionのindexがint値で渡される。第二引数には取り出したSelectionが渡される。
		
		// href属性の値を取り出す
		lk, _ := sel.Attr("href")
		println(n, sel.Text(), "(", lk, ")")
	})
}