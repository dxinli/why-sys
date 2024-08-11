package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
	"why-sys/internal/biz"
	"why-sys/internal/data/ent/role"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

func (repo *roleRepo) GrantUserRole(ctx context.Context, userId, roleId uint64) error {
	// 校验用户
	u, err := repo.data.db.User.Get(ctx, int(userId))
	if err != nil {
		return err
	}
	// 校验角色
	r, err := repo.data.db.Role.Get(ctx, int(roleId))
	if err != nil {
		return err
	}
	_, err = repo.data.casbinAdapter.AddRoleForUser(strconv.Itoa(u.ID), strconv.Itoa(r.ID))
	if err != nil {
		return err
	}
	return nil
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *roleRepo) Create(ctx context.Context, role *biz.Role) (int, error) {
	createRole, err := repo.data.db.Role.Create().SetRoleName(role.Name).SetRoleDesc(role.Desc).Save(ctx)
	if err != nil {
		return 0, err
	}

	return createRole.ID, nil
}

func (repo *roleRepo) Delete(ctx context.Context, role *biz.Role) error {
	//TODO implement me
	panic("implement me")
}

func (repo *roleRepo) Get(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *roleRepo) List(ctx context.Context, r *biz.Role) ([]*biz.Role, error) {
	roleQuery := repo.data.db.Role.Query()
	name := r.Name
	if name != "" {
		roleQuery = roleQuery.Where(role.RoleName(name))
	}
	all, err := roleQuery.All(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*biz.Role, 0, len(all))
	for _, item := range all {
		roles = append(roles, &biz.Role{
			ID:   item.ID,
			Name: item.RoleName,
			Desc: item.RoleDesc,
		})
	}
	return roles, nil
}

func (repo *roleRepo) Update(ctx context.Context, role *biz.Role) (*biz.Role, error) {
	//TODO implement me
	panic("implement me")
}
