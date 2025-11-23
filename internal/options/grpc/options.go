package grpc

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Options struct {
	Address string
}

func GetOptions(ctx context.Context) (Options, error) {
	options := Options{Address: ":8003"}
	configData, err := g.Cfg().Get(ctx, "grpc")
	if err != nil {
		g.Log().Errorf(ctx, "查询Grpc配置异常,异常信息:%v", err)
		return options, nil
	}

	err = configData.Scan(&options)

	if err != nil {
		g.Log().Errorf(ctx, "转换Grpc配置异常,配置:%v,异常信息:%v", configData.String(), err)
		return options, nil
	}

	return options, nil
}
