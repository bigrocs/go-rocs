package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionByIdLogic {
	return &GetPermissionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPermissionByIdLogic) GetPermissionById(in *pb.GetPermissionByIdReq) (*pb.GetPermissionByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPermissionByIdResp{}, nil
}
