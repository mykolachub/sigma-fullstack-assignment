package config

import "errors"

var (
	// Database related errors
	ErrFailedCreateUser = errors.New("failed to create user")
	ErrFailedUpdateUser = errors.New("failed to update user")
	ErrFailedDeleteUser = errors.New("failed to delete user")
	ErrFailedTrackPage  = errors.New("failed to track page")
	ErrFailedGetPage    = errors.New("failed to get page")
	ErrFailedGetUser    = errors.New("failed to get users")
	ErrNoUser           = errors.New("no such user")
	ErrUserExists       = errors.New("user already exists")

	// User request related errors
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmptyUpdateBody    = errors.New("empty update body")
	ErrFailedReadBody     = errors.New("failed to read body")
	ErrMissingIdPar       = errors.New("missing id parameter")
	ErrMissingPagePar     = errors.New("missing page parameter")

	// Security errors
	ErrFailedHashPassword = errors.New("failed to hash password")

	// Authorization related errors
	ErrNoPermissions    = errors.New("no permissions")
	ErrNoAuthHeader     = errors.New("authorization header not provided")
	ErrInvalidBearer    = errors.New("invalid authorization header format")
	ErrInvalidAuthType  = errors.New("unsupported authorization type")
	ErrInvalidToken     = errors.New("invalid token")
	ErrFailedParseToken = errors.New("failed to parse token")
)

const (
	MsgEmpty       = ""
	MsgUserCreated = "user created"
	MsgUserUpdated = "user created"
	MsgUserDeleted = "user deleted"
	MsgPageCountet = "page counted"
)

const (
	// Middlewares context
	PayloadUserId   = "payload_user_id"
	PayloadUserRole = "payload_user_role"

	// Request
	QueryId = "id"
	Page    = "page"

	// Roles
	AdminRole = "admin"
	UserRole  = "user"

	// Auth
	AuthorizationHeader = "authorization"
	JWTClaimsId         = "id"
	JWTClaimsRole       = "role"
	JWTClaimsExp        = "exp"
)
