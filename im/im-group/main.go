package main

import (
	"flag"
	"fmt"
	"hy-im/im/im-group/initial"
	"os"
)

func main() {
	defaultConfig := "./config.yaml"
	dmEnv := os.Getenv("DM_MODE")
	if dmEnv != "" {
		defaultConfig = fmt.Sprintf("./config-%s.yaml", dmEnv)
	}

	confPath := flag.String("conf", defaultConfig, "config file path")
	flag.Parse()

	initial.Run(*confPath)
}
