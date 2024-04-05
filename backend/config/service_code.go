package config

import "errors"

type ServiceCode struct {
	// Message of the occurred action
	Message string

	// Code of the occurred action
	/*
		1xxx - Database related
		2xxx - Request related
		3xxx - Security related
		4xxx - Authorization related
		5xxx - Server related
		6xxx - Other
	*/
	Code int
}

// Converts service code message to the type error
func (s ServiceCode) ToError() error {
	return errors.New(s.Message)
}

func NewServiceCode(code int, message string) ServiceCode {
	return ServiceCode{Code: code, Message: message}
}

var (
	// Database related errors
	SvcFailedCreateUser = NewServiceCode(11001, "failed to create user")
	SvcFailedUpdateUser = NewServiceCode(11002, "failed to update user")
	SvcFailedDeleteUser = NewServiceCode(11003, "failed to delete user")
	SvcFailedGetUser    = NewServiceCode(11004, "failed to get users")
	SvcNoUser           = NewServiceCode(11005, "no such user")
	SvcUserExists       = NewServiceCode(11006, "user already exists")
	SvcUserCreated      = NewServiceCode(11007, "user created")
	SvcUserUpdated      = NewServiceCode(11008, "user created")
	SvcUserDeleted      = NewServiceCode(11009, "user deleted")
	SvcPageTracked      = NewServiceCode(11010, "page tracked")
	SvcFailedTrackPage  = NewServiceCode(12001, "failed to track page")
	SvcFailedGetPage    = NewServiceCode(12002, "failed to get page")

	// Client request related errors
	SvcInvalidCredentials = NewServiceCode(20001, "invalid credentials")
	SvcEmptyUpdateBody    = NewServiceCode(20002, "empty update body")
	SvcFailedReadBody     = NewServiceCode(20003, "failed to read body")
	SvcMissingUserIdPar   = NewServiceCode(21001, "missing user id parameter")
	SvcMissingPagePar     = NewServiceCode(22001, "missing page parameter")
	SvcPageParNotInt      = NewServiceCode(22002, "page parameter must be number")

	// Security errors
	SvcFailedHashPassword = NewServiceCode(30001, "failed to hash password")

	// Authorization related errors
	SvcNoPermissions     = NewServiceCode(40001, "no permissions")
	SvcNoAuthHeader      = NewServiceCode(40002, "authorization header not provided")
	SvcInvalidBearer     = NewServiceCode(40003, "invalid authorization header format")
	SvcInvalidAuthType   = NewServiceCode(40004, "unsupported authorization type")
	SvcInvalidToken      = NewServiceCode(40005, "invalid token")
	SvcFailedParseToken  = NewServiceCode(40006, "failed to parse token")
	SvcFailedCreateToken = NewServiceCode(40007, "failed to generate token")

	// Server
	SvcServiceBusy = NewServiceCode(50001, "service is busy, try again later")

	// Other
	SvcEmptyMsg = NewServiceCode(60001, "")
)
