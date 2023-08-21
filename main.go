package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func save(selected string) {
	// Write to text file
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "PEER"
	p.SetRect(0, 0, 25, 5)

	controls := widgets.NewParagraph()
	controls.Text = "|q| Quit \t |p| Pause \t |s| Save"
	controls.SetRect(0, 6, 33, 7)

	ui.Render(p, controls)

	// State
	paused := false
	selected := ""

	for e := range ui.PollEvents() {
		switch e.ID {
		case "q", "<C-c>": // press 'q' or 'C-c' to quit
			return
		case "p":
			paused = !paused
		case "s":
			save(selected)
		// Maybe?
		case "<Resize>":
			// payload := e.Payload.(ui.Resize)
			// width, height := payload.Width, payload.Height
		}
		// We need another event for mouse clicks, switch selected
	}
}
