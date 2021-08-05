package validation

import (
	"github.com/go-playground/validator/v10"
)

type ArtistCategoryRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func ArtistCategoryMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		//- タグ名
		var tag string
		//- エラーメッセージ
		var errorMessage string
		//- 変数名
		fieldName := err.Field()

		switch fieldName {
		case "Email":
			tag = err.Tag()
			switch tag {
			case "required":
				errorMessage = "必須項目です。"
			case "email":
				errorMessage = "メールアドレスの形式で入力してください。"
			}
		case "Password":
			tag = err.Tag()
			switch tag {
			case "required":
				errorMessage = "必須項目です。"
			}
		}

		errorMessages[fieldName] = errorMessage
	}

	return errorMessages
}
