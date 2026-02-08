package config

import "time"

const (
	JwtSignKey                 = "jwt_secret"
	AccessTokenSubject         = "ac"
	RefreshTokenSubject        = "rt"
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	AuthMiddlewareContextKey   = "claims"
	SigningMethodJwt           = "HS256"
	PaginationLimitDefault     = 20
	PaginationOffsetDefault    = 0
	PaginationLimitMax         = 40
)
