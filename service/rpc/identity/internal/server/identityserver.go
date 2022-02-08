// Code generated by goctl. DO NOT EDIT!
// Source: identity.proto

package server

import (
	"context"

	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/identity"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/logic"
	"github.com/skywalkerwei/pluton-faster/service/rpc/identity/internal/svc"
)

type IdentityServer struct {
	svcCtx *svc.ServiceContext
	identity.UnimplementedIdentityServer
}

func NewIdentityServer(svcCtx *svc.ServiceContext) *IdentityServer {
	return &IdentityServer{
		svcCtx: svcCtx,
	}
}

// 生成token，只针对用户服务开放访问
func (s *IdentityServer) GenerateToken(ctx context.Context, in *identity.GenerateTokenReq) (*identity.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

// 清除token，只针对用户服务开放访问
func (s *IdentityServer) ClearToken(ctx context.Context, in *identity.ClearTokenReq) (*identity.ClearTokenResp, error) {
	l := logic.NewClearTokenLogic(ctx, s.svcCtx)
	return l.ClearToken(in)
}

// validateToken ，只很对用户服务、授权服务api开放
func (s *IdentityServer) ValidateToken(ctx context.Context, in *identity.ValidateTokenReq) (*identity.ValidateTokenResp, error) {
	l := logic.NewValidateTokenLogic(ctx, s.svcCtx)
	return l.ValidateToken(in)
}