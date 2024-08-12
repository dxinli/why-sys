package service

import (
	"github.com/google/wire"
	pb "why-sys/api/auth/v1"
	"why-sys/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthService)

type AuthService struct {
	pb.UnimplementedAuthServer
	pb.UnimplementedRoleServer
	pb.UnimplementedMenuServer
	pb.UnimplementedPermissionServer

	uc *biz.AuthUsecase
	rc *biz.RoleUsecase
	mc *biz.MenuUsecase
	pc *biz.PermissionUsecase
}

func NewAuthService(uc *biz.AuthUsecase, rc *biz.RoleUsecase, mc *biz.MenuUsecase, pc *biz.PermissionUsecase) *AuthService {
	return &AuthService{uc: uc, rc: rc, mc: mc, pc: pc}
}
