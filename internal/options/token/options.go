package token

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// Options token的配置
type Options struct {
	Issuer     string //颁发
	SecretKey  string //密钥
	Expires    int64  //过期时间
	MultiLogin bool   //是否运行多登录
}

// GetOptions 从配置文件中查询Token配置信息
func GetOptions(ctx context.Context) (Options, error) {
	options := Options{
		Issuer:     "crab",
		SecretKey:  "",
		Expires:    118,
		MultiLogin: true,
	}

	configData, err := g.Cfg().Get(ctx, "token")
	if err != nil {
		g.Log().Errorf(ctx, "查询token配置异常,异常信息:%v", err)
		return options, nil
	}

	err = configData.Scan(&options)

	if err != nil {
		g.Log().Errorf(ctx, "转换token配置异常,配置:%v,异常信息:%v", configData.String(), err)
		return options, nil
	}

	return options, nil
}
