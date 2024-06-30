package driven

import (
	"github.com/nullexp/finman-auth-service/internal/port/model"
)

type Claims interface {
	Valid() error
}

type JwtClaim interface {
	GetExpireTime() int64
	GetSubject() string
	GetIssuer() string
	GetAudience() []string
	GetIssuedAt() int64
	GetIdentity() string
	IsExpired() bool
}

type TokenService interface {
	CreateToken(sb model.Subject) (string, error)
	GetToken(tokenString string) (model.StandardClaims, error)
	CheckToken(tokenString string) (bool, error)
	GetSubject(subject string) (out model.Subject, err error)
}
