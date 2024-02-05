package main

import (
	"context"
	"google.golang.org/grpc"
	"murino/murino/user"
	"strconv"
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

func (g *GrpcServer) LoginUser(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	getUser, _ := g.UserRepository.CheckUser(ctx, req.Email)
	userId := strconv.Itoa(int(getUser.ID))
	token, _ := generateJWT(userId)
	tokenAddr := *token
	return &user.LoginResponse{Status: true, Message: "Successfully", Data: {AccessToken: tokenAddr}}, nil
}
