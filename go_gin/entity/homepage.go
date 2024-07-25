package entity

import "go_gin/sider/datacenter/vfcore"

// AddrInfo 域&站地址
type AddrInfo struct {
	AreaAddr    int32
	StationAddr int32
}

func GetAssetsNums(systemId int64, addrInfo AddrInfo) (rspInfo []HardwareProperty, err error) {
	systemInfo, err := GetSystemInfoBySystemId(systemId)
	if err != nil {
		return
	}

	fatherId, err := GetLatestFatherIdByProjectName(systemInfo.ProjectName)
	if err != nil {
		return
	}

	switch systemInfo.SystemType {
	case vfcore.SYSTEM_TYPE_ECS700:
		typMap := GetVFAssetsNums(fatherId, addrInfo)
		fbNum := GetNumsByType(fatherId, vfcore.TYPE_FUCTION_BLOCK_IMPL, addrInfo)
		rspInfo = []HardwareProperty{
			{"AI", typMap[vfcore.TYPE_AITAG_NODE]},
			{"AO", typMap[vfcore.TYPE_AOTAG_NODE]},
			{"DI", typMap[vfcore.TYPE_DITAG_NODE]},
			{"DO", typMap[vfcore.TYPE_DOTAG_NODE]},
			{"NA", typMap[vfcore.TYPE_NATAG_NODE]},
			{"ND", typMap[vfcore.TYPE_NDTAG_NODE]},
			{"NN", typMap[vfcore.TYPE_NNTAG_NODE]},
			{"FB", fbNum},
		}
	case vfcore.SYSTEM_TYPE_TCS900:
		typeMap := GetSISAssetsNums(fatherId)
		rspInfo = []HardwareProperty{
			{"I/O变量", typeMap[vfcore.TYPE_SIS_HARDWARE_TAG_IMPL]},
			{"内存变量", typeMap[vfcore.TYPE_SIS_MEMORY_TAG_IMPL]},
			{"操作变量", typeMap[vfcore.TYPE_SIS_COM_TAG_IMPL]},
			{"同步变量", typeMap[vfcore.TYPE_SIS_SYNC_TAG_IMPL]},
			{"通信变量", typeMap[vfcore.TYPE_SIS_COMMU_TAG_IMPL]},
			{"扩展站间通信变量", typeMap[vfcore.TYPE_SIS_ESSTATION_TAG_IMPL]},
			{"扩展通信接口变量", typeMap[vfcore.TYPE_SIS_ESCOMMU_TAG_IMPL]},
			{"FB", typeMap[vfcore.TYPE_SIS_FUCTION_BLOCK_IMPL]},
		}
	case vfcore.SYSTEM_TYPE_TRICON:
		typeMap := GetSISAssetsNums(fatherId)
		rspInfo = []HardwareProperty{
			{"MEMORY", typeMap[vfcore.TYPE_SIS_MEMORY_TAG_IMPL]},
			{"INPUT", typeMap[vfcore.TYPE_SIS_INPUT_TAG_IMPL]},
			{"OUTPUT", typeMap[vfcore.TYPE_SIS_OUTPUT_TAG_IMPL]},
			{"FB", typeMap[vfcore.TYPE_SIS_FUCTION_BLOCK_IMPL]},
		}
	}
	return
}

func GetVFAssetsNums(fatherId int64, addrInfo AddrInfo) map[int]int {
	var assetNums []struct {
		PageType int
		Count    int
	}

	tx := db.Table(vfcore.TABLE_ASSET_TREE+" a").
		Select("a.page_type, COUNT(*) count").
		Joins("LEFT JOIN "+vfcore.TABLE_ASSET_TREE+" b ON a.id = b.pre_id AND a.father_id = b.father_id").
		Where("a.page_type BETWEEN ? AND ?", vfcore.TYPE_AITAG_NODE, vfcore.TYPE_NNTAG_NODE).
		Where("b.father_id = ?", fatherId)

	if addrInfo.AreaAddr >= 0 {
		if addrInfo.StationAddr >= 0 {
			tx = tx.Where("a.area_addr = ? AND a.station_addr = ?", addrInfo.AreaAddr, addrInfo.StationAddr)
		} else {
			tx = tx.Where("a.area_addr = ?", addrInfo.AreaAddr)
		}
	}

	err := tx.Group("a.page_type").Find(&assetNums).Error
	PrintSqlErr(err)

	typeMap := make(map[int]int)
	for _, v := range assetNums {
		typeMap[v.PageType] = v.Count
	}
	return typeMap
}

func GetSISAssetsNums(fatherId int64) map[int]int {
	var assetNums []struct {
		PageType int
		Count    int
	}

	err := db.Table(vfcore.TABLE_ASSET_TREE).
		Select("page_type, COUNT(*) AS count").
		Where("father_id = ?", fatherId).
		Group("page_type").Find(&assetNums).Error
	PrintSqlErr(err)

	typeMap := make(map[int]int)
	for _, v := range assetNums {
		typeMap[v.PageType] = v.Count
	}
	return typeMap
}

func GetNumsByType(fatherId, pageType int64, addrInfo AddrInfo) (nums int64) {
	tx := db.Table(vfcore.TABLE_ASSET_TREE).Where("father_id = ? AND page_type = ?", fatherId, pageType)
	if addrInfo.AreaAddr >= 0 {
		if addrInfo.StationAddr >= 0 {
			tx = tx.Where("area_addr = ? AND station_addr = ?", addrInfo.AreaAddr, addrInfo.StationAddr)
		} else {
			tx = tx.Where("station_addr = ?", addrInfo.StationAddr)
		}
	}
	err := tx.Count(&nums).Error
	PrintSqlErr(err)
	return
}
