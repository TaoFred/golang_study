package config

import "github.com/spf13/viper"

type DB struct {
	Addr         string
	Type         string
	DatabaseName string
	UserName     string
	Password     string
	MaxConn      int
}

type Limits struct {
	MaxQuery  int64
	MaxExport int64
	MaxSystem int64
}
type InplantConfig struct {
	DB
	APIPort       string
	WebSocketPort string
	WebPagePort   string
	Limits
}

var Config *InplantConfig

func Init(vip viper.Viper) *InplantConfig {
	Config = &InplantConfig{
		DB: DB{
			Addr:         vip.GetString("db_addr"),
			Type:         vip.GetString("db_driver"),
			DatabaseName: vip.GetString("db_name"),
			UserName:     vip.GetString("db_user"),
			Password:     vip.GetString("db_pwd"),
			MaxConn:      vip.GetInt("db_max_comm"),
		},
		APIPort:       vip.GetString("apiport"),
		WebSocketPort: vip.GetString("websocket_port"),
		WebPagePort:   vip.GetString("webpage_port"),
		Limits: Limits{
			MaxQuery:  vip.GetInt64("max_query"),
			MaxExport: vip.GetInt64("max_export"),
			MaxSystem: vip.GetInt64("max_system"),
		},
	}
	return Config
}
