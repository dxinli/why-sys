package service

import (
	"context"
	pb "why-sys/api/auth/v1"
	"why-sys/internal/biz"
)

func (s *AuthService) CreateMenu(ctx context.Context, detail *pb.MenuDetail) (*pb.CreateMenuResponse, error) {
	menu := &biz.Menu{
		Name:      detail.Name,
		Path:      detail.Path,
		Sort:      detail.Sort,
		ParentID:  int(detail.ParentId),
		Component: detail.Component,
		Level:     uint16(detail.Level),
		Desc:      detail.Description,
		Leaf:      detail.Leaf,
	}
	id, err := s.mc.CreateMenu(ctx, menu)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMenuResponse{
		Msg: "创建成功，菜单id:" + id,
	}, nil
}

func (s *AuthService) ListMenu(ctx context.Context, request *pb.ListMenuRequest) (*pb.MenuDetail, error) {
	//TODO implement me
	panic("implement me")
}
