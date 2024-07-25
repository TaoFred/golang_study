package vfcore

import (
	"time"

	"gorm.io/gorm"
)

const (
	SYSTEM_TYPE_ECS700 = "ECS-700"
	SYSTEM_TYPE_TRICON = "Tricon"
	SYSTEM_TYPE_TCS900 = "TCS-900"
)
const (
	// 组态数据
	TABLE_CFG_MAIN = "configuration_main" // 采集表

	// 平台数据
	TABLE_LOG    = "audit_log"      // 操作日志表
	TABLE_MODULE = "module_info"    // 模块表,侧边栏Tab
	TABLE_SYSTEM = "system_info"    // 工程系统表
	TABLE_DEVICE = "equipment_info" // 工程分组表

	TABLE_USER        = "user_info"         // 用户表
	TABLE_ROLE        = "role_info"         // 角色表
	TABLE_ROLE_SYSTEM = "role_system_merge" // 角色-工程关联表
	TABLE_ROLE_MODULE = "role_module_merge" // 角色-模块关联表

	TABLE_ASSET_TREE = "asset_tree" // 组织树表
	// TABLE_ASSET_PROGRAM           = "asset_program"           // 资产-程序表
	TABLE_ASSET_NODE              = "asset_node"              // 资产-节点表
	TABLE_ASSET_NODE_PIN          = "asset_node_pin"          // 资产-节点引脚表
	TABLE_ASSET_NODE_PIN_RELATION = "asset_node_pin_relation" // 资产-节点引脚关系表

	VIEW_MAIN_NEWEST = "v_configuration_main_newest" // 采集表中最新fatherId的映射视图

)

type TableAssetNodePinRelation struct {
	BaseModel
	FatherId     int64 `gorm:"column:father_id;     type:int8;index;                        comment:采集ID"                 json:"FatherId"`
	BeforeNodeId int64 `gorm:"column:before_node_id;type:int8;uniqueIndex:uniq_pin_relation;comment:前置节点Id[冗余字段]"   json:"BeforeNodeId"`
	BeforePinId  int64 `gorm:"column:before_pin_id; type:int8;uniqueIndex:uniq_pin_relation;comment:前置引脚Id"             json:"BeforePinId"`
	AfterNodeId  int64 `gorm:"column:after_node_id; type:int8;uniqueIndex:uniq_pin_relation;comment:后置节点Id[冗余字段]"   json:"AfterNodeId"`
	AfterPinId   int64 `gorm:"column:after_pin_id;  type:int8;uniqueIndex:uniq_pin_relation;comment:后置引脚Id"             json:"AfterPinId"`
	WireType     int32 `gorm:"column:wire_type;     type:int8;not null;                     comment:0：Normal，1：Feedback" json:"WireType"`
}

func (TableAssetNodePinRelation) TableName() string {
	return TABLE_ASSET_NODE_PIN_RELATION
}

type TableAssetNodePin struct {
	BaseModel
	BlockNodeId       int64  `gorm:"column:block_node_id;       type:int8;               not null;           comment:节点Id"                                      json:"BlockNodeId"`
	PinId             string `gorm:"column:pin_id;              type:varchar(256);       not null;           comment:接口返回的引脚Id"                            json:"PinId"`
	Name              string `gorm:"column:name;                type:varchar(256);       not null;           comment:引脚名称"                                    json:"Name"`
	InitDataType      int64  `gorm:"column:init_data_type;      type:int8;               not null;           comment:节点引脚输入或输出的数据类型"                json:"InitDataType"`
	InitDataValue     string `gorm:"column:init_data_value;     type:varchar(256);       not null;           comment:节点引脚初始化的默认值"                      json:"InitDataValue"`
	IsInput           int32  `gorm:"column:is_input;            type:int4;               not null;           comment:节点引脚是否是输入引脚"                      json:"IsInput"`
	IsOutput          int32  `gorm:"column:is_output;           type:int4;               not null;           comment:节点引脚是否是输出引脚"                      json:"IsOutput"`
	IsInvert          int32  `gorm:"column:is_invert;           type:int4;               not null;           comment:是否取反，取反为1，不取反为0"                json:"IsInvert"`
	WireTypeBefore    int32  `gorm:"column:wire_type_before;    type:int4;               not null;           comment:0：Normal，1：Feedback"                      json:"WireTypeBefore"`
	Unfold            int32  `gorm:"column:Unfold;              type:int2;               not null;default:0; comment:引脚是否有连接对象"                          json:"Unfold"`
	UnfoldArr         string `gorm:"column:unfold_arr;          type:varchar(256);       not null;default:0; comment:跟当前引脚连接的节点id列表,原类型为[]string" json:"UnfoldArr"`
	TypicalLoopFBInst int32  `gorm:"column:typical_loop_fb_inst;type:int2;               not null;default:0; comment:引脚前是否是典型回路节点连接线（不允许收起）" json:"TypicalLoopFBInst"`
	FbIdBefore        string `gorm:"column:fb_id_before;        type:varchar(256);       not null;default:'';comment:接口返回的节点引脚前置节点ID"                json:"FbIdBefore"`
	ParaIdBefore      string `gorm:"column:para_id_before;      type:varchar(256);       not null;default:'';comment:接口返回的节点引脚前置节点引脚ID"            json:"ParaIdBefore"`
	UniqueName        string `gorm:"column:unique_name;         type:varchar(512);index;not null;default:'';comment:标识唯一性的名称（不排除不唯一的可能）"      json:"UniqueName"`
	FatherId          int64  `gorm:"column:father_id;           type:int8;         index;                    comment:采集ID[冗余字段]"                            json:"FatherId"`
}

func (TableAssetNodePin) TableName() string {
	return TABLE_ASSET_NODE_PIN
}

type TableAssetNode struct {
	BaseModel
	ProgramId         int64  `gorm:"column:program_id;          type:int8;         index;not null;default:0; comment:归属程序Id, 0表示空"                                      json:"ProgramId"`
	ProgramName       string `gorm:"column:program_name;        type:varchar(128); index;not null;default:'';comment:程序名称[冗余]"                                           json:"ProgramName"`
	BlockHeadName     string `gorm:"column:block_head_name;     type:varchar(256);       not null;           comment:节点名称"                                                 json:"BlockName"`
	BlockFootName     string `gorm:"column:block_foot_name;     type:varchar(256);       not null;default:'';comment:节点类型子项（图谱节点左下角名称）"                       json:"FootName"`
	BlockFootDesc     string `gorm:"column:block_foot_desc;     type:varchar(256);       not null;default:'';comment:节点描述信息"                                             json:"Desc"`
	BlockHeadType     int32  `gorm:"column:block_head_type;     type:int4;               not null ;          comment:节点类型ID"                                               json:"BlockType"`
	BlockHeadTypeName string `gorm:"column:block_head_type_name;type:varchar(256);       not null;           comment:节点类型名称"                                             json:"BlockTypeName"`
	BlockHeadPath     string `gorm:"column:block_head_path;     type:varchar(256);       not null;           comment:[域地址.站地址] 或者 [域地址]"                            json:"BlockPath"`
	BlockBodyContent  string `gorm:"column:block_body_content;  type:varchar(256);       not null;           comment:显示信息"                                                 json:"BlockContent"`
	BlockCombo        string `gorm:"column:block_combo;         type:varchar(256);       not null;default:'';comment:仪表回路专属-节点归属 [现场测 | 辅助柜 | 系统测]"         json:"ComboId"`
	BlockGroup        int32  `gorm:"column:block_group;         type:int4;               not null;default:0; comment:仪表回路专属-节点层级 (Y向)"                              json:"BlockGroup"`
	BlockLevel        int32  `gorm:"column:block_level;         type:int4;               not null;default:0; comment:仪表回路专属-节点层级 (X向)"                              json:"BlockLevel"`
	Xpos              int    `gorm:"column:x_pos;               type:int4;               not null;default:0; comment:X坐标"                                                    json:"Xpos"`
	Ypos              int    `gorm:"column:y_pos;               type:int4;               not null;default:0; comment:Y坐标"                                                    json:"Ypos"`
	BlockPositionX    string `gorm:"column:block_position_x;    type:varchar(256);       not null;default:'';comment:自定义专属-节点层级(无用)"                                json:"positionX"`
	FBInstDire        int32  `gorm:"column:fb_inst_dire;        type:int2;               not null;default:0; comment:节点中引脚的布局：false(默认)-输入左侧,输出右边;true-反之" json:"FBInstDire"`
	TypicalLoopNode   int32  `gorm:"column:typical_loop_node;   type:int2;               not null;default:0; comment:当前节点是否是典型回路的组成节点"                         json:"TypicalLoopNode"`
	BeforeNodeUnfold  int32  `gorm:"column:before_node_unfold;  type:int2;               not null;default:0; comment:当前节点前置是否可以被展开"                               json:"BeforeNodeUnfold"`
	AfterNodeUnfold   int32  `gorm:"column:after_node_unfold;   type:int2;               not null;default:0; comment:当前节点后置是否可以被展开"                               json:"AfterNodeUnfold"`
	InnerId           string `gorm:"column:inner_id;            type:varchar(256);       not null;default:'';comment:内部Id"                                                   json:"InnerId"`
	InnerName         string `gorm:"column:inner_name;          type:varchar(256);       not null;default:'';comment:内部名称"                                                 json:"InnerName"`
	SheetTitle        string `gorm:"column:sheet_title;         type:varchar(256);index;not null;default:'';comment:页面名称"                                                 json:"SheetTitle"`
	UniqueName        string `gorm:"column:unique_name;         type:varchar(256);index;not null;default:'';comment:标识唯一性的名称（不排除不唯一的可能）"                   json:"UniqueName"`
	UserNodeData      string `gorm:"column:user_node_data;      type:text;         not null;comment:用户节点数据[冗余]"                                       json:"UserNodeData"`
	AssetId           int64  `gorm:"column:asset_id;            type:int8;         	     not null;default:0; comment:关联资产Id，自定义节点不存在"                             json:"AssetId"`
	FatherId          int64  `gorm:"column:father_id;           type:int8;         index;not null;default:0; comment:采集ID"                                                   json:"FatherId"`
	SystemId          int64  `gorm:"column:system_id;           type:int8;         index;not null;default:0; comment:系统ID"                                                  json:"SystemId"`
}

func (TableAssetNode) TableName() string {
	return TABLE_ASSET_NODE
}

type TableRoleModule struct {
	RoleId     int64  `gorm:"column:role_id;    type:int8;        uniqueIndex:uniq_role_module;not null;comment:用户角色 0|超级管理员" json:"role_id"`
	ModulePath string `gorm:"column:module_path;type:varchar(128);uniqueIndex:uniq_role_module;not null;comment:页面相对路径"         json:"module_path"`
}

func (TableRoleModule) TableName() string {
	return TABLE_ROLE_MODULE
}

type TableRoleSystem struct {
	RoleId   int64 `gorm:"column:role_id;  type:int8;uniqueIndex:uniq_role_system;not null;comment:用户角色 0|超级管理员" json:"role_id"`
	SystemId int64 `gorm:"column:system_id;type:int8;uniqueIndex:uniq_role_system;not null;comment:系统Id"               json:"system_id"`
}

func (TableRoleSystem) TableName() string {
	return TABLE_ROLE_SYSTEM
}

type TableRole struct {
	RoleId     int64     `gorm:"column:role_id;    type:bigint;   primaryKey;autoIncrement;                   comment:用户角色 0|超级管理员" json:"RoleId"`
	RoleName   string    `gorm:"column:role_name;  type:varchar(128);uniqueIndex;not null;                    comment:角色名称"             json:"RoleName"`
	RoleDesc   string    `gorm:"column:role_desc;  type:varchar(128);not null;default:'';                     comment:角色描述"             json:"RoleDesc"`
	RoleStatus int32     `gorm:"column:role_status;type:int4;        not null;                                comment:角色状态"             json:"RoleStatus"`
	CreateUser string    `gorm:"column:create_user;type:varchar(128);not null;default:'-';                    comment:创建人"               json:"CreateUser"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp; autoCreateTime;default:CURRENT_TIMESTAMP;  comment:记录创建时间"          json:"CreateTime"`
}

func (TableRole) TableName() string {
	return TABLE_ROLE
}

type TableLog struct {
	BaseModel
	Message       string `gorm:"column:message;        type:text;         not null; comment:消息内容"                                   json:"message"`
	UserId        int64  `gorm:"column:user_id;        type:int8;         not null; comment:用户Id"                                     json:"user_id"`
	UserName      string `gorm:"column:user_name;      type:varchar(128); not null; comment:用户名"                                     json:"user_name"`
	MsgType       int64  `gorm:"column:msg_type        type:int8;         not null; comment:消息类型 1|系统信息  2|用户信息"             json:"msg_type"`
	MsgGrade      int64  `gorm:"column:msg_grade       type:int8;         not null; comment:日志等级 0|所有 1|正常 2|一般 3|警告 4|错误" json:"msg_grade"`
	ErrorCode     string `gorm:"column:error_code      type:varchar(64);  not null; comment:错误码  "                                   json:"error_code"`
	RelationGroup string `gorm:"column:relation_group; type:varchar(64);  not null; comment:关系分组 用于标识同一个批次的消息  "         json:"relation_group"`
}

func (TableLog) TableName() string {
	return TABLE_LOG
}

type BaseModel struct {
	Id         int64          `gorm:"column:id;         type:bigint;  primaryKey;autoIncrement;                comment:主键Id"      json:"Id"`
	CreateTime time.Time      `gorm:"column:create_time;type:timestamp;autoCreateTime;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"CreateTime"`
	UpdateTime time.Time      `gorm:"column:update_time;type:timestamp;autoUpdateTime;default:CURRENT_TIMESTAMP;comment:记录更新时间" json:"UpdateTime"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:timestamp;index;                                   comment:记录删除时间" json:"DeleteTime,omitempty"`
}

type TableUser struct {
	BaseModel
	Name       string `gorm:"column:name;       type:varchar(128);not null;default:''; comment:昵称"                          json:"name"`
	Department string `gorm:"column:department; type:varchar(128);not null;default:''; comment:部门"                          json:"department"`
	Position   string `gorm:"column:position;   type:varchar(128);not null;default:''; comment:岗位"                          json:"position"`
	Telephone  string `gorm:"column:telephone;  type:varchar(32); not null;default:''; comment:电话"                          json:"telephone"`
	Role       int64  `gorm:"column:role;       type:int8;        not null;default:0;  comment:用户角色 0|超级管理员"           json:"role"`
	UserName   string `gorm:"column:user_name;  type:varchar(128);uniqueIndex;not null;comment:用户名"                         json:"user_name"`
	Password   string `gorm:"column:password;   type:varchar(128);not null;            comment:用户密码 明文密码经MD5加密后保存" json:"password"`
	Status     int32  `gorm:"column:status;     type:int4;        not null;            comment:账号状态 1|启用 0|禁用"          json:"status"`
	LoginTime  int64  `gorm:"column:login_time; type:int8;                             comment:登录时间"                       json:"login_time"`
	CreateUser string `gorm:"column:create_user;type:varchar(128);not null;            comment:创建人"                         json:"create_user"`
}

func (TableUser) TableName() string {
	return TABLE_USER
}

type TableSystem struct {
	SystemId         int64          `gorm:"column:system_id;                 type:bigint;      primaryKey; autoIncrement;               comment:系统Id"`
	ProjectName      string         `gorm:"column:project_name;              type:varchar(128);   not null; unique;                        comment:工程名"`
	SystemName       string         `gorm:"column:system_name;               type:varchar(128);   not null; unique;                        comment:系统名"`
	ProjectType      string         `gorm:"column:project_type;              type: varchar(128);  not null; default:'DCS',                 comment:工程类型"`
	SystemType       string         `gorm:"column:system_type;               type:varchar(128);   not null; default:'';                    comment:系统类型"`
	ChangeRuleId     int64          `gorm:"column:change_rule_id;            type:int8;           default:0;                               comment:审计规则Id"                                json:"ChangeRuleId"`
	DefectRuleId     int64          `gorm:"column:defect_rule_id;            type:int8;           default:0;                               comment:缺陷规则Id"                                json:"DefectRuleId"`
	IsEdit           int16          `gorm:"column:is_edit;                   type:int2;           default:0;                               comment:是否可以编辑(仅手动添加数据可编辑) 0|否 1|是" json:"IsEdit"`
	LinkDeviceId     int64          `gorm:"column:link_device_id;            type:int8;           default:0;                               comment:本系统关联的设备id"                         json:"LinkDeviceId"`
	CollectBeginTime int64          `gorm:"column:collect_begin_time;        type:int8;                                                    comment:采集开始时间"                               json:"CollBeginTime"`
	CollectInterval  int32          `gorm:"column:collect_interval;          type:int4;                                                    comment:采集间隔(天)"                               json:"CollInterval"`
	IsUsed           int16          `gorm:"column:is_used;                   type:int2;           default:0;                               comment:是否启用 0|否 1|是"                         json:"IsUsed"`
	OpcUAAddr        string         `gorm:"column:opcua_addr;                type:varchar(32);    not null;default:'';                     comment:opcua地址"                                  json:"OpcUAAddr"`
	IsSuccess        int16          `gorm:"column:is_success;                type:int2;           default:0;                               comment:上次采集是否成功 0|否 1|是"                  json:"Success"`
	CollectUser      string         `gorm:"column:collect_user;              type:varchar(128);   not null;                                comment:采集人"                                     json:"CollUser"`
	IsComplete       int16          `gorm:"column:is_complete;               type:int2;           default:1;                               comment:是否结束采集 0|否 1|是"                      json:"Complete"`
	CreateTime       time.Time      `gorm:"column:create_time;               type:timestamp;    autoCreateTime;default:CURRENT_TIMESTAMP;comment:记录创建时间"                                json:"CreateTime"`
	UpdateTime       time.Time      `gorm:"column:update_time;               type:timestamp;    autoUpdateTime;default:CURRENT_TIMESTAMP;comment:记录更新时间"                                json:"UpdateTime"`
	DeleteTime       gorm.DeletedAt `gorm:"column:delete_time;index;         type:timestamp;                                             comment:记录删除时间"                                json:"DeleteTime"`
	ProjectPath      string         `gorm:"column:project_path;              type:varchar(512);   not null;default:'';                     comment:SIS组态文件路径" json:"ProjectPath"`
	RealtimeSwitch   int16          `gorm:"column:realtime_switch;           type:int2;not null;  default:0;                               comment:实时数据服务开关 0|关闭 1|开启" json:"RealtimeSwitch" diff:"ignore"`
}

func (TableSystem) TableName() string {
	return TABLE_SYSTEM
}

type TableAssetTree struct {
	BaseModel
	PreId                     int64  `gorm:"column:pre_id;                     type:int8;         index;                                         comment:父节点Id"`
	FatherId                  int64  `gorm:"column:father_id;                  type:int8;         uniqueIndex:uniq_asset_tree;                   comment:采集ID"`
	SystemId                  int64  `gorm:"column:system_id;                  type:int4;               not null;index;                          comment:系统ID[冗余]"`
	ProjectType               string `gorm:"column:project_type;               type:varchar(128);       not null;default:'';                     comment:工程类型[冗余]"`
	SystemType                string `gorm:"column:system_type;                type:varchar(128);       not null;default:'';                     comment:系统类型[冗余]"`
	UniqueId                  string `gorm:"column:unique_id;                  type:varchar(256);uniqueIndex:uniq_asset_tree;not null;          comment:组织树节点唯一ID，"`
	Description               string `gorm:"column:description;                type:varchar(256);index;not null;default:'';                     comment:组织树节点唯一ID详情说明"`
	NodeId                    string `gorm:"column:node_id;                    type:varchar(256);      not null;default:'';                     comment:组织树层级ID，不同层级可能相同"`
	NodeName                  string `gorm:"column:node_name;                  type:varchar(256);uniqueIndex:uniq_asset_tree;not null;index;    comment:组织树节点名称[Full]"`
	NodeShortName             string `gorm:"column:node_short_name;            type:varchar(256);      not null;default:'';index;               comment:组织树节点名称[Short]"`
	NodeType                  string `gorm:"column:node_type;                  type:varchar(128);       not null;default:'';                     comment:节点类型 枚举类型 主机节点、域变量集合、程序集、位号表、典型回路"`
	PageType                  int    `gorm:"column:page_type;                  type:int4;         uniqueIndex:uniq_asset_tree;not null;default:0;comment:关联前端的Type 枚举类型"`
	IsAssetNode               int32  `gorm:"column:is_asset_node;              type:int2;               not null;default:0;                      comment:是否资产节点 0|否 1|是（用于区分是否虚拟节点）"`
	Level                     int32  `gorm:"column:level;                      type:int4;               not null;                                comment:组织树层级"`
	ConfigurationInternalType int32  `gorm:"column:configuration_internal_type;type:int4;               not null;default:0;                      comment:组态内部类型 枚举类型 功能块位号、模出、模入"`
	ConfigurationInternalDesc string `gorm:"column:configuration_internal_desc;type:varchar(256);      not null;default:'';                     comment:组态内部描述"`
	IsShow                    int16  `gorm:"column:is_show;                    type:int2;                        default:1;                      comment:是否显示 1|是 2|否"`
	IsBackup                  int32  `gorm:"column:is_backup;                  type:int2;               not null;default:0;                      comment:是否备用 0|无备份属性，即无论是否筛选备份都需要显示 1|是备份节点 2|不是备份节点"`
	ExtendParam               string `gorm:"column:extend_param;               type:varchar(256);      not null;default:'';                     comment:拓展参数"`
	Manual                    string `gorm:"column:manual;                     type:varchar(256);      not null;default:'';                     comment:帮助手册信息"`
	AreaAddr                  int32  `gorm:"column:area_addr;                  type:int4;               not null;default:0;                      comment:ECS-700控制域地址"`
	StationAddr               int32  `gorm:"column:station_addr;               type:int4;               not null;default:0;                      comment:ECS-700控制站地址"`
	ChassisId                 int32  `gorm:"column:chassis_id;                 type:int4;               not null;default:0;                      comment:Tricon机架Id"`
	BoardId                   int32  `gorm:"column:board_id;                   type:int4;               not null;default:0;                      comment:Tricon卡件Id"`
	ChannelId                 int32  `gorm:"column:channel_id;                 type:int4;               not null;default:0;                      comment:Tricon通道Id"`
	IsRedundancy              int32  `gorm:"column:is_redundancy;              type:int2;                        default:0;                      comment:是否冗余卡件 0|否 1|是"`
	Sort                      int64  `gorm:"column:sort;                       type:int8;         index;not null;default:0;                      comment:同父节点记录的排序顺序"`
}

func (TableAssetTree) TableName() string {
	return TABLE_ASSET_TREE
}

type TableModule struct {
	BaseModel
	Name    string `gorm:"column:name;    type:varchar(128); not null; default:''; comment:菜单描述"     json:"Name"`
	Path    string `gorm:"column:path;    type:varchar(128);           default:''; comment:页面相对路径" json:"Path"`
	Icon    string `gorm:"column:icon;    type:varchar(128);           default:''; comment:菜单名称"     json:"Icon"`
	Submeun string `gorm:"column:submenu; type:varchar(128);           default:''; comment:子菜单"       json:"Submenu"`
}

func (TableModule) TableName() string {
	return TABLE_MODULE
}

type ViewMainNewest struct {
	ProjectName   string    `gorm:"column:project_name;    type:varchar(128); not null; unique; comment:工程名称"           json:"ProjectName"`
	MaxCreateTime time.Time `gorm:"column:max_create_time; type:timestamp;                      comment:工程的最新更新时间"  json:"MaxCreateTime"`
}

func (ViewMainNewest) TableName() string {
	return VIEW_MAIN_NEWEST
}

type TableCfgMain struct {
	BaseModel
	ProjectName string `gorm:"column:project_name; type:varchar(128); not null; default:'';comment:工程名称"`
	IsComplete  int16  `gorm:"column:is_complete; type:int2; default:0; comment:是否采集成功 0|否 1|是"`
	DisplayName string `gorm:"column:dispaly_name; type:varchar(512); not null; default:''; comment:当前采集系统的DisplayName"`
}

func (TableCfgMain) TableName() string {
	return TABLE_CFG_MAIN
}

type TableDevice struct {
	DeviceId   int64     `gorm:"column:id;          type:bigint;       primaryKey; autoIncrement;                 comment:工程分组Id, 唯一标识" json:"device_id" `
	DeviceName string    `gorm:"column:device_name; type:varchar(128); not null;                                  commnet:工程分组名称"         json:"device_name"`
	CreateTime time.Time `gorm:"column:create_time; type:timestamp;    autoCreateTime; default:CURRENT_TIMESTAMP; comment:记录创建时间"         json:"create_time"`
}

func (TableDevice) TableName() string {
	return TABLE_DEVICE
}
