package conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Config struct {
	Log     Log   `yaml:"log"`
	Http    Http  `yaml:"http"`
	Etcd    Etcd  `yaml:"etcd"`
	Release bool  `yaml:"release"`
	Trace   Trace `yaml:"trace"`
	Micro   Micro `yaml:"micro"`
}

type Micro struct {
	Version          string `yaml:"version"`
	RegisterTTL      int    `yaml:"registerTTL"`
	RegisterInterval int    `yaml:"registerInterval"`
}

type Trace struct {
	Addr string `yaml:"addr"`
}

type Log struct {
	Level        string `yaml:"level"`
	File         string `yaml:"file"`         // 日志文件路径
	RotationTime int    `yaml:"rotationTime"` // 日志文件切割时间(小时)
	MaxAge       int    `yaml:"maxAge"`       // 日志文件保留时长(天)
}

type Http struct {
	Addr string `yaml:"addr"`
}

type Etcd struct {
	Addrs []string `yaml:"addrs"`
}

func NewConfig(configFile string) (*Config, error) {
	fileSource := file.NewSource(
		file.WithPath(configFile),
	)
	conf, _ := config.NewConfig()

	if err := conf.Load(fileSource); err != nil {
		return nil, err
	}

	var c Config
	if err := conf.Scan(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
