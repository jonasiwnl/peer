package main

import (
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Peer struct {
	data     chan string
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

func (p *Peer) Run() {
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
	ticker := time.NewTicker(time.Second).C // TODO ticker time should be variable

	for {
		select {
		case e := <-ui.PollEvents():
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
			// TODO We need another event for mouse clicks, switch selected

		case <-ticker:
			if !p.paused {
				// TODO read from channel, update, render
				msg, ok := <-p.data
				if ok {
					para := widgets.NewParagraph()
					para.Text = msg
					para.SetRect(0, 0, 25, 5)
					ui.Render(para)
				}
			}
		}
	}
}
