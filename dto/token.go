package dto

import (
	"github.com/TrHung-297/chat-v2/constant"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	"github.com/google/uuid"
	"time"
)

const (
	GPlayType        int = 0
	DeviceKindPC     int = 1
	DeviceKindMobile int = 3
	DeviceKindWeb    int = 5

)

type TokenDetails struct {
	UserId             string
	UserName           string
	DeviceKind         int
	AccessUUID         uuid.UUID
	RefreshUUID        uuid.UUID
	AvatarUrl          string
	AccessToken        *jwt.Token
	RefreshToken       *jwt.Token
	SignedAccessToken  string
	SignedRefreshToken string
	AtExpires          time.Duration
	RtExpires          time.Duration
	SecureExpire       time.Duration
}

const (
	KTokenAuthorizedKey  string = "authorized"
	KTokenAccessUUIDKey  string = "access_uuid"
	KTokenRefreshUUIDKey string = "refresh_uuid"
	KTokenUserIDKey      string = "user_id"
	KTokenExpKey         string = "exp"
	KTokenSecureExpKey   string = "secure_exp"
	KTokenAvatarURLKey   string = "avatarUrl"
	KTokenUserNameKey    string = "userName"
	KTokenDeviceKindKey  string = "device_kind"
	KTokenDeviceIPKey    string = "device_ip"
	KTokenDeviceIDKey    string = "device_id"
	KTokenThirdPartyType string = "third_party_type"
)

func Sign(token *TokenDetails, signer string) (string, string, error) {
	signedAt, err := token.AccessToken.SignedString([]byte(signer))
	if err != nil {
		return "", "", err
	}
	token.SignedAccessToken = signedAt

	signedRt, err := token.RefreshToken.SignedString([]byte(signer))
	if err != nil {
		return "", "", err
	}
	token.SignedRefreshToken = signedRt
	return signedAt, signedRt, nil
}

func NewToken(accountID, username, avatar string, deviceKind int, deviceIP, deviceID string, thirdPartyType int) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.UserId = accountID
	td.UserName = username
	td.DeviceKind = deviceKind

	// Đối với các bên thứ 3 khác mà sản phẩm vẫn là của GPlay thì chỉ tối đa 1h token
	// Đối với các sản phẩm chính GPlay thì app: 1day; mobile-web: 1 month
	if deviceKind == DeviceKindPC {
		td.AtExpires = constant.CacheExpiresInOneDay
	} else if thirdPartyType != GPlayType {
		td.AtExpires = constant.CacheExpiresInOneHour
	} else {
		td.AtExpires = constant.CacheExpiresInOneMonth
	}
	td.SecureExpire = constant.CacheExpiresInOneHour
	td.RtExpires = constant.CacheExpiresInOneMonth

	if deviceID == "" {
		deviceID = uuid.New().String()
	}

	accessUUID := uuid.New()
	td.AccessUUID = accessUUID
	atClaims := jwt.MapClaims{}
	atClaims[KTokenAuthorizedKey] = true
	atClaims[KTokenAccessUUIDKey] = accessUUID.String()
	atClaims[KTokenUserIDKey] = accountID
	atClaims[KTokenExpKey] = time.Now().Add(td.AtExpires).Unix()
	atClaims[KTokenAvatarURLKey] = td.AvatarUrl
	atClaims[KTokenUserNameKey] = td.UserName
	atClaims[KTokenDeviceKindKey] = deviceKind
	atClaims[KTokenDeviceIPKey] = deviceIP
	atClaims[KTokenDeviceIDKey] = deviceID
	atClaims[KTokenThirdPartyType] = thirdPartyType

	atClaims[KTokenSecureExpKey] = time.Now().Add(td.SecureExpire).Unix()
	td.AccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	refreshUUID := uuid.New()
	td.RefreshUUID = refreshUUID
	rtClaims := jwt.MapClaims{}
	rtClaims[KTokenRefreshUUIDKey] = refreshUUID.String()
	rtClaims[KTokenUserIDKey] = accountID
	rtClaims[KTokenExpKey] = time.Now().Add(td.RtExpires).Unix()
	rtClaims[KTokenAvatarURLKey] = td.AvatarUrl
	rtClaims[KTokenUserNameKey] = td.UserName
	rtClaims[KTokenDeviceKindKey] = deviceKind
	rtClaims[KTokenDeviceIPKey] = deviceIP
	atClaims[KTokenDeviceIDKey] = deviceID
	atClaims[KTokenThirdPartyType] = thirdPartyType

	td.RefreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	signKey := viper.GetString("OpenIDJwt.SecretKey")
	_, _, err := Sign(td, signKey)

	return td, err
}
