package conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Config struct {
	Log        Log        `yaml:"log"`
	Etcd       Etcd       `yaml:"etcd"`
	Kafka      Kafka      `yaml:"kafka"`
	Release    bool       `yaml:"release"`
	Trace      Trace      `yaml:"trace"`
	Prometheus Prometheus `yaml:"prometheus"`
	Hystrix    Hystrix    `yaml:"hystrix"`
	Micro      Micro      `yaml:"micro"`
	Tcp        Tcp        `yaml:"tcp"`
	Ws         Tcp        `yaml:"ws"`
}

type Micro struct {
	Version          string `yaml:"version"`
	RegisterTTL      int    `yaml:"registerTTL"`
	RegisterInterval int    `yaml:"registerInterval"`
	LimitRate        int    `yaml:"limitRate"`
}

type Prometheus struct {
	Enable bool   `yaml:"enable"`
	Addr   string `yaml:"addr"`
}
type Hystrix struct {
	Enable bool   `yaml:"enable"`
	Addr   string `yaml:"addr"`
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

type Etcd struct {
	Addrs []string `yaml:"addrs"`
}

type Kafka struct {
	Addrs  []string `yaml:"addrs"`
	Enable bool     `yaml:"enable"`
}

type Tcp struct {
	Addr     string `yaml:"addr"`
	IdleTime int    `yaml:"idletime"`
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
