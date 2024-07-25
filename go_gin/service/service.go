package service

import (
	"fmt"
	"go_gin/config"
	"go_gin/entity"
	"go_gin/sider/common/registry"
)

type service struct {
}

func Run(c *config.InplantConfig) {

	svc := new(service)

	// 初始化数据库
	entity.InitDB(&c.DB)

	registryInfo, _ := registry.ReadRegistry()
	fmt.Printf("========= Integrity DataPath:   %s =========\n", registryInfo.DataPath)
	fmt.Printf("========= Integrity InstallDir: %s =========\n", registryInfo.InstallDir)
	fmt.Printf("========= Integrity Version:    %s =========\n", registryInfo.Version)
	fmt.Printf("========= Integrity DBVersion:  %s =========\n", registryInfo.DatabaseVersion)

	go entity.SetConfig(c.Limits)

	svc.http(c)

}
