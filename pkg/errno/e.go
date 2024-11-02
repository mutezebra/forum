package errno

// Verify
const (
	InvalidEmailFormat = 1000 + iota
	InvalidUserNameFormat
	InvalidPasswordFormat
	InvalidVerifyCodeFormat
	SpaceUserName
	LackToken
	WrongToken
	TokenExpire
	LackOFAvatarFile
	InvalidParam
)

const (
	ThreadTitleLengthInvalid = 1000 + iota
	ThreadContentTooShort
)
