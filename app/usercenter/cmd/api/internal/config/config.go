package config

import (
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Salt string

	Mysql struct {
		DataSource string
	}
	//Cache cache.CacheConf
	Cache cache.CacheConf
	Redis redis.RedisConf
	//Casdoor 单点登陆, 由casdoor服务验证用户登录,退出等
	CasdoorConfig casdoorsdk.AuthConfig
	// jwt 登陆 常规登陆,由本体服务验证用户登录,退出等
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
