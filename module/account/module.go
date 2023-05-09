package account

import (
	"github.com/TrHung-297/chat-v2/module/account/controller"
	"github.com/TrHung-297/chat-v2/module/account/service"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type AccountModule struct {
	mAccountController *controller.AccountController
}

var accountModuleInstance *AccountModule

func Initialize(apiServer *fiber.App) *AccountModule {
	if accountModuleInstance != nil {
		return accountModuleInstance
	}

	accountService := service.NewAccountService()

	accountModuleInstance = &AccountModule{
		mAccountController: controller.NewAccountController(accountService, viper.GetStringSlice("WhiteList.UserId")),
	}

	// New router
	accountModuleInstance.InitRouter(apiServer)

	return accountModuleInstance
}

func (m *AccountModule) InitRouter(apiServer *fiber.App) {
	gv2 := apiServer.Group("/users/api/v2.0/account/")
	gv2.Post("login", m.mAccountController.LoginAccount)
}