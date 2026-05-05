package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Ezra Hub client")

	cfg := &config{
		Window: int(os.Stdin.Fd()), // sets reference for term window
	}

	startRepl(cfg)
}
