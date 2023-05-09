package herror

import "strconv"

/**
 * Translate
 */
func T(errorCode uint32) string {
	switch errorCode {
	//////////////////////////
	// Client-side
	//////////////////////////
	case ErrorBindData:
		return "Failed to bind data"
	case ErrorValidData:
		return "Failed to valid data"
	case ErrorEasyPassword:
		return "Password not strong"
	case ErrorWrongPassword:
		return "Wrong Password"
	case ErrorWrongUsername:
		return "Wrong Username"
	case ErrorEmptyUsernameOrPassword:
		return "Empty Username or Password"
	case ErrorUserLoggedIn:
		return "User Logged In"
	case ErrorUsernameTaken:
		return "Username exists"
	case ErrorWrongEmailFormat:
		return "Wrong Email Format"
	case ErrorEmailNotMatch:
		return "Email Not Match"
	case ErrorEmailTaken:
		return "Email exists"
	case ErrorPhoneNumberTaken:
		return "Số điện thoại đã tồn tại"
	case ErrorDisplayNameTaken:
		return "Tên hiển thị đã tồn tại"
	case ErrorAccountLocked:
		return "Account Locked"
	case ErrorPasswordExpired:
		return "Password Expired"
	case ErrorRefreshTokenInvalid:
		return "Refresh Token Invalid"
	case ErrorTokenInvalid:
		return "Phiên đăng nhập hết hạn, vui lòng đăng nhập lại"
	case ErrorCaptcha:
		return "Captcha Failed"
	case ErrorFriendRequested:
		return "Friend Request Sent"
	case ErrorUserNotFriend:
		return "Not Friend"
	case ErrorAddSelfAsFriend:
		return "Cannot add self as Friend"
	case ErrorDisplayNameFormat:
		return `Tên hiển thị không hợp lệ`
	case ErrorUnknow3rdProvider:
		return "Unknown 3rd Party"
	case ErrorOauthProviderIdInvalid:
		return "Provider ID Invalid"
	case ErrorOauthGetUserInfo:
		return "Cannot Get Info From OAuth Token"
	case ErrorOauthProviderIdNotMatch:
		return "Oauth Provider ID Not match with Token"
	case ErrorChangeSamePassword:
		return "Same with Current Password "
	case ErrorMaxFriendsReach:
		return "User Max Friend Reach"
	case ErrorMaxFriendsFriendReach:
		return "Friend Max Friends Reach"
	case ErrorIpReachMaxRegistered:
		return "Reach Max Ip Registered"
	case ErrorWrongMail:
		return "Invalid Email Format"
	case ErrorAuthentication:
		return "Unauthorize Key"
	case ErrorReachMaxRegistered:
		return "Reach Max Registered"
	case ErrorEarlyLogin:
		return "Login Gate Not Open"
	case ErrorAccountActived:
		return "Account Activated"
	case ErrorTokenLocaltionChanged:
		return "Token Localtion Changed"
	case ErrorAccountBanned:
		return "Account Banned"
	case ErrorApiKey:
		return "Api Keys Unauthorize"
	case ErrorInvalidPhoneNumber:
		return "Số điên thoại không hợp lệ"
	case ErrorInvalidAddress:
		return "Địa chỉ quá dài, vui lòng nhập địa chỉ ngắn hơn"
	case ErrorDisplaynameChangePermission:
		return "Bạn đã hết lượt đổi tên"
	//case ErrorIpReachMaxRegistered:
	//	return
	//////////////////////////
	// Server-side
	//////////////////////////
	case ErrorConnect:
		return "Failed to connect database"
	case ErrorSaveData:
		return "Failed to save data"
	case ErrorRetrieveData:
		return "Failed to retrieve data"
	case ErrorLogin:
		return "Failed to login. Please try again!"
	case ErrorCreateAccount:
		return "Failed to Create Account"
	case ErrorCreateToken:
		return "Failed to Create Token"
	case ErrorServerConfig:
		return "Server Missing Config"
	case ErrorServerLogic:
		return "Failed to handle server logic"
	case ErrorResetPassword:
		return "Failded to Reset Password"
	case ErrorCreateEmailConfirm:
		return "Failded to Create Email Confirmation"
	case ErrorSystemMaintenance:
		return "The system is maintenance"
	case ErrorAccountDeleted:
		return "Account is Deleted"
	}

	return "Unknown Error Code: " + strconv.Itoa(int(errorCode))
}