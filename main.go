package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type caller func()

func main() {
	noui := flag.Bool("noui", false, "Disable busy UI")
	flag.Parse()
	start := time.Now()
	robotgo.EventHook(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("Ran for :", time.Since(start))
		os.Exit(0)
	})
	s := robotgo.EventStart()
	go func() {
		select {
		case <-robotgo.EventProcess(s):

		}
	}()

	fmt.Println("Press ctrl+shift+q to exit")
	ticker := time.NewTicker(10 * time.Second)

	runner := func(fn caller) {
		for {
			select {
			case <-ticker.C:
				fn()
				robotgo.MoveMouseSmooth(rand.Intn(1000), rand.Intn(1000), 1.0, 1.0)
				robotgo.MoveMouseSmooth(rand.Intn(1000), rand.Intn(1000), 1.0, 1.0)
				robotgo.KeyTap("shift")
			}
		}
	}

	if !*noui {
		a := app.New()
		w := a.NewWindow("Busy")

		busy := widget.NewLabel(fmt.Sprintf("I've been busy for : %s", "0s"))
		busyContainer := container.NewVBox(
			busy,
			widget.NewButton("Reaally? stahp then!", func() {
				os.Exit(0)
			}),
		)
		w.SetContent(busyContainer)

		w.CenterOnScreen()
		go runner(func() {
			w.RequestFocus()
			w.CenterOnScreen()
			busy.Text = fmt.Sprintf("I've been busy for : %s", time.Since(start).Round(5*time.Second).String())
			busyContainer.Refresh()
		})
		w.ShowAndRun()
	} else {
		runner(func() {})
	}

}
