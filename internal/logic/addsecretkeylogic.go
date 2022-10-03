package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSecretKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSecretKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSecretKeyLogic {
	return &AddSecretKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------secretKeys-----------------------
func (l *AddSecretKeyLogic) AddSecretKey(in *pb.AddSecretKeyReq) (*pb.AddSecretKeyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddSecretKeyResp{}, nil
}
