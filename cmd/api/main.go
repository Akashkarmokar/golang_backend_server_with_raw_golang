package main

import (
	"fmt"

	"github.com/Akashkarmokar/golang_backend_server_with_raw_golang/cmd/internal/config"
)

func main() {
	// Load Configuration

	cfg := config.MustLoad()
	fmt.Println("Config : ", cfg.Env)
}
