package svc

import (
	"bufio"
	"fmt"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/config"
	"go_zero_dashboard_base/app/usercenter/cmd/api/internal/middleware"
	"go_zero_dashboard_base/app/usercenter/model"
	"go_zero_dashboard_base/common/utils"
	"os"
)

//  =============== casdoor单点登陆 ===============

type ServiceContext struct {
	Config               config.Config
	Redis                *redis.Client
	UserModel            model.UserModel
	CasdoorClient        *casdoorsdk.Client
	CasdoorJwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	Certificate, err := readTokenJWTKey(utils.GetConfigPath() + "/token_jwt_key.pem")
	if err != nil {
		fmt.Println("无法读取 token_jwt_key.pem 文件:", err)
		return nil
	}

	// 初始化casdoor
	casdoorClient := casdoorsdk.NewClient(
		c.CasdoorConfig.Endpoint,
		c.CasdoorConfig.ClientId,
		c.CasdoorConfig.ClientSecret,
		Certificate,
		c.CasdoorConfig.OrganizationName,
		c.CasdoorConfig.ApplicationName,
	)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       0,
	})
	CasdoorJwtMiddleware, err := middleware.NewCasdoorJwtMiddleware(casdoorClient)
	if err != nil {
		fmt.Println("无法初始化 jwtHandle:", err)
		return nil
	}
	return &ServiceContext{
		Config:               c,
		Redis:                redisClient,
		UserModel:            model.NewUserModel(mysqlConn, c.Cache),
		CasdoorClient:        casdoorClient,
		CasdoorJwtMiddleware: CasdoorJwtMiddleware.Handle,
	}
}
func readTokenJWTKey(fileFullPath string) (string, error) {
	// 打开文件
	file, err := os.Open(fileFullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建一个Scanner来读取文件
	scanner := bufio.NewScanner(file)

	// 逐行读取文件并将内容拼接成字符串
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	// 检查Scanner是否出现错误
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}
