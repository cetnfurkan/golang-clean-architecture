package service

import (
	"context"

	"golang-clean-architecture/repository"
	"golang-clean-architecture/target/ent"
	xgrpc "golang-clean-architecture/target/grpc/user"

	"github.com/cetnfurkan/core/mapper"
)

type (
	UserServiceGrpcImpl struct {
		xgrpc.UnimplementedUserServiceServer
		repository repository.UserRepository
	}
)

// NewUserServiceGrpcImpl creates a new user service gRPC implementation instance.
//
// It takes a user repository instance and returns a new user service gRPC interface instance.
func NewUserServiceGrpcImpl(repository repository.UserRepository) xgrpc.UserServiceServer {
	return &UserServiceGrpcImpl{repository: repository}
}

func (service *UserServiceGrpcImpl) mustEmbedUnimplementedUserServiceServer() {}

func (service *UserServiceGrpcImpl) ListUsers(ctx context.Context, in *xgrpc.ListUsersRequest) (*xgrpc.ListUsersResponse, error) {
	users, err := service.repository.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersDto, err := mapper.ToDto[[]*ent.User, []*xgrpc.User](users)
	if err != nil {
		return nil, err
	}

	return &xgrpc.ListUsersResponse{
		Users: usersDto,
	}, nil
}
