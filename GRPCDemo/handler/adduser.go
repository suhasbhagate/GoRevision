package handler

import (
	"context"

	"github.com/suhasbhagate/GoCode/GoCodeRevision/GRPCDemo/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server)AddUserService(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error){
	uid := req.GetUser().GetUserId()
	unm := req.GetUser().GetUserName()
	if uid <= 0 || len(unm) == 0{
		return &pb.AddUserResponse{Status: false}, status.Errorf(codes.InvalidArgument, "Invalid Input")
	}
	newUser := User{
		UserId: uid,
		UserName: unm,
	}
	UserList = append(UserList, newUser)
	return &pb.AddUserResponse{Status: true}, nil
}

