package main

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Peer struct {
	output   *os.File
	paused   bool
	selected string
}

func (p *Peer) save() {
	_, err := p.output.WriteString(p.selected + "\n")

	if err != nil {
		log.Fatal(err)
	}
}

func (p *Peer) run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	para := widgets.NewParagraph()
	para.Text = "PEER"
	para.SetRect(0, 0, 25, 5)

	controls := widgets.NewParagraph()
	controls.Text = "|q| Quit \t |p| Pause \t |s| Save"
	controls.SetRect(0, 6, 33, 7)

	ui.Render(para, controls)

	for e := range ui.PollEvents() {
		switch e.ID {
		case "q", "<C-c>": // press 'q' or 'C-c' to quit
			return
		case "p":
			p.paused = !p.paused
		case "s":
			p.save()

		// Maybe?
		case "<Resize>":
			// payload := e.Payload.(ui.Resize)
			// width, height := payload.Width, payload.Height
		}
		// We need another event for mouse clicks, switch selected
	}
}

func main() {
	output, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	// close output file on exit and check for its returned error
	defer func() {
		if err := output.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	p := Peer{output, false, ""}
	p.run()

	log.Println("Exiting...")
}
