package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"strconv"
)

func main(){
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	e.SetText("0")
	w.SetContent(
		// 縦に一列にラベルを表示する
		container.NewVBox(
		// 横一列にラベルを表示する
		// container.NewHBox(
			l, e,
			widget.NewButton("Click me!", func(){
				// eから入力値を受け取る
				n, _ := strconv.Atoi(e.Text)
				// labelに設定されたテキストを変更する。
				l.SetText("Total: " + strconv.Itoa(total(n)))
			}),
		),
	)

	a.Settings().SetTheme(theme.DarkTheme())
	w.ShowAndRun()
}

func total(n int) int {
	
	t := 0
	for i := 1; i <= n; i++{
		t += i
	}

	return t
}