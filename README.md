# A TUI frontend for realtime data (logs, packet sniffers, etc)
## written in Go, because I need concurrency

## files
`main.go` - entry point, frontend, user input, fout\
`data.go` - handle fetching and parsing network data

## goals
- [ ] connect to embedded device
- [ ] show logs
- [ ] pause, unpause logs
- [x] add ability to save to .txt
