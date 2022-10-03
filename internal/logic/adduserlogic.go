package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------users-----------------------
func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.AddUserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserResp{}, nil
}
