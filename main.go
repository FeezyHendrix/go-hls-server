package main

import (
	"fmt"

	"github.com/feezyhendrix/go-hls-server/cmd/api"
)

func main() {
	fmt.Println("Starting HLS Server")
	api.Run()
}
