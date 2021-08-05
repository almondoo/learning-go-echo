package handler

import (
	"learning-go-echo/application/usecase"
	"learning-go-echo/interface/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	Get() echo.HandlerFunc
	Set() echo.HandlerFunc
}

type roleHandler struct {
	roleUsecase usecase.RoleUsecase
}

func NewRoleHandler(roleUsecase usecase.RoleUsecase) RoleHandler {
	return &roleHandler{roleUsecase: roleUsecase}
}

func (ah *roleHandler) Get() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		// var responseErr = make(map[string][]string)
		// loginReq := &validation.RoleLoginRequest{
		// 	Email:    c.FormValue("email"),
		// 	Password: c.FormValue("password"),
		// }
		// if err := validation.RoleLoginValidate(c, loginReq); err != nil {
		// 	return err
		// }

		// admin, err := ah.roleUsecase.Login(c.Request())
		// if err != nil {
		// 	responseErr["main"] = []string{"メールアドレスかパスワードが違います"}
		// 	return c.JSON(http.StatusBadRequest, responseErr)
		// }

		// ts, err := ah.token.CreateRoleToken(int(admin.ID))
		// if err != nil {
		// 	responseErr["token_error"] = []string{"トークンエラー"}
		// 	return c.JSON(http.StatusBadRequest, responseErr)
		// }
		// saveErr := ah.auth.CreateAuth(int(admin.ID), ts)
		// if saveErr != nil {
		// 	return c.JSON(http.StatusInternalServerError, saveErr.Error())
		// }

		// adminData := make(map[string]interface{})
		// adminData["access_token"] = ts.AccessToken
		// adminData["refresh_token"] = ts.RefreshToken
		// adminData["id"] = admin.ID
		// adminData["name"] = admin.Name
		// adminData["email"] = admin.Email

		// return response.Response(c, http.StatusOK, adminData)
		return c.CustomResponse(http.StatusOK, "get")
	})
}

func (ah *roleHandler) Set() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {

		return c.CustomResponse(http.StatusOK, "set")
	})
}
