package entity

import (
	"errors"
	"fmt"
	"go_gin/config"
	"go_gin/sider/common/stringcore"
	"go_gin/sider/datacenter/vfcore"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(c *config.DB) {
	dial, err := getGormConnectString(c)
	if err != nil {
		panic(err)
	}

	conf := gorm.Config{DisableForeignKeyConstraintWhenMigrating: false}

	engine, err := gorm.Open(dial, &conf)
	if err != nil {
		panic(err)
	}

	sqlDB, err := engine.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(c.MaxConn)

	// 初始化数据模型
	InitDataSourceModel(engine)

	// 默认创建admin用户用于初次登录
	insertInitUserInfo()

	// 默认往module_info表中添加菜单栏数据
	insertModuleInfo()

}

func ClearDataBeforeNew() {
	dropModuleInfo()
}
func dropModuleInfo() {
	db.Exec("DROP TABLE IF EXISTS module_info")
}

func insertModuleInfo() {
	moduleList := getModuleInitData()
	db.Table(vfcore.TABLE_MODULE).Create(&moduleList)
	fmt.Println("create table module_info success")
}

func insertInitUserInfo() {
	isExist := false
	db.Model(&vfcore.TableUser{UserName: "admin"}).Select("true").Take(&isExist)
	if isExist {
		return
	}
	adminInfo := getSuperAdminInitData()
	db.Table(vfcore.TABLE_USER).Create(&adminInfo)
}

func getGormConnectString(c *config.DB) (gorm.Dialector, error) {
	if c == nil {
		return nil, errors.New("please init config")
	}
	var strBuilder strings.Builder
	switch c.Type {
	case DB_TYPE_MYSQL:
		strBuilder.WriteString(c.UserName)
		strBuilder.WriteString(":")
		strBuilder.WriteString(c.Password)
		strBuilder.WriteString(fmt.Sprintf("@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Addr, c.DatabaseName))
		return mysql.New(mysql.Config{DSN: strBuilder.String()}), nil // DSN: Data Source Name 数据源名称；DNS: Domain Name System 域名系统
	case DB_TYPE_PGSQL:
		ipInfo, err := stringcore.ParseIP(c.Addr)
		if err != nil {
			return nil, errors.New("db ip error")
		}
		strBuilder.WriteString(fmt.Sprintf("host=%s ", ipInfo.Host))
		strBuilder.WriteString(fmt.Sprintf("user=%s ", c.UserName))
		strBuilder.WriteString(fmt.Sprintf("password=%s ", c.Password))
		strBuilder.WriteString(fmt.Sprintf("dbname=%s ", c.DatabaseName))
		strBuilder.WriteString(fmt.Sprintf("port=%d", ipInfo.Port))
		return postgres.New(postgres.Config{DSN: strBuilder.String()}), nil
	default:
		return nil, errors.New("unknown database")
	}
}
