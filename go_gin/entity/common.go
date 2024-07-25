package entity

import "go_gin/sider/datacenter/vfcore"

const (
	ACCOUNT_ADMIN  = "admin"
	ACCOUNT_SYSTEM = "system"
)
const (
	DB_TYPE_MYSQL = "mysql"
	DB_TYPE_PGSQL = "postgresql"
)

const (
	HOME         = "主页"
	ASSET        = "工控资产"
	CHANGE       = "基线审计"
	DEFECT       = "组态缺陷"
	SPARES       = "备用资源"
	ARCHITECTURE = "硬件架构"
	JOURNAL      = "操作日志"
	ACCOUNT      = "权限管理"
	CONFIG       = "配置管理"
	METERLOOP    = "仪表回路"
)
const (
	MODULE_ANYONE       = ""
	MODULE_HOME         = "/main/home"
	MODULE_ASSET        = "/main/asset"
	MODULE_CHANGE       = "/main/changes"
	MODULE_DEFECT       = "/main/defects"
	MODULE_SPARES       = "/main/spares"
	MODULE_ARCHITECTURE = "/main/architecture"
	MODULE_JOURNAL      = "/main/journal"
	MODULE_ACCOUNT      = "/main/accountManagement"
	MODULE_CONFIG       = "/main/config"
	MODULE_METERLOOP    = "/main/meterloopConfig"
)

const (
	ICON_HOME         = "i-home"
	ICON_ASSET        = "i-asset"
	ICON_CHANGE       = "i-changes"
	ICON_DEFECT       = "i-defects"
	ICON_SPARES       = "i-spares"
	ICON_ARCHITECTURE = "i-architecture"
	ICON_JOURNAL      = "i-journal"
	ICON_ACCOUNT      = "i-account"
	ICON_CONFIG       = "i-config"
	ICON_METERLOOP    = "i-meterloop"
)

func getModuleInitData() []vfcore.TableModule {
	return []vfcore.TableModule{
		{Name: HOME, Path: MODULE_HOME, Icon: ICON_HOME},
		{Name: ASSET, Path: MODULE_ASSET, Icon: ICON_ASSET},
		{Name: CHANGE, Path: MODULE_CHANGE, Icon: ICON_CHANGE},
		{Name: DEFECT, Path: MODULE_DEFECT, Icon: ICON_DEFECT},
		{Name: SPARES, Path: MODULE_SPARES, Icon: ICON_SPARES},
		{Name: ARCHITECTURE, Path: MODULE_ARCHITECTURE, Icon: ICON_ARCHITECTURE},
		{Name: JOURNAL, Path: MODULE_JOURNAL, Icon: ICON_JOURNAL},
		{Name: ACCOUNT, Path: MODULE_ACCOUNT, Icon: ICON_ACCOUNT},
		{Name: CONFIG, Path: MODULE_CONFIG, Icon: ICON_CONFIG},
		{Name: METERLOOP, Path: MODULE_METERLOOP, Icon: ICON_METERLOOP},
	}
}

func getSuperAdminInitData() vfcore.TableUser {
	return vfcore.TableUser{
		Name:       "Administrator",
		UserName:   ACCOUNT_ADMIN,
		Password:   "21232F297A57A5A743894A0E4A801FC3",
		Status:     1,
		Role:       0,
		CreateUser: ACCOUNT_SYSTEM,
	}
}
