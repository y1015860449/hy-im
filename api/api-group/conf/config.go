package conf

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

type Config struct {
	Release    bool       `yaml:"release"`
	Log        Log        `yaml:"log"`
	Etcd       Etcd       `yaml:"etcd"`
	Kafka      Kafka      `yaml:"kafka"`
	Redis      Redis      `yaml:"redis"`
	Mysql      Mysql      `yaml:"mysql"`
	Token      Token      `yaml:"token"`
	Trace      Trace      `yaml:"trace"`
	Prometheus Prometheus `yaml:"prometheus"`
	Hystrix    Hystrix    `yaml:"hystrix"`
	Micro      Micro      `yaml:"micro"`
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

type Redis struct {
	Addrs        string `yaml:"addrs"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"maxidleconns"`
	MaxOpenConns int    `yaml:"maxopenconns"`
	MaxLifeTime  int    `yaml:"maxlifetime"`
}

type Mysql struct {
	Hostname     []string `yaml:"hoatname"`
	Database     string   `yaml:"database"`
	Username     string   `yaml:"username"`
	Password     string   `yaml:"password"`
	HostPort     string   `yaml:"hostport"`
	RwSeparate   bool     `yaml:"rwseparate"`
	MaxIdleConns int      `yaml:"maxidleconns"`
	MaxOpenConns int      `yaml:"maxopenconns"`
	MaxLifeTime  int      `yaml:"maxlifetime"`
}

type Token struct {
	ExpiredSec int64 `yaml:"expiredsec"`
	SafetySec  int64 `yaml:"safetytsec"`
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
