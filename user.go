package main

import (
	"context"

	"gitlab.com/my0sot1s/pguser/pb"
)

func (u *User) ListUsers(ctx context.Context, in *pb.UserRequest) (*pb.Users, error) {
	return nil, nil
}
func (u *User) CreateUserWithPhone(ctx context.Context, in *pb.User) (*pb.User, error) {
	return nil, nil
}
