package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Permission struct {
	RoleID        uint64
	PermissionIds []uint64
}

type PermissionRepo interface {
	Create(context.Context, uint64, []uint64) error
}

type PermissionUsecase struct {
	repo PermissionRepo
	log  *log.Helper
}

func NewPermissionUsecase(repo PermissionRepo, logger log.Logger) *PermissionUsecase {
	return &PermissionUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (u PermissionUsecase) GrantPermission(ctx context.Context, permission *Permission) error {
	return u.repo.Create(ctx, permission.RoleID, permission.PermissionIds)
}
