package handler

import (
	"context"

	"github.com/suhasbhagate/GoCode/GoCodeRevision/GRPCDemo/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUserService(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error){
	var user *pb.User
	uid := req.GetUserId()
	if uid <= 0{
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Input")
	}
	for _, v := range UserList{
		if uid == v.UserId{
			user = &pb.User{
				UserId: v.UserId,
				UserName: v.UserName,
			}
		}
	}
	return &pb.GetUserResponse{User: user}, nil
}