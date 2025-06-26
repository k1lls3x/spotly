package main

import (
	"spotly/internal/config"
)
func main() {
	config.SetupLogger()
	config.Init()
}
