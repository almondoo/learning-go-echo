package handler

import (
	"learning-go-echo/interface/context"
	"learning-go-echo/interface/fileupload"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CropHandler interface {
	Crop() echo.HandlerFunc
}

type cropHandler struct {
	fu fileupload.FileUpload
}

func NewCropHandler(fu fileupload.FileUpload) CropHandler {
	return &cropHandler{fu: fu}
}

func (ch *cropHandler) Crop() echo.HandlerFunc {
	return context.CastContext(func(c *context.CustomContext) error {
		x, err := strconv.Atoi(c.FormValue("x"))
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, "リクエストされた値が不正です。")
		}
		y, err := strconv.Atoi(c.FormValue("y"))
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, "リクエストされた値が不正です。")
		}
		width, err := strconv.Atoi(c.FormValue("width"))
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, "リクエストされた値が不正です。")
		}
		height, err := strconv.Atoi(c.FormValue("height"))
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, "リクエストされた値が不正です。")
		}

		cropData := map[string]int{
			"x":      x,
			"y":      y,
			"width":  width,
			"height": height,
		}

		file, err := c.FormFile("file")
		if err != nil {
			return c.CustomResponse(http.StatusBadRequest, "リクエストされた値が不正です。")
		}
		if err = ch.fu.Crop(file, cropData); err != nil {
			return c.CustomResponse(http.StatusInternalServerError, "クロップに失敗しました")
		}

		return c.CustomResponse(http.StatusOK, "成功しました。")
	})
}
