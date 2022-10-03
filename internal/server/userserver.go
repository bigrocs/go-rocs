// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/go-rocs/user-rpc/internal/logic"
	"github.com/go-rocs/user-rpc/internal/svc"
	"github.com/go-rocs/user-rpc/types/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// -----------------------casbinRule-----------------------
func (s *UserServer) AddCasbinRule(ctx context.Context, in *pb.AddCasbinRuleReq) (*pb.AddCasbinRuleResp, error) {
	l := logic.NewAddCasbinRuleLogic(ctx, s.svcCtx)
	return l.AddCasbinRule(in)
}

func (s *UserServer) UpdateCasbinRule(ctx context.Context, in *pb.UpdateCasbinRuleReq) (*pb.UpdateCasbinRuleResp, error) {
	l := logic.NewUpdateCasbinRuleLogic(ctx, s.svcCtx)
	return l.UpdateCasbinRule(in)
}

func (s *UserServer) DelCasbinRule(ctx context.Context, in *pb.DelCasbinRuleReq) (*pb.DelCasbinRuleResp, error) {
	l := logic.NewDelCasbinRuleLogic(ctx, s.svcCtx)
	return l.DelCasbinRule(in)
}

func (s *UserServer) GetCasbinRuleById(ctx context.Context, in *pb.GetCasbinRuleByIdReq) (*pb.GetCasbinRuleByIdResp, error) {
	l := logic.NewGetCasbinRuleByIdLogic(ctx, s.svcCtx)
	return l.GetCasbinRuleById(in)
}

func (s *UserServer) SearchCasbinRule(ctx context.Context, in *pb.SearchCasbinRuleReq) (*pb.SearchCasbinRuleResp, error) {
	l := logic.NewSearchCasbinRuleLogic(ctx, s.svcCtx)
	return l.SearchCasbinRule(in)
}

// -----------------------configs-----------------------
func (s *UserServer) AddConfig(ctx context.Context, in *pb.AddConfigReq) (*pb.AddConfigResp, error) {
	l := logic.NewAddConfigLogic(ctx, s.svcCtx)
	return l.AddConfig(in)
}

func (s *UserServer) UpdateConfig(ctx context.Context, in *pb.UpdateConfigReq) (*pb.UpdateConfigResp, error) {
	l := logic.NewUpdateConfigLogic(ctx, s.svcCtx)
	return l.UpdateConfig(in)
}

func (s *UserServer) DelConfig(ctx context.Context, in *pb.DelConfigReq) (*pb.DelConfigResp, error) {
	l := logic.NewDelConfigLogic(ctx, s.svcCtx)
	return l.DelConfig(in)
}

func (s *UserServer) GetConfigById(ctx context.Context, in *pb.GetConfigByIdReq) (*pb.GetConfigByIdResp, error) {
	l := logic.NewGetConfigByIdLogic(ctx, s.svcCtx)
	return l.GetConfigById(in)
}

func (s *UserServer) SearchConfig(ctx context.Context, in *pb.SearchConfigReq) (*pb.SearchConfigResp, error) {
	l := logic.NewSearchConfigLogic(ctx, s.svcCtx)
	return l.SearchConfig(in)
}

// -----------------------frontPermits-----------------------
func (s *UserServer) AddFrontPermit(ctx context.Context, in *pb.AddFrontPermitReq) (*pb.AddFrontPermitResp, error) {
	l := logic.NewAddFrontPermitLogic(ctx, s.svcCtx)
	return l.AddFrontPermit(in)
}

func (s *UserServer) UpdateFrontPermit(ctx context.Context, in *pb.UpdateFrontPermitReq) (*pb.UpdateFrontPermitResp, error) {
	l := logic.NewUpdateFrontPermitLogic(ctx, s.svcCtx)
	return l.UpdateFrontPermit(in)
}

func (s *UserServer) DelFrontPermit(ctx context.Context, in *pb.DelFrontPermitReq) (*pb.DelFrontPermitResp, error) {
	l := logic.NewDelFrontPermitLogic(ctx, s.svcCtx)
	return l.DelFrontPermit(in)
}

func (s *UserServer) GetFrontPermitById(ctx context.Context, in *pb.GetFrontPermitByIdReq) (*pb.GetFrontPermitByIdResp, error) {
	l := logic.NewGetFrontPermitByIdLogic(ctx, s.svcCtx)
	return l.GetFrontPermitById(in)
}

func (s *UserServer) SearchFrontPermit(ctx context.Context, in *pb.SearchFrontPermitReq) (*pb.SearchFrontPermitResp, error) {
	l := logic.NewSearchFrontPermitLogic(ctx, s.svcCtx)
	return l.SearchFrontPermit(in)
}

// -----------------------logs-----------------------
func (s *UserServer) AddLog(ctx context.Context, in *pb.AddLogReq) (*pb.AddLogResp, error) {
	l := logic.NewAddLogLogic(ctx, s.svcCtx)
	return l.AddLog(in)
}

func (s *UserServer) UpdateLog(ctx context.Context, in *pb.UpdateLogReq) (*pb.UpdateLogResp, error) {
	l := logic.NewUpdateLogLogic(ctx, s.svcCtx)
	return l.UpdateLog(in)
}

func (s *UserServer) DelLog(ctx context.Context, in *pb.DelLogReq) (*pb.DelLogResp, error) {
	l := logic.NewDelLogLogic(ctx, s.svcCtx)
	return l.DelLog(in)
}

func (s *UserServer) GetLogById(ctx context.Context, in *pb.GetLogByIdReq) (*pb.GetLogByIdResp, error) {
	l := logic.NewGetLogByIdLogic(ctx, s.svcCtx)
	return l.GetLogById(in)
}

func (s *UserServer) SearchLog(ctx context.Context, in *pb.SearchLogReq) (*pb.SearchLogResp, error) {
	l := logic.NewSearchLogLogic(ctx, s.svcCtx)
	return l.SearchLog(in)
}

// -----------------------orders-----------------------
func (s *UserServer) AddOrder(ctx context.Context, in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	l := logic.NewAddOrderLogic(ctx, s.svcCtx)
	return l.AddOrder(in)
}

func (s *UserServer) UpdateOrder(ctx context.Context, in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	l := logic.NewUpdateOrderLogic(ctx, s.svcCtx)
	return l.UpdateOrder(in)
}

func (s *UserServer) DelOrder(ctx context.Context, in *pb.DelOrderReq) (*pb.DelOrderResp, error) {
	l := logic.NewDelOrderLogic(ctx, s.svcCtx)
	return l.DelOrder(in)
}

func (s *UserServer) GetOrderById(ctx context.Context, in *pb.GetOrderByIdReq) (*pb.GetOrderByIdResp, error) {
	l := logic.NewGetOrderByIdLogic(ctx, s.svcCtx)
	return l.GetOrderById(in)
}

func (s *UserServer) SearchOrder(ctx context.Context, in *pb.SearchOrderReq) (*pb.SearchOrderResp, error) {
	l := logic.NewSearchOrderLogic(ctx, s.svcCtx)
	return l.SearchOrder(in)
}

// -----------------------permissions-----------------------
func (s *UserServer) AddPermission(ctx context.Context, in *pb.AddPermissionReq) (*pb.AddPermissionResp, error) {
	l := logic.NewAddPermissionLogic(ctx, s.svcCtx)
	return l.AddPermission(in)
}

func (s *UserServer) UpdatePermission(ctx context.Context, in *pb.UpdatePermissionReq) (*pb.UpdatePermissionResp, error) {
	l := logic.NewUpdatePermissionLogic(ctx, s.svcCtx)
	return l.UpdatePermission(in)
}

func (s *UserServer) DelPermission(ctx context.Context, in *pb.DelPermissionReq) (*pb.DelPermissionResp, error) {
	l := logic.NewDelPermissionLogic(ctx, s.svcCtx)
	return l.DelPermission(in)
}

func (s *UserServer) GetPermissionById(ctx context.Context, in *pb.GetPermissionByIdReq) (*pb.GetPermissionByIdResp, error) {
	l := logic.NewGetPermissionByIdLogic(ctx, s.svcCtx)
	return l.GetPermissionById(in)
}

func (s *UserServer) SearchPermission(ctx context.Context, in *pb.SearchPermissionReq) (*pb.SearchPermissionResp, error) {
	l := logic.NewSearchPermissionLogic(ctx, s.svcCtx)
	return l.SearchPermission(in)
}

// -----------------------roles-----------------------
func (s *UserServer) AddRole(ctx context.Context, in *pb.AddRoleReq) (*pb.AddRoleResp, error) {
	l := logic.NewAddRoleLogic(ctx, s.svcCtx)
	return l.AddRole(in)
}

func (s *UserServer) UpdateRole(ctx context.Context, in *pb.UpdateRoleReq) (*pb.UpdateRoleResp, error) {
	l := logic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *UserServer) DelRole(ctx context.Context, in *pb.DelRoleReq) (*pb.DelRoleResp, error) {
	l := logic.NewDelRoleLogic(ctx, s.svcCtx)
	return l.DelRole(in)
}

func (s *UserServer) GetRoleById(ctx context.Context, in *pb.GetRoleByIdReq) (*pb.GetRoleByIdResp, error) {
	l := logic.NewGetRoleByIdLogic(ctx, s.svcCtx)
	return l.GetRoleById(in)
}

func (s *UserServer) SearchRole(ctx context.Context, in *pb.SearchRoleReq) (*pb.SearchRoleResp, error) {
	l := logic.NewSearchRoleLogic(ctx, s.svcCtx)
	return l.SearchRole(in)
}

// -----------------------secretKeys-----------------------
func (s *UserServer) AddSecretKey(ctx context.Context, in *pb.AddSecretKeyReq) (*pb.AddSecretKeyResp, error) {
	l := logic.NewAddSecretKeyLogic(ctx, s.svcCtx)
	return l.AddSecretKey(in)
}

func (s *UserServer) UpdateSecretKey(ctx context.Context, in *pb.UpdateSecretKeyReq) (*pb.UpdateSecretKeyResp, error) {
	l := logic.NewUpdateSecretKeyLogic(ctx, s.svcCtx)
	return l.UpdateSecretKey(in)
}

func (s *UserServer) DelSecretKey(ctx context.Context, in *pb.DelSecretKeyReq) (*pb.DelSecretKeyResp, error) {
	l := logic.NewDelSecretKeyLogic(ctx, s.svcCtx)
	return l.DelSecretKey(in)
}

func (s *UserServer) GetSecretKeyById(ctx context.Context, in *pb.GetSecretKeyByIdReq) (*pb.GetSecretKeyByIdResp, error) {
	l := logic.NewGetSecretKeyByIdLogic(ctx, s.svcCtx)
	return l.GetSecretKeyById(in)
}

func (s *UserServer) SearchSecretKey(ctx context.Context, in *pb.SearchSecretKeyReq) (*pb.SearchSecretKeyResp, error) {
	l := logic.NewSearchSecretKeyLogic(ctx, s.svcCtx)
	return l.SearchSecretKey(in)
}

// -----------------------users-----------------------
func (s *UserServer) AddUser(ctx context.Context, in *pb.AddUserReq) (*pb.AddUserResp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UserServer) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServer) DelUser(ctx context.Context, in *pb.DelUserReq) (*pb.DelUserResp, error) {
	l := logic.NewDelUserLogic(ctx, s.svcCtx)
	return l.DelUser(in)
}

func (s *UserServer) GetUserById(ctx context.Context, in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UserServer) SearchUser(ctx context.Context, in *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	l := logic.NewSearchUserLogic(ctx, s.svcCtx)
	return l.SearchUser(in)
}
