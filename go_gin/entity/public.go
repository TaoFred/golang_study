package entity

import (
	"fmt"
	"go_gin/sider/datacenter/vfcore"
)

func GetSystemInfosBySystemIds(systemIds []int64) (systemInfos []vfcore.TableSystem, err error) {
	err = db.Table(vfcore.TABLE_SYSTEM).Where("id in (?)", systemIds).Find(&systemInfos).Error
	PrintSqlErr(err)
	return
}

func GetLatestFatherIdByProjectName(projectName string) (fatherId int64, err error) {
	err = db.Table(vfcore.TABLE_CFG_MAIN+" cm").Joins("JOIN "+vfcore.VIEW_MAIN_NEWEST+
		" v ON cm.project_name = v.project_name AND cm.create_time = v.max_create_timeAND v.project_name =?",
		projectName).
		Select("cm.id").Take(&fatherId).Error
	PrintSqlErr(err)
	return
}

func GetSystemInfoBySystemId(systemId int64) (systemInfo vfcore.TableSystem, err error) {
	err = db.Table(vfcore.TABLE_SYSTEM).
		Where("system_id = ?", systemId).
		Take(&systemInfo).Error
	PrintSqlErr(err)
	return
}

func PrintSqlErr(err error) {
	if err == nil {
		return
	}
	fmt.Println(err.Error())
}
