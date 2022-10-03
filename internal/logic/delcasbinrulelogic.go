package logic

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCasbinRuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCasbinRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCasbinRuleLogic {
	return &DelCasbinRuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCasbinRuleLogic) DelCasbinRule(in *pb.DelCasbinRuleReq) (*pb.DelCasbinRuleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelCasbinRuleResp{}, nil
}
