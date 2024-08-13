package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
	"why-sys/internal/biz"

	pb "why-sys/api/auth/v1"
)

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.uc.Login(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: token,
	}, nil
}
func (s *AuthService) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	// TODO 通过 ctx 中能够获取到 jwtChaim 不需要重新获取 token
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, errors.New(500, "ERROR_GET_FROM_SERVER_CONTEXT", "请求元数据获取失败")
	}
	authorizationValue := tr.RequestHeader().Get("Authorization")
	if !strings.HasPrefix(authorizationValue, "Bearer ") {
		// 如果没有找到有效的Bearer令牌，则返回错误
		return nil, errors.New(500, "ERROR_NON_BEARER_SCHEME", "认证信息格式错误")
	}
	// 使用SplitN分割字符串，只分割一次
	parts := strings.SplitN(authorizationValue, " ", 2)
	if len(parts) != 2 {
		// 如果没有找到有效的Bearer令牌，则返回错误
		return nil, errors.New(500, "ERROR_NON_TOKEN", "token为空")
	}
	ok = s.uc.Check(ctx, parts[1], req.Uri, req.Method)
	if !ok {
		return nil, errors.New(500, "ERROR_CHECK_TOKEN", "没有权限访问资源")
	}
	return &pb.CheckResponse{Ok: ok, Msg: "成功"}, nil
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := s.uc.Register(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		Msg: "注册成功,用户id:" + id,
	}, nil
}
