package main

import (
	"google.golang.org/grpc"
	"murino/murino/user"
)

type GrpcServer struct {
	UserRepository
	user.UnimplementedUserServiceServer
}

func NewRPCServer(repository UserRepository) *grpc.Server {
	srv := GrpcServer{
		UserRepository: repository,
	}
	gsrv := grpc.NewServer()
	user.RegisterUserServiceServer(gsrv, &srv)
	return gsrv
}
