package config

import (
	"log"
	"github.com/BurntSushi/toml"
)

type (
	envConfig struct {
		Dev        bool
		LocalDev   bool
		LogConsole bool

		LogDir        string
		SensorData    string
		SensorProject string

		HttpApi string
		WsApi   string

		//redis config
		LoginRedisAddr     string
		LoginRedisPassword string

		CoinRedisAddr     string
		CoinRedisPassword string

		MarryRedisAddr     string
		MarryRedisPassword string

		//mysql config
		MainMysqlDsn string
		LogMysqlDsn  string
		CoinMysqlDsn string
		//supervise

		TcpGoApi          string
		FixRoomApi        string
		ProxyIp           string
		ProxyPort         string
		DirtyFilterServer string
	}
)

var (
	EnvConfigFile = "config/env.toml"

	envConf envConfig
)

func InitEnvConf() {
	if _, err := toml.DecodeFile(EnvConfigFile, &envConf); err != nil {
		log.Fatal(err)
	}
}

func IsDev() bool {
	return envConf.Dev
}

func IsLocalDev() bool {
	return envConf.LocalDev
}

func GetLoginRedisAddr() string {
	return envConf.LoginRedisAddr
}

func GetLoginRedisPassword() string {
	return envConf.LoginRedisPassword
}

func GetMainMysqlDsn() string {
	return envConf.MainMysqlDsn
}

func GetLogDir() string {
	return envConf.LogDir
}
