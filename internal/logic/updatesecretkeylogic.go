package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSecretKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSecretKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSecretKeyLogic {
	return &UpdateSecretKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSecretKeyLogic) UpdateSecretKey(in *pb.UpdateSecretKeyReq) (*pb.UpdateSecretKeyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateSecretKeyResp{}, nil
}
