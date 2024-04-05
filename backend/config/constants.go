package config

const (
	// Middlewares context
	PayloadUserId   = "payload_user_id"
	PayloadUserRole = "payload_user_role"

	// Request
	QueryId     = "id"
	SearchParam = "search"
	PageParam   = "page"
	UserId      = "user_id"
	Page        = "page"

	// Roles
	AdminRole = "admin"
	UserRole  = "user"

	// Auth
	AuthorizationHeader = "authorization"
	JWTClaimsId         = "id"
	JWTClaimsRole       = "role"
	JWTClaimsExp        = "exp"
)
