package logic

import (
	"context"
	"fmt"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	user, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	fmt.Print(user, err)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIdResp{User: &pb.User{
		Id:        user.Id,
		Username:  user.Username,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Password:  user.Password,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Origin:    user.Origin,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}}, nil
}
