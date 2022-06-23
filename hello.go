package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main(){
	// 新しいアプリケーションを作成
	a := app.New()

	// ウィンドウを作成
	w := a.NewWindow("Hello")
	// 表示コンテンツの設定
	w.SetContent(
		// Hello Fyne!という内容でラベルを作成
		widget.NewLabel("Hello Fyne!"),
	)

	//ウィンドウを表示し実行
	w.ShowAndRun()
}