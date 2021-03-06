package main

import (
	"flag"
	"fmt"
	"hy-im/api/api-gateway/initial"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	defaultConfig := "./config.yaml"
	dmEnv := os.Getenv("DM_MODE")
	if dmEnv != "" {
		defaultConfig = fmt.Sprintf("./config-%s.yaml", dmEnv)
	}

	confPath := flag.String("conf", defaultConfig, "config file path")
	flag.Parse()

	initial.Run(*confPath)
}
