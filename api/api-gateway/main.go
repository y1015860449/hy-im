package main

import (
	"dm-im/http-gateway/initial"
	"flag"
	"fmt"
	_ "github.com/dm-common/health"
	_ "github.com/dm-common/hystrix"
	_ "github.com/dm-common/pprof"
	_ "github.com/dm-common/prometheus"
	_ "go.uber.org/automaxprocs"
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
