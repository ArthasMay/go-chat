package config

import (
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"
	"sync"
)

var once sync.Once
var realPath string
var Conf *Config

type Config struct {
	 Connect ConnectConfig
}

func init() {
	Init()
}

func Init() {
	once.Do(func() {
		env := GetMode()
		realPath := getCurrentDir()
		configFilePath := realPath + "/" + env + "/"
		viper.SetConfigType("toml")
		viper.SetConfigName("/connect")
		viper.AddConfigPath(configFilePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		viper.Unmarshal(&Conf.Connect)
	})
}

func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-1], "/")
	return dir
}

func GetMode() string {
	env := os.Getenv("RUN_MODE")
	if env == "" {
		env = "dev"
	}
	return env
}

func GetGinMode() string {
	env := GetMode()
	if env == "dev" {
		return "debug"
	}
	if env == "test" {
		return "debug"
	}
	if env == "prod" {
		return  "release"
	}
	return "release"
}

type ConnectBase struct {
	CerPath string `mapstructure:"certPath"`
	KeyPath string `mapstructure:"keyPath"`
}

type ConnectRpcAddressWebsockets struct {
	Address string `mapstructure:"address"`
}

type ConnectRpcAddressTcp struct {
	Address string `mapstructure:"address"`
}

type ConnectWebsocket struct {
	ServerId string `mapstructure:"serverId"`
	Bind     string `mapstructure:"bind"`
}

type ConnectTcp struct {
	ServerId      string `mapstructure:"serverId"`
	Bind          string `mapstructure:"bind"`
	SendBuf       int    `mapstructure:"sendBuf"`
	ReceiveBuf    int    `mapstructure:"receivedBuf"`
	KeepAlive     bool   `mapstructure:"keepalive"`
	Reader        int    `mapstructure:"reader"`
	ReadBuf       int    `mapstructure:"readBuf"`
	ReadBufSize   int    `mapstructure:"readBufSize"`
	Writer        int    `mapstructure:"writer"`
	WriterBuf     int    `mapstructure:"writerBuf"`
	WriterBufSize int    `mapstructure:"writeBufSize"`
}

type ConnectBucket struct {
	CpuNum        int    `mapstructure:"cpuNum"`
	Channel       int    `mapstructure:"channel"`
	Room          int    `mapstructure:"room"`
	SrvProto      int    `mapstructure:"svrProto"`
	RoutineAmount uint64 `mapstructure:"routineAmount"`
	RoutineSize   int    `mapstructure:"routineSize"`
}

type ConnectConfig struct {
	ConnectBase                ConnectBase                  `mapstructure:"connect-base"`
	ConnectRpcAddressWebsockets ConnectRpcAddressWebsockets `mapstructure:"connect-rpcAddress-websockets"`
	ConnectRpcAddressTcp       ConnectRpcAddressTcp         `mapstructure:"connect-rpcAddress-tcp"`
	ConnectBucket              ConnectBucket                `mapstructure:"connect-bucket"`
	ConnectWebsocket           ConnectWebsocket             `mapstructure:"connect-websocket"`
	ConnectTcp                 ConnectTcp                   `mapstructure:"connect-tcp"`
}




