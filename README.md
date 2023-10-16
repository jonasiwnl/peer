# A TUI frontend for realtime data (logs, packet sniffers, etc)
## written in Go, because I need concurrency (chan <3)

## files
`main.go` - entry point and testing\
`peer.go` - peer struct: handles all TUI

## goals
- [x] show logs
- [x] pause, unpause logs
- [x] add cmg line args for output, data source, etc
- [ ] add ability to select logs and save to .txt
- [ ] data source!

eventually,
- [ ] connect to embedded device
