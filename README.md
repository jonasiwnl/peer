# A TUI frontend for realtime data (logs, packet sniffers, etc)
## written in Go, because I need concurrency

## files
`main.go` - entry point and testing\
`peer.go` - peer struct: handles all TUI\

## goals
- [ ] show logs
- [x] pause, unpause logs
- [x] add ability to save to .txt
- [ ] add cmg line args for output, data source, etc

eventually,
- [ ] connect to embedded device
