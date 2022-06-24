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
	r := widget.NewRadioGroup(
		[]string{"One", "Two", "Three"},
		func(s string){
			if s == "" {
				l.SetText("not selected")
			}else{
				l.SetText("selected: " + s)
			}
		},
	)
	// チェック状態を変更
	r.SetSelected("One")
	w.SetContent(
		container.NewVBox(
			l, r,
		),
	)

	w.ShowAndRun()
}