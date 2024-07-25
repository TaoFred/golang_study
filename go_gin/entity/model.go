package entity

import (
	"go_gin/sider/datacenter/vfcore"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDataSourceModel(engine *gorm.DB) {
	db = engine
	ClearDataBeforeNew()
	db.AutoMigrate(
		new(vfcore.ViewMainNewest),

		new(vfcore.TableCfgMain),

		new(vfcore.TableLog),
		new(vfcore.TableModule),
		new(vfcore.TableSystem),
		new(vfcore.TableDevice),

		new(vfcore.TableUser),
		new(vfcore.TableRole),
		new(vfcore.TableRoleSystem),
		new(vfcore.TableRoleModule),

		new(vfcore.TableAssetTree),
		new(vfcore.TableAssetNode),
		new(vfcore.TableAssetNodePin),
		new(vfcore.TableAssetNodePinRelation),
	)
}
