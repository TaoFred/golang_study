package usage

type EamDeviceInfo struct {
	DevId               int64   `json:"id"`                  // 设备ID，保证唯一，且删除后不会再分配一样的
	Name                string  `json:"name"`                // 设备名称
	Remark              string  `json:"remark"`              // 备注
	Code                string  `json:"code"`                // 设备编码
	EamType             string  `json:"eamType"`             // 设备类型
	EamTypeName         string  `json:"eamTypeName"`         // 设备类型名称
	Model               string  `json:"model"`               // 设备型号
	FundAbc             int8    `json:"fundAbc"`             // 资产ABC，ABC资产分级， 0-未分级 1-C 2-B 3-A
	InstallDate         string  `json:"installDate"`         // 安装日期
	InstallPlace        string  `json:"installPlace"`        // 区域位置
	InstallArea         string  `json:"installArea"`         // 安装区域，详细安装位置，一般指一大片区域，比如说A楼，A楼里面又分了很多位置
	AreaPath            string  `json:"areaPath"`            // 区域位置路径, 用"/"分割
	AreaNum             string  `json:"areaNum"`             // 设备位置位号
	Score               float32 `json:"score"`               // 设备健康值，0-100，越高越健康
	Major               string  `json:"major"`               // 所属专业
	MajorName           string  `json:"majorName"`           // 所属专业名称，动静电仪控...
	OperationStatus     string  `json:"operationStatus"`     // 运行状态
	OperationStatusName string  `json:"operationStatusName"` // 运行状态名称
	UseDept             string  `json:"useDept"`             // 使用部门
	UseDeptName         string  `json:"useDeptName"`         // 使用部门名称
	DutyStaff           string  `json:"dutyStaff"`           // 责任人
	DutyStaffName       string  `json:"dutyStaffName"`       // 责任人姓名
	UseYear             float32 `json:"useYear"`             // 使用年限， 单位Year
	UsefulLife          float32 `json:"usefulLife"`          // 已使用年限，单位Year
	ProduceDate         string  `json:"produceDate"`         // 出厂日期
	ProduceCode         string  `json:"produceCode"`         // 出厂编号
	ProduceFirm         string  `json:"produceFirm"`         // 生产厂家
	Vendor              string  `json:"vendor"`              // 供应商
	State               string  `json:"state"`               // 设备状态
	StateName           string  `json:"stateName"`           // 设备状态名称
	EamAssetCode        string  `json:"eamAssetCode"`        // 关联的ERP资产编号
	LogicTag            string  `json:"logicTag"`            // 关联的控制系统位号
	MainCode            string  `json:"mainCode"`            // 主设备编码
}

type EamDevInfoOptionFunc func(*EamDeviceInfo)

func NewEamDeviceInfo(devName string, options ...EamDevInfoOptionFunc) EamDeviceInfo {
	eamDev := EamDeviceInfo{
		Name: devName,
	}
	for _, option := range options {
		option(&eamDev)
	}
	return eamDev
}

func WithDevId(devId int64) EamDevInfoOptionFunc {
	return func(eamDev *EamDeviceInfo) {
		eamDev.DevId = devId
	}
}

func WithRemark(remark string) EamDevInfoOptionFunc {
	return func(eamDev *EamDeviceInfo) {
		eamDev.Remark = remark
	}
}

// ---------------------  接口类型版本  ---------------------------

// 设备信息
// 内部使用
type PrideDeviceInfo struct {
	devId       int64  `json:"id"`          // 设备ID，保证唯一，且删除后不会再分配一样的
	name        string `json:"name"`        // 设备名称
	remark      string `json:"remark"`      // 备注
	code        string `json:"code"`        // 设备编码
	model       string `json:"model"`       // 设备型号
	fundAbc     int8   `json:"fundAbc"`     // 资产ABC，ABC资产分级， 0-未分级 1-C 2-B 3-A
	installDate string `json:"installDate"` // 安装日期
	logicTag    string `json:"logicTag"`    // 关联的控制系统位号
	mainCode    string `json:"mainCode"`    // 主设备编码
}

type PrideDevInfoOption interface {
	apply(*PrideDeviceInfo)
}

type funcPrideDevInfoOption struct {
	f func(*PrideDeviceInfo)
}

func (fdo *funcPrideDevInfoOption) apply(pdi *PrideDeviceInfo) {
	fdo.f(pdi)
}

func newPrideDevInfoOption(f func(*PrideDeviceInfo)) PrideDevInfoOption {
	return &funcPrideDevInfoOption{
		f: f,
	}
}

func WithDevIdPride(devId int64) PrideDevInfoOption {
	return newPrideDevInfoOption(func(pdi *PrideDeviceInfo) {
		pdi.devId = devId
	})
}

func WithRemarkPride(remark string) PrideDevInfoOption {
	return newPrideDevInfoOption(func(pdi *PrideDeviceInfo) {
		pdi.remark = remark
	})
}

func NewPrideDeviceInfo(devName string, options ...PrideDevInfoOption) PrideDeviceInfo {
	pdi := PrideDeviceInfo{
		name: devName,
	}
	for _, option := range options {
		option.apply(&pdi)
	}
	return pdi
}
