package server

import (
	"golang-clean-architecture/repository"
	"golang-clean-architecture/service"
	"golang-clean-architecture/target/ent"
	xuser "golang-clean-architecture/target/grpc/user"

	"github.com/cetnfurkan/core/cache"
	"github.com/cetnfurkan/core/config"
	"github.com/cetnfurkan/core/server"
	"google.golang.org/grpc"
)

// NewGRPCServer creates a new gRPC server instance.
//
// It takes a config instance, a database instance and a cache instance
// and returns a new server interface instance.
//
// It will panic if it fails to create a new gRPC server instance.
func NewGRPCServer(cfg *config.Server, db *ent.Client, cache cache.Cache, grpcEntryName string, boot []byte) server.Server {
	return server.NewGRPCServer(
		cfg,
		grpcEntryName,
		boot,
		server.WithRegisterGrpcFunc(registerGrpcFunctions(db)),
		server.WithRegisterGrpcGatewayFunc(xuser.RegisterUserServiceHandlerFromEndpoint),
	)
}

func registerGrpcFunctions(db *ent.Client) func(*grpc.Server) {
	return func(s *grpc.Server) {
		userRepository := repository.NewUserPostgresRepository(db)
		xuser.RegisterUserServiceServer(s, service.NewUserServiceGrpcImpl(userRepository))
	}
}
