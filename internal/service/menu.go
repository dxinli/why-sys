package service

import (
	"context"
	"strconv"
	pb "why-sys/api/auth/v1"
	"why-sys/internal/biz"
)

func (s *AuthService) CreateMenu(ctx context.Context, detail *pb.MenuDetail) (*pb.CreateMenuResponse, error) {
	id, err := s.mc.CreateMenu(ctx, &biz.Menu{
		Name:      detail.Name,
		Path:      detail.Path,
		Sort:      detail.Sort,
		ParentID:  int(detail.ParentId),
		Component: detail.Component,
		Level:     int16(detail.Level),
		Desc:      detail.Description,
		Leaf:      detail.Leaf,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateMenuResponse{
		Msg: "创建成功，菜单id:" + strconv.Itoa(id),
	}, nil
}

func (s *AuthService) ListMenu(ctx context.Context, request *pb.ListMenuRequest) (*pb.MenuDetail, error) {
	//TODO implement me
	panic("implement me")
}
