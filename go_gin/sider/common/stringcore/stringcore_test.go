package stringcore

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestJudgeV4(t *testing.T) {
	type want struct {
		addr uint8
		flag bool
	}
	type test struct {
		input string
		want
	}

	tests := []test{
		{
			input: "-1",
			want:  want{0, false},
		},
		{
			input: "0",
			want:  want{0, true},
		},
		{
			input: "126",
			want:  want{126, true},
		},
		{
			input: "255",
			want:  want{255, true},
		},
		{
			input: "256",
			want:  want{0, false},
		},
		{
			input: "02",
			want:  want{0, false},
		},
		{
			input: "df",
			want:  want{0, false},
		},
		{
			input: "",
			want:  want{0, false},
		},
	}

	for _, tc := range tests {
		gotUint8, gotBool := judgeV4(tc.input)
		assert.Equal(t, gotUint8, tc.addr)
		assert.Equal(t, gotBool, tc.flag)
	}
}

func TestJudetPort(t *testing.T) {
	type want struct {
		port uint16
		flag bool
	}
	tests := []struct {
		give string
		want
	}{
		{
			give: "0",
			want: want{0, false},
		},
		{
			give: "1",
			want: want{1, true},
		},
		{
			give: "65535",
			want: want{65535, true},
		},
		{
			give: "65536",
			want: want{0, false},
		},
		{
			give: "034",
			want: want{0, false},
		},
		{
			give: "1232442",
			want: want{0, false},
		},
		{
			give: "",
			want: want{0, false},
		},
	}

	for _, tc := range tests {
		gotPort, gotFlag := judegPort(tc.give)
		assert.Equal(t, gotPort, tc.want.port)
		assert.Equal(t, gotFlag, tc.want.flag)
	}
}

func TestParseIp(t *testing.T) {
	var emptyIPInfo IPInfo
	type want struct {
		IPInfo
		err error
	}

	tests := []struct {
		give string
		want
	}{
		{
			give: "0.0.0.-1:1",
			want: want{
				emptyIPInfo,
				errors.New("incorrect ip format, should be in 0.0.0.0~255.255.255.255"),
			},
		},
		{
			give: "0.0.0.01:1",
			want: want{
				emptyIPInfo,
				errors.New("incorrect ip format, should be in 0.0.0.0~255.255.255.255"),
			},
		},
		{
			give: "0.0.0.0:0",
			want: want{
				emptyIPInfo,
				errors.New("incorrect port, should be in 1~65535"),
			},
		},
		{
			give: "0.0.0.0:1",
			want: want{
				IPInfo{
					Host:  "0.0.0.0",
					Port:  1,
					Addrs: []uint8{0, 0, 0, 0},
				},
				nil,
			},
		},
		{
			give: "127.127.127.127:32768",
			want: want{
				IPInfo{
					Host:  "127.127.127.127",
					Port:  32768,
					Addrs: []uint8{127, 127, 127, 127},
				},
				nil,
			},
		},
		{
			give: "255.255.255.255:65535",
			want: want{
				IPInfo{
					Host:  "255.255.255.255",
					Port:  65535,
					Addrs: []uint8{255, 255, 255, 255},
				},
				nil,
			},
		},
		{
			give: "255.255.255.256:65535",
			want: want{
				emptyIPInfo,
				errors.New("incorrect ip format, should be in 0.0.0.0~255.255.255.255"),
			},
		},
		{
			give: "255.255.255.255:65536",
			want: want{
				emptyIPInfo,
				errors.New("incorrect port, should be in 1~65535"),
			},
		},
		{
			give: "255.255.255:65535",
			want: want{
				emptyIPInfo,
				errors.New("incorrect ip format"),
			},
		},
	}
	for _, tc := range tests {
		ipInfo, err := ParseIP(tc.give)
		assert.Equal(t, ipInfo, tc.want.IPInfo)
		assert.Equal(t, err, tc.want.err)
	}
}

func TestIsValidateIpv4(t *testing.T) {

	tests := []struct {
		give string
		want bool
	}{
		{
			give: "127.0.0.1",
			want: true,
		},
		{
			give: "0.0.0.0",
			want: true,
		},
		{
			give: "255.255.255.255",
			want: true,
		},
		{
			give: "-1.0.0.1",
			want: false,
		},
		{
			give: "127.256.0.1",
			want: false,
		},
		{
			give: "127.127.127.127",
			want: true,
		},
		{
			give: "127.03.0.1",
			want: false,
		},
		{
			give: "-0.0.0.1",
			want: false,
		},
		{
			give: "03.0.0.1",
			want: false,
		},
	}

	for _, tc := range tests {
		get := isValidIPv4(tc.give)
		assert.Equal(t, get, tc.want)
	}
}
