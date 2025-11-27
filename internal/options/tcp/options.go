package tcp

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// Options TCP的配置项
type Options struct {
	Address    string //启动地址
	BufferSize int    //每次接受字节大小
}

// GetOptions 获取配置并解析到对象
func GetOptions(ctx context.Context) (Options, error) {
	options := Options{
		Address:    ":8001",
		BufferSize: 1024,
	}

	configData, err := g.Cfg().Get(ctx, "tcp")
	if err != nil {
		g.Log().Errorf(ctx, "查询Grpc配置异常,异常信息:%v", err)
		return options, nil
	}

	err = configData.Scan(&options)

	if err != nil {
		g.Log().Errorf(ctx, "转换Grpc配置异常,配置:%v,异常信息:%v", configData.String(), err)
		return options, nil
	}

	if options.BufferSize < 1024 {
		options.BufferSize = 1024
	}

	return options, nil
}
