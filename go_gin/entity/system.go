package entity

import (
	"errors"
	"fmt"
	"go_gin/sider/datacenter/vfcore"
	"strconv"
	"strings"
)

type SystemRequest struct {
	SystemId     int64
	SystemName   string
	ProjectName  string
	SystemType   string
	ProjectType  string
	ChangeRuleId int64
	DefectRuleId int64
	IsEdit       bool
}

// 查询系统返回的数据结构
type SystemResponse struct {
	SystemId       int64    `gorm:"column:id"`
	ProjectName    string   `gorm:"column:project_name"`
	ProjectType    string   `gorm:"column:project_type"`
	SystemName     string   `gorm:"column:system_name"`
	SystemType     string   `gorm:"column:system_type"`
	IsEdit         bool     `gorm:"column:is_edit"`
	LinkDeviceId   int64    `gorm:"column:link_device_id"`
	DefectRuleId   int64    `gorm:"column:defect_rule_id"`
	ChangeRuleId   int64    `gorm:"column:chang_rule_id"`
	DefectRuleName string   `gorm:"-"`
	ChangeRuleName string   `gorm:"-"`
	Role           []string `gorm:"-"`
}

func AddSystem(req *SystemRequest) (systemInfo vfcore.TableSystem, err error) {
	err = checkSystemRequest(req)
	if err != nil {
		return
	}

	if checkIsSystemNumBeyondLimit() {
		return systemInfo, errors.New("系统数量已达上限，请删除后重试")
	}

	if checkIsProjectNameExist(req.ProjectName) {
		return systemInfo, errors.New("工程名称重复，请重新输入")
	}

	if checkIsSystemNameExist(req.SystemName) {
		return systemInfo, errors.New("系统名称重复，请重新输入")
	}

	addSystemInfo := vfcore.TableSystem{
		ProjectName:      req.ProjectName,
		SystemName:       req.SystemName,
		ProjectType:      req.ProjectType,
		SystemType:       req.SystemType,
		ChangeRuleId:     req.ChangeRuleId,
		DefectRuleId:     req.DefectRuleId,
		IsEdit:           1,
		CollectBeginTime: 0,
		CollectInterval:  0,
		IsUsed:           0,
		OpcUAAddr:        "",
		IsSuccess:        1,
		IsComplete:       1,
	}

	err = insertSystemInfo(&addSystemInfo)
	if err != nil {
		return systemInfo, errors.New("创建系统失败")
	}
	return addSystemInfo, nil
}

func ModifySystem(req *SystemRequest) (err error) {
	if req.SystemId <= 0 {
		return fmt.Errorf("系统Id异常, SystemId = %d", req.SystemId)
	}

	err = checkSystemRequest(req)
	if err != nil {
		return
	}

	if checkIsProjectNameExistExceptSystemId(req.ProjectName, req.SystemId) {
		return errors.New("系统名称重复，请重新输入")
	}

	if checkIsSystemtNameExistExceptSystemId(req.SystemName, req.SystemId) {
		return errors.New("系统名称重复，请重新输入")
	}

	err = updateSystemInfo(req.SystemId, map[string]interface{}{
		"project_name":   req.ProjectName,
		"system_name":    req.SystemName,
		"system_type":    req.SystemType,
		"project_type":   req.ProjectType,
		"change_rule_id": req.ChangeRuleId,
		"defect_rule_id": req.DefectRuleId,
	})

	return
}

func DeleteSystem(systemIds string) (err error) {
	idStrs := strings.Split(systemIds, ",")
	var idsInt64 []int64
	for _, idStr := range idStrs {
		idInt64, _ := strconv.ParseInt(idStr, 10, 64)
		if idInt64 <= 0 {
			return errors.New("参数错误")
		}
		idsInt64 = append(idsInt64, idInt64)
	}

	systemInfos, err := GetSystemInfosBySystemIds(idsInt64)
	if err != nil {
		return
	}

	for _, systemInfo := range systemInfos {
		if systemInfo.IsUsed == 1 {
			return fmt.Errorf("系统: (%s) 已启用自动采集，请关闭启用后重试", systemInfo.SystemName)
		}
	}

	err = db.Table(vfcore.TABLE_SYSTEM).Where("id in (?)", idsInt64).Unscoped().Delete(nil).Error
	if err != nil {
		return
	}
	return
}

func GetAllSystem(req *SystemRequest) {

}

func GetSystemByDevId(deviceId int64) (list []SystemResponse, err error) {
	// 对入参进行校验
	if deviceId < 0 {
		return list, errors.New("参数错误")
	}

	// 需要显示设备Id下的系统和未关联设备Id的系统，未关联的设备Id为零
	searchDeviceIds := []int64{0} // 加上未关联设备的系统
	// 查询工程分组表中是否存在该设备Id
	if deviceId != 0 {
		isExist := false
		db.Table(vfcore.TABLE_DEVICE).Select("true").Where("id = ?", deviceId).Take(&isExist)
		if !isExist {
			return list, errors.New("参数错误")
		}
		searchDeviceIds = append(searchDeviceIds, deviceId)
	}

	// 查询系统信息
	err = db.Table(vfcore.TABLE_SYSTEM).Where("link_device_id in (?)", searchDeviceIds).Order("id DESC").Find(&list).Error
	if err != nil {
		PrintSqlErr(err)
	}

	// 查询审计规则和缺陷规则名
	for idx, item := range list {
		err = db.Table(vfcore.TABLE_ROLE_SYSTEM+"AS rs").Distinct("role_name").
			Joins("JOIN "+vfcore.TABLE_ROLE+" ro ON ro.role_id = rs.role_id").
			Where("system_id = ?", item.SystemId).Find(&list[idx].Role).Error
		PrintSqlErr(err)
		if item.ChangeRuleId == 0 {
			list[idx].ChangeRuleName = "默认"
		}

		if item.DefectRuleId == 0 {
			list[idx].DefectRuleName = "默认"
		}
	}
	return
}

func checkIsProjectNameExistExceptSystemId(projectName string, systemId int64) bool {
	isExist := false
	db.Table(vfcore.TABLE_SYSTEM).Select("true").
		Where("project_name = ?", projectName).
		Where("id != ?", systemId).
		Take(&isExist)
	return isExist
}

func checkIsSystemtNameExistExceptSystemId(systemName string, systemId int64) bool {
	isExist := false
	db.Table(vfcore.TABLE_SYSTEM).Select("true").
		Where("system_name = ?", systemName).
		Where("id != ?", systemId).
		Take(&isExist)
	return isExist
}

func insertSystemInfo(systemInfo *vfcore.TableSystem) error {
	err := db.Table(vfcore.TABLE_SYSTEM).Create(systemInfo).Error
	PrintSqlErr(err)
	return err
}

func updateSystemInfo(systemId int64, queryMap map[string]interface{}) error {
	err := db.Table(vfcore.TABLE_SYSTEM).Where("id = ?", systemId).
		Updates(queryMap).Error
	PrintSqlErr(err)
	return err
}

func checkIsSystemNumBeyondLimit() bool {
	isBeyondLimit := false
	maxSystemNumLimit := GetLimits().MaxSystem
	currentSystemNum := int64(0)
	db.Table(vfcore.TABLE_SYSTEM).Count(&currentSystemNum)
	if currentSystemNum >= maxSystemNumLimit {
		isBeyondLimit = true
	}
	return isBeyondLimit
}

func checkIsProjectNameExist(projectName string) bool {
	isExist := false
	db.Table(vfcore.TABLE_SYSTEM).Select("true").Where("project_name = ?", projectName).Take(&isExist)
	return isExist
}

func checkIsSystemNameExist(systemName string) bool {
	isExist := false
	db.Table(vfcore.TABLE_SYSTEM).Select("true").Where("system_name = ?", systemName).Take(&isExist)
	return isExist
}

func checkSystemRequest(req *SystemRequest) error {
	req.ProjectName = strings.Trim(req.ProjectName, " ")
	if req.ProjectName == "" {
		return errors.New("工程名称不允许为空")
	}

	req.SystemName = strings.Trim(req.SystemName, " ")
	if req.SystemName == "" {
		return errors.New("系统名称不允许为空")
	}

	return nil
}
