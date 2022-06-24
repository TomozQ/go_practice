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
	c := widget.NewCheck("Check!", func(f bool){	//チェックが変更された際に実行される変数
		if f {
			l.SetText("Checked!")
		}else {
			l.SetText("Not Checked")
		}
	})
	// チェック状態を変更
	c.SetChecked(true)
	w.SetContent(
		container.NewVBox(
			l, c,
		),
	)

	w.ShowAndRun()
}