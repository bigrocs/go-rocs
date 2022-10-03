package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCasbinRuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCasbinRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCasbinRuleLogic {
	return &UpdateCasbinRuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCasbinRuleLogic) UpdateCasbinRule(in *pb.UpdateCasbinRuleReq) (*pb.UpdateCasbinRuleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateCasbinRuleResp{}, nil
}
