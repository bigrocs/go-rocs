package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSecretKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSecretKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSecretKeyLogic {
	return &DelSecretKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSecretKeyLogic) DelSecretKey(in *pb.DelSecretKeyReq) (*pb.DelSecretKeyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelSecretKeyResp{}, nil
}
