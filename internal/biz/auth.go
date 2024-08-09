package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type User struct {
	Username string
	Password string
	Email    string

	jwt.RegisteredClaims // v5版本新加的方法
}

type SecurityUser struct {
	AuthorityId, Path, Method string
}

type AuthRepo interface {
	Login(context.Context, *User) (int, error)

	Create(context.Context, *User) (int, error)

	Check(context.Context, *SecurityUser) bool
}

type AuthUsecase struct {
	log *log.Helper
	AuthRepo
}

func createAccessJwtToken(username string, secretKey []byte) (string, error) {
	claims := User{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // 过期时间1小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(secretKey)

	return s, err
}

func parseJwt(tokenString, secretKey string) (*User, error) {
	t, err := jwt.ParseWithClaims(tokenString, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*User); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func NewAuthUsecase(repo AuthRepo, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{log.NewHelper(logger), repo}
}

func (uc *AuthUsecase) Login(ctx context.Context, user *User) (string, error) {
	_, err := uc.AuthRepo.Login(ctx, user)
	if err != nil {
		return "", err
	}
	token, err := createAccessJwtToken(user.Username, []byte("test_jwt"))
	if err != nil {
		return "", err
	}
	// 生成 jwt
	return token, nil
}

func (uc *AuthUsecase) Register(ctx context.Context, user *User) (string, error) {
	id, err := uc.AuthRepo.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(id), nil
}

func (uc *AuthUsecase) Check(ctx context.Context, token, uri, method string) bool {
	// 解析 token 获取用户信息
	user, err := parseJwt(token, "test_jwt")
	if err != nil {
		return false
	}
	return uc.AuthRepo.Check(ctx, &SecurityUser{Path: uri, Method: method, AuthorityId: user.Username})
}
