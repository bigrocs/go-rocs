package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPermissionLogic {
	return &AddPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------permissions-----------------------
func (l *AddPermissionLogic) AddPermission(in *pb.AddPermissionReq) (*pb.AddPermissionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddPermissionResp{}, nil
}
