package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func main(){
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	//  入力フォーム
	ne := widget.NewEntry()
	// パスワード入力フォーム
	pe := widget.NewPasswordEntry()

	w.SetContent(
		container.NewVBox(
			l, 
			// フォームを作成
			widget.NewForm(
				// フォーム内に部品を置いていく
				widget.NewFormItem("Name", ne), // 入力フォーム
				widget.NewFormItem("Pass", pe), // パスワード入力フォーム
			),
			// ボタン配置
			widget.NewButton("OK", func(){
				// ne.Text() -> neの入力値
				l.SetText(ne.Text + " & " + pe.Text)
			}),
		),
	)

	w.ShowAndRun()
}

// Formはwebとは認識が違う。
// FyneではあくまでデザインのためにFormの構造体を利用する。