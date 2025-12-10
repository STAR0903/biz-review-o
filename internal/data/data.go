package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "review-o/api/review/v1"
	"review-o/internal/conf"
	"time"

	validate "github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDiscovery, NewReviewClient, NewData, NewOperationRepo)

// Data .
type Data struct {
	rc     v1.ReviewClient
	logger *log.Helper
}

// NewData .
func NewData(rc v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{rc: rc, logger: log.NewHelper(logger)}, cleanup, nil
}

// OperationRepo 是操作仓库接口
func NewDiscovery(r *conf.Registry, logger log.Logger) registry.Discovery {
	c := api.DefaultConfig()
	c.Address = r.Consul.Address
	c.Scheme = r.Consul.Scheme

	client, err := api.NewClient(c)
	if err != nil {
		log.NewHelper(logger).Errorf("api.NewClient err: %v", err)
		panic(err)
	}
	return consul.New(client)
}

// NewReviewClient 返回一个 ReviewClient 实例
func NewReviewClient(conf *conf.Registry, d registry.Discovery, logger log.Logger) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Endpoint),
		grpc.WithDiscovery(d),
		grpc.WithTimeout(3600*time.Second),
		grpc.WithMiddleware(
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
	)
	if err != nil {
		log.NewHelper(logger).Errorf("grpc.DialInsecure err: %v", err)
		panic(err)
	}
	return v1.NewReviewClient(conn)
}
