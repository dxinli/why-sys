package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"why-sys/internal/biz"
)

type menuRepo struct {
	data *Data
	log  *log.Helper
}

func NewMenuRepo(data *Data, logger log.Logger) biz.MenuRepo {
	return &menuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (m menuRepo) CreateMenu(ctx context.Context, menu *biz.Menu) (int, error) {
	me, err := m.data.db.Menu.Create().
		SetMenuName(menu.Name).
		SetComponent(menu.Component).
		SetMenuDesc(menu.Desc).
		SetLeaf(menu.Leaf).
		SetLevel(menu.Level).
		SetPath(menu.Path).SetParentID(menu.ParentID).SetSort(menu.Sort).Save(ctx)
	if err != nil {
		return 0, err
	}

	return me.ID, nil
}
