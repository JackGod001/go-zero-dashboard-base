package model

import (
	"fmt"
	"go_zero_dashboard_base/common/constants"
)

//创建一个方法拼接传进来的sql并返回  // 拼接 data_status
//	query = fmt.Sprintf("%s and `data_status` = 1", query)

func SqlBuildDataStatus(query string, data_status int64) string {
	return fmt.Sprintf("%s and `data_status` = %d", query, data_status)
}

func SqlBuildNormalData(query string) string {
	return SqlBuildDataStatus(query, constants.DataStatusNormal)
}

func SqlBuildDeletedData(query string) string {
	return SqlBuildDataStatus(query, constants.DataStatusDelete)
}
