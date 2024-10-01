package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	v1 "review-b/api/review/v1"
	"review-b/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewServiceDiscovery, NewReviewClient, NewBusinessRepo)

// Data .
type Data struct {
	rc  v1.ReviewClient // grpc 客户端
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		rc:  rc,
		log: log.NewHelper(logger),
	}, cleanup, nil
}

// NewReviewClient 创建grpc客户端
func NewReviewClient(d registry.Discovery) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///reviewService"),
		grpc.WithDiscovery(d),
	)
	if err != nil {
		panic(err)
	}
	return v1.NewReviewClient(conn)
}

// NewServiceDiscovery 服务发现
func NewServiceDiscovery(c *conf.Consul) registry.Discovery {
	consulCfg := api.DefaultConfig()
	consulCfg.Address = c.GetAddress()
	consulCfg.Scheme = c.GetScheme()

	cli, err := api.NewClient(consulCfg)
	if err != nil {
		panic(err)
	}

	return consul.New(cli)
}
