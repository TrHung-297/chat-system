package controller

import (
	"github.com/TrHung-297/chat-v2/dto"
	"github.com/TrHung-297/chat-v2/herror"
	"github.com/TrHung-297/chat-v2/infrastructure/controller"
	"github.com/TrHung-297/chat-v2/infrastructure/response"
	"github.com/TrHung-297/chat-v2/infrastructure/util"
	"github.com/TrHung-297/chat-v2/module/account/service"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AccountController struct {
	controller.BaseController
	AccountService     service.IAccountService
}

func NewAccountController(acc service.IAccountService, userWhiteList []string) *AccountController {
	return &AccountController{
		AccountService:     acc,
	}
}

func (controller *AccountController) LoginAccount(ctx *fiber.Ctx) error {
	start := time.Now()
	credentials := dto.Credentials{}
	if err := ctx.BodyParser(&credentials); err != nil {
		gerr := herror.New(herror.ErrorBindData, err, util.FuncName())
		message, resp := response.NewGerrorResponse(gerr, "")
		return controller.WriteUnauthorizedRequest(ctx, message, resp)
	}

	g_log.V(3).Infof("Request login for: %s - data: %v", credentials.Username, credentials)

	err := util.UsernameValidator(credentials.Username)
	if err != nil {
		message, resp := response.NewGerrorResponse(herror.New(herror.ErrorWrongUsername, err, util.FuncName()), "")
		return controller.WriteBadRequest(ctx, message, resp)
	}

	// Check Password
	ip := ctx.IP()
	_, accessToken, refreshToken, firstLogin, gerr := controller.AccountService.LogIn(credentials.Username, credentials.Password, credentials.ForceLogin, ip, credentials.ClientID)
	if gerr != nil {
		message, resp := response.NewGerrorResponse(gerr, "")
		return controller.WriteBadRequest(ctx, message, resp)
	}
	mainProcess := time.Since(start)

	// Generate Temp Username + Password for Turn Server
	//
	responseBody := map[string]interface{}{
		"AccessToken":  accessToken,
		"RefreshToken": refreshToken,
		"FirstLogin":   firstLogin,
	}
	processTime := time.Since(start)
	g_log.V(3).Infof("[LogIn] - Time log login : %v, all: %v, username: %s", mainProcess, processTime, credentials.Username)
	return controller.WriteSuccess(ctx, responseBody)
}