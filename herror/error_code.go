package herror

type Error struct {
	Error error
	Code  uint32
	Line  string
}

func New(code uint32, err error, line string) *Error {
	return &Error{
		Error: err,
		Code:  code,
		Line:  line,
	}
}

/********************************************************************/
/* Client-side Error Code											*/
/********************************************************************/

const (
	ErrorBindData                    uint32 = 400000
	ErrorValidData                   uint32 = 400001
	ErrorEasyPassword                uint32 = 400002
	ErrorEasyNewPassword             uint32 = 400003
	ErrorWrongPassword               uint32 = 400004
	ErrorWrongUsername               uint32 = 400005
	ErrorUserLoggedIn                uint32 = 400006
	ErrorEmptyUsernameOrPassword     uint32 = 400007
	ErrorUsernameTaken               uint32 = 400008
	ErrorEmailTaken                  uint32 = 400009
	ErrorPhoneNumberTaken            uint32 = 400010
	ErrorDisplayNameTaken            uint32 = 400011
	ErrorWrongEmailFormat            uint32 = 400012
	ErrorEmailNotMatch               uint32 = 400013
	ErrorAccountLocked               uint32 = 400014
	ErrorPasswordExpired             uint32 = 400015
	ErrorTokenInvalid                uint32 = 400016
	ErrorRefreshTokenInvalid         uint32 = 400017
	ErrorCaptcha                     uint32 = 400020
	ErrorFriendRequested             uint32 = 400021
	ErrorPasswordResetCode           uint32 = 400022
	ErrorUserNotFriend               uint32 = 400023
	ErrorAddSelfAsFriend             uint32 = 400024
	ErrorDisplayNameFormat           uint32 = 400025
	ErrorUnknow3rdProvider           uint32 = 400026
	ErrorOauthProviderIdInvalid      uint32 = 400027
	ErrorOauthGetUserInfo            uint32 = 400028
	ErrorOauthProviderIdNotMatch     uint32 = 400033
	ErrorChangeSamePassword          uint32 = 400029
	ErrorMaxFriendsReach             uint32 = 400030
	ErrorMaxFriendsFriendReach       uint32 = 400031
	ErrorEarlyLogin                  uint32 = 400035
	ErrorEmailActived                uint32 = 400038
	ErrorDisplaynameChangePermission uint32 = 400039
	ErrorApiKey                      uint32 = 400040
	ErrorInvalidAddress              uint32 = 400041
	ErrorInvalidPhoneNumber          uint32 = 400042
	ErrorAccountDeleted              uint32 = 400043

	ErrorAuthentication        uint32 = 40200
	ErrorIpReachMaxRegistered  uint32 = 40300
	ErrorWrongMail             uint32 = 40301
	ErrorInvalidEarlyCode      uint32 = 40302
	ErrorReachMaxRegistered    uint32 = 40303
	ErrorAccountActived        uint32 = 40304
	ErrorTokenLocaltionChanged uint32 = 40305
	ErrorAccountBanned         uint32 = 40306
)

// var ErrMissing = "MISSING"
// var ErrTaken = "TAKEN"
// var ErrFormatInvalid = "FORMAT_INVALID"
// var ErrInsecure = "INSECURE"
// var ErrFailed = "FAILED"
// var ErrLocked = "LOCKED"
// var ErrExpired = "EXPIRED"
// var ErrNotFound = "NOT_FOUND"
// var ErrInvalidOrExpired = "INVALID_OR_EXPIRED"

/********************************************************************/
/* Server-side Error Code											*/
/********************************************************************/
const (
	ErrorConnect            uint32 = 500000
	ErrorSaveData           uint32 = 500001
	ErrorRetrieveData       uint32 = 500002
	ErrorLogin              uint32 = 500003
	ErrorCreateAccount      uint32 = 500004
	ErrorCreateToken        uint32 = 500005
	ErrorServerConfig       uint32 = 500006
	ErrorServerLogic        uint32 = 500007
	ErrorResetPassword      uint32 = 500008
	ErrorCreateEmailConfirm uint32 = 500009
	ErrorOther              uint32 = 500010
	ErrorSystemMaintenance  uint32 = 500011
)