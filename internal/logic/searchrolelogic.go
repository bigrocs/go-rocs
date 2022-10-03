package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRoleLogic {
	return &SearchRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchRoleLogic) SearchRole(in *pb.SearchRoleReq) (*pb.SearchRoleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchRoleResp{}, nil
}
