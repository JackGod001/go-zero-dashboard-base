package constants

import "strconv"

const Cdsg = '1'

const (
	//正常
	DataStatusNormal int64 = iota + 1
	//删除
	DataStatusDelete
)

// 获取 DataStatusNormal 的字符串
func GetDataStatusNormalStr() string {
	return strconv.FormatInt(DataStatusNormal, 10) // 将 int64 转换为字符串
}

// 获取 DataStatusDelete 的字符串
func GetDataStatusDeleteStr() string {
	return strconv.FormatInt(DataStatusDelete, 10)
}
