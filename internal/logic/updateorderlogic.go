package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateOrderResp{}, nil
}
