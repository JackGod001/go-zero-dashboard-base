#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package


#生成的表名
#tables=user,user_auth,global_config,platform_app_order,platform_app_statistics,platform_app_withdraw_record,app_platform_address,applications
# todo 生成的表model
tables=*
#表生成的genmodel目录
modeldir=./genModel

# todo 数据库配置
host=127.0.0.1
port=3306
dbname=go_zero_dashboard_base
username=go_zero_dashboard_base
passwd=BGjfAJKxWbZ4X7RP

currentPath=$(pwd)
cd ../../goctl/1.6.2
# 设置goctl模板路径
GOCTL_TEMPLATE_DIR=$(pwd)
echo "goctl 模板目录: "$GOCTL_TEMPLATE_DIR

cd  $currentPath

#echo "开始创建库：$dbname 的表：$2"
$GOCTL_TEMPLATE_DIR/goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero --home=$GOCTL_TEMPLATE_DIR
