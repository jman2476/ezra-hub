package main

import (
	"fmt"
	"os"
	"time"

	apicaller "github.com/jman2476/ezra-hub/app/client/internal/api-caller"
)

func main() {
	fmt.Println("Starting Ezra Hub client")
	ezraClient := apicaller.NewClient(30 * time.Second)

	cfg := &config{
		Window: int(os.Stdin.Fd()), // sets reference for term window
		Client: ezraClient,
	}

	startRepl(cfg)
}
