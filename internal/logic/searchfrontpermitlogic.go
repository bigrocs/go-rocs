package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFrontPermitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFrontPermitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFrontPermitLogic {
	return &SearchFrontPermitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFrontPermitLogic) SearchFrontPermit(in *pb.SearchFrontPermitReq) (*pb.SearchFrontPermitResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchFrontPermitResp{}, nil
}
