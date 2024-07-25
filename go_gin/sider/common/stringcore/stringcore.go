package stringcore

import (
	"errors"
	"strconv"
	"strings"
)

type IPInfo struct {
	Host  string
	Port  uint16
	Addrs []uint8
}

// 校验并解析ip地址，仅支持IPv4
func ParseIP(strAddr string) (info IPInfo, err error) {
	addrArr := strings.Split(strAddr, ":")
	switch len(addrArr) {
	case 0, 1:
		return info, errors.New("incorrect ip format")
	case 2:
	default:
		return info, errors.New("IPv6 is not supported")
	}

	hostArr := strings.Split(addrArr[0], ".")
	ipLen := len(hostArr)
	if ipLen != 4 {
		return info, errors.New("incorrect ip format")
	}
	var addrs []uint8
	for _, addr := range hostArr {
		v, ok := judgeV4(addr)
		if !ok {
			return info, errors.New("incorrect ip format, should be in 0.0.0.0~255.255.255.255")
		}
		addrs = append(addrs, v)
	}

	port, ok := judegPort(addrArr[1])
	if !ok {
		return info, errors.New("incorrect port, should be in 1~65535")
	}

	info = IPInfo{
		Host:  addrArr[0],
		Port:  port,
		Addrs: addrs,
	}
	return info, nil
}

func judgeV4(addr string) (uint8, bool) {
	if len(addr) > 3 {
		return 0, false
	}
	num, err := strconv.Atoi(addr)
	if err != nil {
		return 0, false
	}
	if num < 0 || num > 255 {
		return 0, false
	}
	if len(addr) > 1 && addr[0] == '0' {
		return 0, false
	}
	return uint8(num), true
}

func judegPort(port string) (uint16, bool) {
	if len(port) > 5 {
		return 0, false
	}
	num, err := strconv.Atoi(port)
	if err != nil {
		return 0, false
	}
	if num <= 0 || num > 65535 {
		return 0, false
	}
	if len(port) > 1 && port[0] == '0' {
		return 0, false
	}
	return uint16(num), true
}

// 判断IPv4地址是否合法
func isValidIPv4(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}
