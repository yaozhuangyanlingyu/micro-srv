package server

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/yaozhuangyanlingyu/micro-srv/loader"
	"github.com/yaozhuangyanlingyu/micro-srv/logger"
	"github.com/yaozhuangyanlingyu/micro-srv/wrap"
)

// 创建micro.Service用
type AplumServiceConfig struct {
	ServiceName      string `validate:"required"`        // 服务名称
	ServiceAddr      string `validate:"service_address"` // 服务地址
	ConsulAddr       string `validate:"consul_addr"`     // consul地址
	RegisterTTL      int    `validate:"gt=0"`            // 本服务注册到注册中心的存活时间（生命周期）。单位：秒。官方默认值：90s。例：设置为60s，如果注册中心（Consul）在60s内没有收到本服务的健康信息（心跳包），则认为本服务无效。
	RegisterInterval int    `validate:"gt=0"`            // 本服务的向注册中心发送健康信息（心跳包）的时间间隔。单位：秒。官方默认值：30s。例：设置为15s。在18:00:00向注册中心发送一次健康信息，18:00:15再次向注册中心发送健康信息，18:00:30，18:00:45...以此类推。
	LogLevel         string `validate:"required"`        // 日志级别
	WxHost           string `validate:"wx_host"`         // 微信报警服务IP
	WxEmail          string `validate:"wx_email"`        // 微信报警接收人Email
}

func (_this *AplumServiceConfig) SetDefaultValue() {
	const (
		DefaultRegisterTTL      = 90
		DefaultRegisterInterval = 30
	)
	if _this.RegisterTTL == 0 {
		_this.RegisterTTL = DefaultRegisterTTL
	}
	if _this.RegisterInterval == 0 {
		_this.RegisterInterval = DefaultRegisterInterval
	}
}

// 创建micro服务
func NewAplumService(param AplumServiceConfig) micro.Service {
	// 初始化默认参数
	param.SetDefaultValue()

	// 初始化日志组件
	logger.InitLogger(param.LogLevel, param.ServiceName, param.WxHost, param.WxEmail)

	// 创建micro服务
	service := micro.NewService(
		micro.Name(param.ServiceName),          // 服务名称
		micro.Address(param.ServiceAddr),       // 服务地址
		micro.Registry(consulRegistry()),       // consul注册
		micro.WrapHandler(wrap.LogRespWrapper), // 中间件
	)
	return service
}

// consul注册
func consulRegistry() registry.Registry {
	return consul.NewRegistry(
		registry.Addrs(loader.Config.GetString("consul.address")),
	)
}
