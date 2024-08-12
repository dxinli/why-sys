package service

import (
	"context"
	"strconv"
	pb "why-sys/api/auth/v1"
	"why-sys/internal/biz"
)

func (s *AuthService) CreateRole(ctx context.Context, request *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	id, err := s.rc.CreateRole(ctx, &biz.Role{
		Name: request.Name,
		Desc: request.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoleResponse{
		Msg: "创建成功:角色id:" + strconv.Itoa(id),
	}, nil
}

func (s *AuthService) DeleteRole(ctx context.Context, request *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthService) GetRole(ctx context.Context, request *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthService) ListRole(ctx context.Context, request *pb.ListRoleRequest) (*pb.ListRoleResponse, error) {
	roles, err := s.rc.ListRole(ctx, &biz.Role{
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}
	roleList := make([]*pb.GetRoleResponse, 0, len(roles))
	for _, role := range roles {
		roleList = append(roleList, &pb.GetRoleResponse{
			Id:          strconv.Itoa(role.ID),
			Name:        role.Name,
			Description: role.Desc,
		})
	}

	return &pb.ListRoleResponse{
		Data: roleList,
	}, nil
}

func (s *AuthService) UpdateRole(ctx context.Context, request *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthService) GrantUserRole(ctx context.Context, req *pb.GrantUserRoleRequest) (*pb.GrantUserRoleResponse, error) {
	err := s.rc.GrantUserRole(context.Background(), req.UserId, req.RoleId)
	if err != nil {
		return nil, err
	}
	return &pb.GrantUserRoleResponse{
		Msg: "授权成功",
	}, nil
}
