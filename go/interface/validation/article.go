package validation

import (
	"github.com/go-playground/validator/v10"
)

type CreateArticleRequest struct {
	Title     string `json:"title" form:"title" query:"title" validate:"required"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" query:"thumbnail" validate:"required,max=255"`
	Text      string `json:"text" form:"text" query:"text" validate:"required,max=20000"`
}

func (form *CreateArticleRequest) Validate() (bool, map[string]string) {
	var errorMessages = make(map[string]string)
	err := validator.New().Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Title":
					switch errors[i].Tag() {
					case "required":
						errorMessages["title"] = "必須項目です。"
					}

				case "Thumbnail":
					switch errors[i].Tag() {
					case "required":
						errorMessages["thumbnail"] = "必須項目です。"
					case "max":
						errorMessages["thumbnail"] = "異常な値が入りました。"
					}

				case "Text":
					switch errors[i].Tag() {
					case "required":
						errorMessages["text"] = "必須項目です。"
					case "max":
						errorMessages["text"] = "20000字以内で入力してください。"
					}
				}

			}
		}
		return false, errorMessages
	}

	return true, nil
}

func CreateArticleMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			switch errors[i].StructField() {
			case "Title":
				switch errors[i].Tag() {
				case "required":
					errorMessages["title"] = "必須項目です。"
				}

			case "Thumbnail":
				switch errors[i].Tag() {
				case "required":
					errorMessages["thumbnail"] = "必須項目です。"
				case "max":
					errorMessages["thumbnail"] = "異常な値が入りました。"
				}

			case "Text":
				switch errors[i].Tag() {
				case "required":
					errorMessages["text"] = "必須項目です。"
				case "max":
					errorMessages["text"] = "20000字以内で入力してください。"
				}
			}

		}
		return errorMessages
	}

	return nil
}

type EditArticleRequest struct {
	Id        uint   `json:"id" form:"id" query:"id" validate:"required,numeric"`
	Title     string `json:"title" form:"title" query:"title" validate:"required"`
	Thumbnail string `json:"thumbnail" form:"thumbnail" query:"thumbnail" validate:"required,max=255"`
	Text      string `json:"text" form:"text" query:"text" validate:"required,max=20000"`
}

func (form *EditArticleRequest) Validate() (bool, map[string]string) {
	var errorMessages = make(map[string]string)
	err := validator.New().Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Title":
					switch errors[i].Tag() {
					case "required":
						errorMessages["title"] = "必須項目です。"
					}

				case "Thumbnail":
					switch errors[i].Tag() {
					case "required":
						errorMessages["thumbnail"] = "必須項目です。"
					case "max":
						errorMessages["thumbnail"] = "異常な値が入りました。"
					}

				case "Text":
					switch errors[i].Tag() {
					case "required":
						errorMessages["text"] = "必須項目です。"
					case "max":
						errorMessages["text"] = "20000字以内で入力してください。"
					}
				}

			}
		}
		return false, errorMessages
	}

	return true, nil
}

func EditArticleMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			switch errors[i].StructField() {
			case "Title":
				switch errors[i].Tag() {
				case "required":
					errorMessages["title"] = "必須項目です。"
				}

			case "Thumbnail":
				switch errors[i].Tag() {
				case "required":
					errorMessages["thumbnail"] = "必須項目です。"
				case "max":
					errorMessages["thumbnail"] = "異常な値が入りました。"
				}

			case "Text":
				switch errors[i].Tag() {
				case "required":
					errorMessages["text"] = "必須項目です。"
				case "max":
					errorMessages["text"] = "20000字以内で入力してください。"
				}
			}

		}
		return errorMessages
	}

	return nil
}
