package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// 定义接收配置文件的结构体
type DataBaseConnection struct {
	IpAddress    string
	Port         int
	UserName     string
	Password     int
	DataBaseName string
}

func main() {
	config := viper.New()
	config.AddConfigPath("./config")
	// config.SetConfigName("appConfig")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	// fmt.Println(config.Get("server.http"))

	var configData DataBaseConnection
	configDataMap := make(map[string]interface{})

	// err = config.Unmarshal(&configData)
	// if err != nil {
	// 	panic(fmt.Errorf("read config file to struct err: %s \n", err))
	// }

	err = config.Unmarshal(&configDataMap)
	if err != nil {
		panic(fmt.Errorf("read config file to struct err: %s \n", err))
	}

	fmt.Println(configData)
	fmt.Println(configDataMap)
}

type Test struct {
	Name string
	Age  uint8
}

func (t *Test) GetName() string {
	return t.Name
}

func (t *Test) GetAge() uint8 {
	return t.Age
}

func test(t Test) {
	fmt.Println(t.Age, t.Name)
}
