package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchConfigLogic {
	return &SearchConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchConfigLogic) SearchConfig(in *pb.SearchConfigReq) (*pb.SearchConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchConfigResp{}, nil
}
