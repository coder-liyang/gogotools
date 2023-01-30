package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GoGoTools")

	clock := widget.NewLabel(time.Now().Format("Time: 15:04:05"))
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	clock.SetText(time.Now().Format("Time: 15:04:05"))
}
