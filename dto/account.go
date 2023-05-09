package dto

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

type Credentials struct {
	Username   string `json:"Username" form:"Username" query:"Username"`
	Password   string `json:"Password" form:"Password" query:"Password"`
	ForceLogin int    `json:"ForceLogin"`
	ClientID   string `json:"ClientID" form:"ClientID" query:"ClientID"`
	UserID     string `json:"UserID" form:"UserID" query:"UserID"`
	ThirdParty int    `json:"ThirdParty" form:"ThirdParty" query:"ThirdParty"`
}

func GenerateTempUsername(username string) string {
	now := strconv.FormatInt(time.Now().Add(24*3600*time.Second).Unix(), 10)
	tempUserName := now + ":" + username
	return tempUserName
}

func GenerateTempPassword(tempUserName string) string {
	secretKey := viper.GetString("TurnServer.SecretKey")
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(tempUserName))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}