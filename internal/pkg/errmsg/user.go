package errmsg

const (
	ErrorMsg_UserNotFound        = "User Not Found"
	ErrorMsg_UserDuplicated      = "User Duplicated"
	ErrorMsg_UserCreation        = "User Creation Error"
	ErrorMsg_WrongPassword       = "Email Or Password Are Incorrect"
	ErrorMsg_UserAlreadyVerified = "User's Email is Already Verified!"

	// Redis Cache Errors
	ErrorMsg_UserRedisSetError         = "Error Setting User in Redis Cache"
	ErrorMsg_UserRedisGetIDError       = "Error Getting User by ID in Redis Cache"
	ErrorMsg_UserRedisGetEmailError    = "Error Getting User by Email in Redis Cache"
	ErrorMsg_UserRedisDeleteEmailError = "Error Deleting User by Email in Redis Cache"
	ErrorMsg_UserRedisDeleteIDError    = "Error Deleting User by ID in Redis Cache"
)
