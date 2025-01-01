package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {

	//
	// root := tview.NewFlex().SetDirection(tview.FlexRow).
	// 	AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
	// 			AddItem(tview.NewBox().SetBorder(true).SetTitle("Content"), 0, 1, false).
	//    ).AddItem(tview.NewTextView().SetText("help"), 1, 1, false)

	// TODO left to render list
	// TODO right to render details
	// TODO render modal

	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil)
	// AddItem("Quit", "Press to exit", 'q', func() {
	// 	app.Stop()
	// })

	detailsVisible := false
	modalVisible := false

	twoCols := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(list, 0, 2, true).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Right"), 0, 1, false), 0, 1, true).
		AddItem(tview.NewTextView().SetText("help"), 1, 1, false)

	oneCols := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(list, 0, 1, true), 0, 1, true).
		AddItem(tview.NewTextView().SetText("help"), 1, 1, false)

	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Centered Box")

	pages := tview.NewPages().
		AddPage("oneCols", oneCols, true, true).
		AddPage("twoCols", twoCols, true, false).
		AddPage("modal", modal(box, 40, 10), true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 113: // q
			app.Stop()
		case 112: // p
			if detailsVisible == true {
				pages.SwitchToPage("oneCols")
			} else {
				pages.SwitchToPage("twoCols")
			}
			detailsVisible = !detailsVisible
		case 97: // a
			if modalVisible {
				pages.HidePage("modal")
			} else {
				pages.ShowPage("modal")
			}
			modalVisible = !modalVisible
		}

		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
