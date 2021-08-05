package handler

import (
	"learning-go-echo/application/usecase"
	"learning-go-echo/interface/context"
	"learning-go-echo/interface/validation"
	"learning-go-echo/utils/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArtistHanlder interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	Logout() echo.HandlerFunc
	Edit() echo.HandlerFunc
}

type artistHandler struct {
	au usecase.ArtistUsecase
}

func NewArtistHandler(au usecase.ArtistUsecase) ArtistHanlder {
	return &artistHandler{
		au: au,
	}
}

func (ah *artistHandler) Login() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.ArtistLoginRequest{}
		if ok, message := c.BindValidate(request, validation.ArtistLoginMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		token, err := ah.au.Login(request)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, err.Error())
		}

		c.SetCookieToken(token.AccessToken, token.RefreshToken, "artist")

		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "成功",
		})
	})
}

func (ah *artistHandler) Register() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.ArtistRegisterRequest{}
		if ok, message := c.BindValidate(request, validation.ArtistRegisterMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		token, err := ah.au.Register(request)
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		c.SetCookieToken(token.AccessToken, token.RefreshToken, "artist")

		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "成功",
		})
	})
}

func (ah *artistHandler) Logout() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		c.DeleteCookie(constant.ArtistAccessTokenName)
		c.DeleteCookie(constant.ArtistRefreshTokenName)

		return c.CustomResponse(http.StatusOK, map[string]string{
			"success": "ログアウト成功",
		})
	})
}

func (ah *artistHandler) Edit() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		request := &validation.ArtistEditRequest{}
		if ok, message := c.BindValidate(request, validation.ArtistEditMessage); !ok {
			return c.CustomResponse(http.StatusBadRequest, message)
		}

		artist, err := ah.au.Edit(request, c.GetAuthorID())
		if err != nil {
			return c.CustomResponse(http.StatusInternalServerError, err.Error())
		}

		return c.CustomResponse(http.StatusOK, map[string]interface{}{
			"artist": map[string]interface{}{
				"ID":              artist.ID,
				"name":            artist.Name,
				"Email":           artist.Email,
				"EmailVerifiedAt": artist.EmailVerifiedAt,
			},
		})
	})
}
