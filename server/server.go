package server

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/yaozhuangyanlingyu/micro-srv/loader"
)

// 创建micro.Service用
type AplumServiceConfig struct {
	ServiceName      string `validate:"required"`        // 服务名称
	ServiceAddr      string `validate:"service_address"` // 服务地址
	ConsulAddr       string `validate:"consul_addr"`     // consul地址
	RegisterTTL      int    `validate:"gt=0"`            // 本服务注册到注册中心的存活时间（生命周期）。单位：秒。官方默认值：90s。例：设置为60s，如果注册中心（Consul）在60s内没有收到本服务的健康信息（心跳包），则认为本服务无效。
	RegisterInterval int    `validate:"gt=0"`            // 本服务的向注册中心发送健康信息（心跳包）的时间间隔。单位：秒。官方默认值：30s。例：设置为15s。在18:00:00向注册中心发送一次健康信息，18:00:15再次向注册中心发送健康信息，18:00:30，18:00:45...以此类推。
}

// 创建micro服务
func NewAplumService(param AplumServiceConfig) micro.Service {
	service := micro.NewService(
		//micro.Name(loader.Config.GetString("server.name")),
		//micro.Address(loader.Config.GetString("server.address")),
		micro.Name(param.ServiceName),
		micro.Address(param.ServiceAddr),
		micro.Registry(consulRegistry()),
	)
	return service
}

// consul注册
func consulRegistry() registry.Registry {
	return consul.NewRegistry(
		registry.Addrs(loader.Config.GetString("consul.address")),
	)
}
