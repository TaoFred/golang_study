package entity

import "go_gin/config"

var limitsInfo config.Limits

func GetLimits() config.Limits {
	return limitsInfo
}

func SetConfig(limits config.Limits) {
	setLimits(limits)
}

func setLimits(limits config.Limits) {
	limitsInfo = limits
}
