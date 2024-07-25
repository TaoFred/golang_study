package main

import (
	"go_gin/config"
	"go_gin/service"
	"go_gin/sider/agent"
)

func main() {
	vip := agent.InitConfig("./conf", "app", "yaml")
	conf := config.Init(*vip)
	service.Run(conf)
}
