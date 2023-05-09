package util

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"runtime"
	"strings"
)

const PasswordBcrypt = 0
const PasswordMd5 = 2

func FuncName() string {
	pc, _, line, _ := runtime.Caller(1)
	result := fmt.Sprintf("%s:%v", runtime.FuncForPC(pc).Name(), line)
	return result
}

func UsernameValidator(username string) error {
	if len(username) < 4 {
		return fmt.Errorf("Username too short")
	}

	if len(username) > 32 {
		return fmt.Errorf("Username too long")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+(?:[_.-][a-zA-Z0-9]+)*$`)
	if !re.MatchString(username) {
		return fmt.Errorf("Username Wrong Format")
	}
	if strings.Contains(username, " ") {
		return fmt.Errorf("Username cannot have spaces")
	}
	return nil
}

func VerifyPassword(currentPass string, hashPassword string, hashType int) bool {
	if hashType == PasswordBcrypt {
		err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(currentPass))
		if err == nil {
			return true
		}
		g_log.V(1).Errorf("Error BCrypt: ", err)
	} else if hashType == PasswordMd5 {
		//currentPasswordMd5 := md5.Sum([]byte(currentPass))
		tmp := strings.Split(hashPassword, "|")
		if len(tmp) == 2 {
			passMd5 := md5.Sum([]byte(currentPass))
			salt := tmp[1]
			passMd5String := fmt.Sprintf("%x%s", passMd5, salt)
			passMd5WithSalt := md5.Sum([]byte(passMd5String))
			passMd5WithSaltString := fmt.Sprintf("%x", passMd5WithSalt)
			if passMd5WithSaltString == tmp[0] {
				return true
			}
			return false
		}
	} else {
		currentPasswordSha256 := sha256.Sum256([]byte(currentPass))
		passwordSha256Checksum := fmt.Sprintf("%x", currentPasswordSha256)
		if passwordSha256Checksum == hashPassword {
			return true
		}
	}
	return false
}