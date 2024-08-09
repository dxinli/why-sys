package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"why-sys/internal/biz"
	"why-sys/internal/data/ent/user"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func (a *authRepo) Check(ctx context.Context, user *biz.SecurityUser) bool {
	ok, err := a.data.casbinAdapter.Enforce(user.AuthorityId, user.Path, user.Method)
	if err != nil {
		return false
	}
	return ok
}

func (a *authRepo) Create(ctx context.Context, u *biz.User) (int, error) {
	user, err := a.data.db.User.Create().
		SetUserName(u.Username).
		SetPassword(u.Password).
		SetEmail(u.Email).
		Save(ctx)

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (a *authRepo) Login(context context.Context, loginUser *biz.User) (int, error) {
	// 查询用户
	id, err := a.data.db.User.Query().Where(
		user.UserName(loginUser.Username),
		user.Password(loginUser.Password),
	).OnlyID(context)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
