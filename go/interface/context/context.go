package context

import (
	"errors"
	"learning-go-echo/infrastructure/auth"
	"learning-go-echo/interface/validation"
	"learning-go-echo/utils/constant"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
	Vwd validation.ValidatorWithDB
}

type callFunc func(c *CustomContext) error

func CastContext(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*CustomContext))
	}
}

var GrobalID uint

func (c *CustomContext) BindValidate(i interface{}, bindMessage func(err error) map[string]string) (bool, interface{}) {
	if err := c.Bind(i); err != nil {
		return false, err
	}
	if err := c.Validate(i); err != nil {
		errorMessage := bindMessage(err)
		return false, errorMessage
	}

	return true, nil
}

func (c *CustomContext) CustomResponse(code int, i interface{}) error {
	response := make(map[string]interface{})
	if code < 300 {
		response["status"] = "ok"
		response["data"] = i
	} else {
		response["status"] = "ng"
		response["data"] = map[string]interface{}{
			"error": i,
		}
	}

	return c.JSON(code, response)
}

//- Cookieを設置
func (c *CustomContext) SetCookieToken(accessToken, refreshToken, prefix string) {
	c.CreateCookie(prefix+constant.NotPrefixAccessTokenName, accessToken, time.Now().Add(auth.AccessExpires))
	c.CreateCookie(prefix+constant.NotPrefixRefreshTokenName, refreshToken, time.Now().Add(auth.RefreshExpires))
}

//- Cookieから取得
func (c *CustomContext) GetCookieToken(prefix string) *auth.TokenData {
	return &auth.TokenData{
		AccessToken:  c.GetCookieValue(prefix + constant.NotPrefixAccessTokenName),
		RefreshToken: c.GetCookieValue(prefix + constant.NotPrefixRefreshTokenName),
	}
}

func (c *CustomContext) CreateCookie(name, value string, expires time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expires
	cookie.Domain = os.Getenv("CORS_DOMAIN")
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func (c *CustomContext) GetCookieValue(name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (c *CustomContext) DeleteCookie(name string) error {
	cookie, err := c.Cookie(name)
	if err != nil {
		return errors.New("Cookieがありません。")
	}
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return nil
}

func (c *CustomContext) SetAuthorID(id uint) {
	GrobalID = id
}

func (c *CustomContext) GetAuthorID() uint {
	return GrobalID
}
