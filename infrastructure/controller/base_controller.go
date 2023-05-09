package controller

import (
	"encoding/json"
	"github.com/TrHung-297/chat-v2/infrastructure/response"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(b)
}

// BaseController define a controller object
type BaseController struct {
}

// WriteSuccess writes Success response
func (controller *BaseController) WriteSuccess(ctx *fiber.Ctx, v interface{}) error {
	res := response.Response{
		Message: "Success",
		Data:    v,
	}

	return ctx.JSON(res)
}

// WriteSuccessEmptyContent writes Success response with empty content
func (controller *BaseController) WriteSuccessEmptyContent(ctx *fiber.Ctx) error {
	res := response.Response{
		Message: "Success",
		Data:    nil,
	}

	// Log response
	// g_log.V(3).Infof(util.ToJSON(response))

	// Return
	return ctx.JSON(res)
}

// WriteBadRequest writes BadRequest response (client-side herror)
func (controller *BaseController) WriteBadRequest(ctx *fiber.Ctx, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(ctx, http.StatusBadRequest, message, errorRes)
}

// WriteUnauthorizedRequest writes UnauthorizedRequest response (client-side herror)
func (controller *BaseController) WriteUnauthorizedRequest(ctx *fiber.Ctx, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(ctx, http.StatusUnauthorized, message, errorRes)
}

// WriteNotFoundRequest writes NotFoundRequest response (client-side response)
func (controller *BaseController) WriteNotFoundRequest(ctx *fiber.Ctx, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(ctx, http.StatusNotFound, message, errorRes)
}

// WriteConflictRequest writes ConflictRequest response (client-side herror)
func (controller *BaseController) WriteConflictRequest(ctx *fiber.Ctx, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(ctx, http.StatusConflict, message, errorRes)
}

// WriteInternalServerError writes InternalServerError response (server-side herror)
func (controller *BaseController) WriteInternalServerError(ctx *fiber.Ctx, message string, errorRes response.ErrorResponse) error {
	return controller.writeError(ctx, http.StatusInternalServerError, message, errorRes)
}

// IsValid validates body object
func (controller *BaseController) IsValid(m interface{}) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// writeError writes an herror response
func (controller *BaseController) writeError(ctx *fiber.Ctx, statusCode int, message string, err response.ErrorResponse) error {
	res := response.Response{
		Message: message,
		Data:    err,
	}

	// Log herror
	g_log.V(5).Errorf("[API] - Request: %s, Response: %v", ctx.Request().RequestURI(), ToJSON(res))

	// Return
	return ctx.Status(statusCode).JSON(res)
}