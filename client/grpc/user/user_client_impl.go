package user

import (
	"context"

	"golang-clean-architecture/config"
	xgrpc "golang-clean-architecture/target/grpc/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	UserClientImpl struct {
		cfg *config.Config
	}
)

func NewUserClientImpl(cfg *config.Config) xgrpc.UserServiceClient {
	return &UserClientImpl{
		cfg: cfg,
	}
}

func (c *UserClientImpl) ListUsers(ctx context.Context, in *xgrpc.ListUsersRequest, opts ...grpc.CallOption) (*xgrpc.ListUsersResponse, error) {
	conn, err := grpc.Dial(c.cfg.UserService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := xgrpc.NewUserServiceClient(conn)
	return client.ListUsers(ctx, in, opts...)
}
