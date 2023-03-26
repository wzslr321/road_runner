package auth

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Tokener interface {
	ValidateToken(authorization []string) error
}

type Jwt struct{}

var (
	ErrMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	ErrInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

func (j *Jwt) ValidateToken(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	// implement oauth2?
	return token == "some-secret-token"
}
