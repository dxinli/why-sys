package service

import (
	"context"
	pb "why-sys/api/auth/v1"
	"why-sys/internal/biz"
)

func (s *AuthService) GrantRolePermission(ctx context.Context, req *pb.GrantRolePermissionRequest) (*pb.GrantRolePermissionResponse, error) {
	err := s.pc.GrantPermission(ctx, &biz.Permission{
		RoleID:        req.RoleId,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GrantRolePermissionResponse{
		Msg: "授权成功",
	}, nil
}
