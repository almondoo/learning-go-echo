package handler

import (
	"learning-go-echo/application/usecase"
	"learning-go-echo/interface/context"
	"learning-go-echo/interface/validation"
	"learning-go-echo/utils/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Logout() echo.HandlerFunc
}

type adminHandler struct {
	au usecase.AdminUsecase
}

func NewAdminHandler(au usecase.AdminUsecase) AdminHandler {
	return &adminHandler{au: au}
}

func (ah *adminHandler) Login() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.AdminLoginRequest{}
		if ok, message := c.BindValidate(request, validation.AdminLoginMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		token, err := ah.au.Login(request)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]string{
				"error": "メールアドレスかパスワードが一致しません。",
			})
		}

		c.SetCookieToken(token.AccessToken, token.RefreshToken, "admin")

		return c.CustomResponse(http.StatusOK, map[string]string{
			"ログイン": "成功",
		})
	})
}

func (ah *adminHandler) Register() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.AdminRegisterRequest{}
		if ok, message := c.BindValidate(request, validation.AdminRegisterMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		token, err := ah.au.Register(request)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]string{
				"error": "管理者作成に失敗しました。",
			})
		}

		c.SetCookieToken(token.AccessToken, token.RefreshToken, "admin")

		return c.CustomResponse(http.StatusOK, "ユーザー作成")
	})
}

func (ah *adminHandler) Logout() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		if err := ah.au.Logout(c.GetCookieToken("admin")); err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]string{
				"error": "ログアウトに失敗しました。",
			})
		}

		c.DeleteCookie(constant.AdminAccessTokenName)
		c.DeleteCookie(constant.AdminRefreshTokenName)

		return c.CustomResponse(http.StatusOK, "logout")
	})
}
