package main

import (
	"reflect"
	"sort"
	"testing"
)

// type testStruct struct {
// 	id   int
// 	name string // 对此字段进行排序
// }

// var tests = []testStruct{
// 	{1, "response"},
// 	{2, "文档管理系统调研"},
// 	{3, "概要设计说明书-文档仓库"},
// 	{4, "文档仓库关联工控资产"},
// 	{5, "概要设计说明书-全局搜索 - 副本"},
// 	{6, "概要设计说明书-全局搜索"},
// 	{7, "图片2"},
// 	{8, "图片1"},
// 	{9, "minio"},
// 	{10, "演示文稿1"},
// 	{11, "组件结构图"},
// 	{12, "安管平台"},
// 	{13, "文档仓库结构"},
// 	{14, "filebrowser"},
// 	{15, "minio-server后台服务"},
// 	{16, "ECS700"},
// 	{17, "mino-server"},
// 	{18, "minio控制台"},
// 	{19, "端口分配_2023-07-27"},
// }

func TestSliceSort(t *testing.T) {
	sort.Slice(tests, func(i, j int) bool {
		return CompareObjectByStr(tests[i].name, tests[j].name)
	})

	// windows排序的结果
	wants := []testStruct{
		{16, "ECS700"},
		{14, "filebrowser"},
		{9, "minio"},
		{15, "minio-server后台服务"},
		{18, "minio控制台"},
		{17, "mino-server"},
		{1, "response"},
		{12, "安管平台"},
		{19, "端口分配_2023-07-27"},
		{5, "概要设计说明书-全局搜索 - 副本"},
		{6, "概要设计说明书-全局搜索"},
		{3, "概要设计说明书-文档仓库"},
		{7, "图片1"},
		{8, "图片2"},
		{4, "文档仓库关联工控资产"},
		{13, "文档仓库结构"},
		{2, "文档管理系统调研"},
		{10, "演示文稿1"},
		{11, "组件结构图"},
	}

	if !reflect.DeepEqual(tests, wants) {
		t.Error("fail")
	}
}
