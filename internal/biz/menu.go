package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Menu struct {
	ID         int
	Name, Desc string
	ParentID   int
	Children   []*Menu
	Parent     *Menu
	Path       string
	Component  string
	Sort       float64
	Level      int16
	Leaf       bool
}

type MenuRepo interface {
	CreateMenu(context.Context, *Menu) (int, error)
}

type MenuUsecase struct {
	repo MenuRepo
	log  *log.Helper
}

func NewMenuUsecase(repo MenuRepo, logger log.Logger) *MenuUsecase {
	return &MenuUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (mc MenuUsecase) CreateMenu(ctx context.Context, menu *Menu) (int, error) {
	return mc.repo.CreateMenu(ctx, menu)
}
