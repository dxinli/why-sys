package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"why-sys/internal/biz"
	"why-sys/internal/data/ent/menu"
)

type permissionRepo struct {
	data *Data
	log  *log.Helper
}

func NewPermissionRepo(data *Data, logger log.Logger) biz.PermissionRepo {
	return &permissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p permissionRepo) Create(ctx context.Context, roleId uint64, permissionIds []uint64) error {
	// 校验角色
	r, err := p.data.db.Role.Get(ctx, int(roleId))
	if err != nil {
		return err
	}
	// 校验权限
	ids := make([]int, len(permissionIds))
	for index, pid := range permissionIds {
		ids[index] = int(pid)
	}
	count, err := p.data.db.Menu.Query().Where(menu.IDIn(ids...)).Count(ctx)
	if err != nil {
		return err
	}
	if count != len(permissionIds) {
		return errors.New(500, "Permission NOT EXISTS", "权限不存在")
	}
	policies := make([][]string, len(permissionIds))
	for index, pid := range permissionIds {
		policies[index] = []string{strconv.Itoa(r.ID), strconv.FormatUint(pid, 10), "*"}
	}
	_, err = p.data.casbinAdapter.AddPolicies(policies)
	if err != nil {
		return err
	}
	return nil
}
