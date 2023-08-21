# A TUI frontend for realtime data (logs, packet sniffers, etc)
## written in Go, because I need concurrency

## files
`main.rs` - entry point\
`display.rs` - handle tui frontend, displaying new data, and user input\
`data.rs` - handle fetching and parsing network data

## goals
- [ ] connect to embedded device
- [ ] show logs
- [ ] pause, unpause logs
- [ ] add ability to save to .txt
