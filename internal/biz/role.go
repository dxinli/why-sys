package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Role struct {
	ID         int
	Name, Desc string
}

type RoleRepo interface {
	Create(context.Context, *Role) (int, error)
	Delete(context.Context, *Role) error
	Get(context.Context, *Role) (*Role, error)
	List(context.Context, *Role) ([]*Role, error)
	Update(context.Context, *Role) (*Role, error)
	GrantUserRole(ctx context.Context, userId, roleId uint64) error
}

type RoleUsecase struct {
	repo RoleRepo
	log  *log.Helper
}

func NewRoleUsecase(repo RoleRepo, logger log.Logger) *RoleUsecase {
	return &RoleUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *RoleUsecase) CreateRole(ctx context.Context, role *Role) (int, error) {
	return uc.repo.Create(ctx, role)
}

func (uc *RoleUsecase) DeleteRole(ctx context.Context, role *Role) error {
	return uc.repo.Delete(ctx, role)
}

func (uc *RoleUsecase) GetRole(ctx context.Context, role *Role) (*Role, error) {
	return uc.repo.Get(ctx, role)
}

func (uc *RoleUsecase) ListRole(ctx context.Context, role *Role) ([]*Role, error) {
	return uc.repo.List(ctx, role)
}

func (uc *RoleUsecase) UpdateRole(ctx context.Context, role *Role) (*Role, error) {
	return uc.repo.Update(ctx, role)
}

func (uc *RoleUsecase) GrantUserRole(ctx context.Context, userId uint64, roleId uint64) error {
	return uc.repo.GrantUserRole(ctx, userId, roleId)
}
