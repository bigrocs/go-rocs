package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSecretKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSecretKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSecretKeyLogic {
	return &SearchSecretKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSecretKeyLogic) SearchSecretKey(in *pb.SearchSecretKeyReq) (*pb.SearchSecretKeyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchSecretKeyResp{}, nil
}
