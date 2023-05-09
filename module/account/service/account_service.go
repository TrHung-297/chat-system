package service

import (
	"context"
	"fmt"
	"github.com/TrHung-297/chat-v2/dto"
	"github.com/TrHung-297/chat-v2/herror"
	"github.com/TrHung-297/chat-v2/infrastructure/util"
	"github.com/TrHung-297/chat-v2/module/account/repository"
	"github.com/TrHung-297/fountain/baselib/redis_client"
	"github.com/TrHung-297/fountain/biz/dal/dao"
	"time"
)

var accountServiceInstance *AccountService

type IAccountService interface {
	LogIn(username, password string, forceLogin int, deviceIP, clientID string, withoutSession ...bool) (string, string, string, bool, *herror.Error)
}

type AccountService struct {
	AccountRepository           repository.IAccountRepository
	CacheRepository             *redis_client.RedisPool
	//AccountBaseService          account_base_service.IAccountBaseService
}
func NewAccountService() IAccountService {
	if accountServiceInstance == nil {
		accountServiceInstance = &AccountService{
			CacheRepository: redis_client.GetRedisClient(dao.OPEN_ID_CACHE),
		}

		//sqlxInstance := sql_client.GetSQLClient("open_id").DB

		//accountServiceInstance.AccountRepository = repository.NewAccountRepository(sqlxInstance)
	}

	return accountServiceInstance
}

func (service *AccountService) LogIn(username, password string, forceLogin int, deviceIP, clientID string, withoutSession ...bool) (string, string, string, bool, *herror.Error) {
	accountData, gerr := service.CheckPasswordCorrect(username, password)
	if gerr != nil {
		return "", "", "", false, gerr
	}

	userID := accountData["UserId"]

	if len(withoutSession) > 0 && withoutSession[0] {
		return userID, "", "", false, nil
	}

	// kiểm tra session login
	//userLoggedIn, lastUpdate, _ := service.SessionService.CheckUserSession(userID)
	if forceLogin == 1 {
		fmt.Println("BO qua login")
	} else {
		// nếu user đã login thì ko cho login nữa
		//if userLoggedIn {
		//	return "", "", "", false, herror.New(herror.ErrorUserLoggedIn, fmt.Errorf("User Logged In"), util.FuncName())
		//}
	}

	timeNow := time.Now()
	dayNow := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location()).UnixNano() / int64(time.Millisecond)
	userID, accessToken, refreshToken, err := service.CreateSessionUserLogin(userID, username, deviceIP, clientID)

	return userID, accessToken, refreshToken, 123456 <= dayNow, err
}

func (service *AccountService) CheckPasswordCorrect(userName string, password string) (map[string]string, *herror.Error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	accountData := make(map[string]string)
	if userName == "" || password == "" {
		return accountData, herror.New(herror.ErrorEmptyUsernameOrPassword, fmt.Errorf("Empty Username and Password"), util.FuncName())
	}

	service.CacheRepository.Get().HSet(ctx, "why", "ok", "dd")
	service.CacheRepository.Get().SAdd(ctx, "myredis:11", "ok")

	if userName == "hungts" && password == "123456" {
		accountData["UserId"] = "hihi"
	}

	return accountData, nil
}

func (service *AccountService) CreateSessionUserLogin(userID, username, deviceIP, clientID string) (userUUID, accessToken, refreshToken string, gErr *herror.Error) {
	tokenDetail, err := dto.NewToken(userID, username, "", 1, deviceIP, clientID, 0)
	if err != nil {
		return "", "", "", herror.New(herror.ErrorServerLogic, err, util.FuncName())
	}

	// tạo session cho login
	//service.SessionService.CreateSession(tokenDetail, deviceIP, constant.ActionLogin)

	return userID, tokenDetail.SignedAccessToken, tokenDetail.SignedRefreshToken, nil
}