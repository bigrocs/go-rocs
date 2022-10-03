package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPermissionLogic {
	return &SearchPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPermissionLogic) SearchPermission(in *pb.SearchPermissionReq) (*pb.SearchPermissionResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchPermissionResp{}, nil
}
