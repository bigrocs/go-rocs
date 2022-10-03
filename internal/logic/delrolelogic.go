package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRoleLogic {
	return &DelRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelRoleLogic) DelRole(in *pb.DelRoleReq) (*pb.DelRoleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelRoleResp{}, nil
}
