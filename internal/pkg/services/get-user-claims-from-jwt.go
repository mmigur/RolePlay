package services

import (
	"RolePlayModule/internal/utils/config"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func GetUserClaimsFromJWT(ctx context.Context, cfg config.Config) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")

	claims, err := DecodeJWT(token, cfg)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "decode token err: %v", err)
	}
	return claims, nil
}
