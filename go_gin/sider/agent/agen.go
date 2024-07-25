package agent

import "github.com/spf13/viper"

func InitConfig(path, name, ctype string) *viper.Viper {
	vip := viper.New()
	vip.AddConfigPath(path)
	vip.SetConfigName(name)
	vip.SetConfigType(ctype)
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	return vip
}
