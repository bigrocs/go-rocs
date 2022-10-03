package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPermissionLogic {
	return &DelPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPermissionLogic) DelPermission(in *pb.DelPermissionReq) (*pb.DelPermissionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelPermissionResp{}, nil
}
