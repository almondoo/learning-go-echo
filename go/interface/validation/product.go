package validation

import (
	"github.com/go-playground/validator/v10"
)

type CreateProductRequest struct {
	ProductCategoryID uint     `json:"product_category_id" form:"product_category_id" query:"product_category_id" validate:"required"`
	Title             string   `json:"title" form:"title" query:"title" validate:"required,max=255"`
	Thumbnail         string   `json:"thumbnail" form:"thumbnail" query:"thumbnail" validate:"required"`
	Price             uint     `json:"price" form:"price" query:"price" validate:"required,min=1,max=999999999"`
	Description       string   `json:"description" form:"description" query:"description" validate:"required,max=2000"`
	VerticalLength    uint32   `json:"vertical_length" form:"vertical_length" query:"vertical_length" validate:"required,min=1,max=999999"`
	HorizontalLength  uint32   `json:"horizontal_length" form:"horizontal_length" query:"horizontal_length" validate:"required,min=1,max=999999"`
	IsPublished       uint8    `json:"is_published" form:"is_published" query:"is_published" validate:"min=0,max=1"`
	ProductImages     []string `json:"product_image" form:"product_image" query:"product_image"`
}

func (form *CreateProductRequest) Validate() (bool, map[string]string) {
	var errorMessages = make(map[string]string)
	err := validator.New().Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "ProductCategoryID":
					switch errors[i].Tag() {
					case "required":
						errorMessages["product_category_id"] = "必須項目です。"
					}

				case "Title":
					switch errors[i].Tag() {
					case "required":
						errorMessages["title"] = "必須項目です。"
					}

				case "Thumbnail":
					switch errors[i].Tag() {
					case "required":
						errorMessages["thumbnail"] = "必須項目です。"
					}

				case "Price":
					switch errors[i].Tag() {
					case "required":
						errorMessages["price"] = "必須項目です。"
					case "min":
						errorMessages["price"] = "1円以上で入力してください。"
					case "max":
						errorMessages["price"] = "999999999円以内で入力してください。"
					}

				case "Description":
					switch errors[i].Tag() {
					case "required":
						errorMessages["description"] = "必須項目です。"
					case "max":
						errorMessages["description"] = "2000文字以内で入力してください。"
					}

				case "VerticalLength":
					switch errors[i].Tag() {
					case "required":
						errorMessages["vertical_length"] = "必須項目です。"
					case "min":
						errorMessages["vertical_length"] = "999999以下で入力してください。"
					case "max":
						errorMessages["vertical_length"] = "1~999999以下で入力してください。"
					}

				case "HorizontalLength":
					switch errors[i].Tag() {
					case "required":
						errorMessages["horizontal_length"] = "必須項目です。"
					case "min":
						errorMessages["horizontal_length"] = "999999以下で入力してください。"
					case "max":
						errorMessages["horizontal_length"] = "1~999999以下で入力してください。"
					}

				case "IsPublished":
					switch errors[i].Tag() {
					case "min":
						errorMessages["is_published"] = "不正な値です。"
					case "max":
						errorMessages["is_published"] = "不正な値です。"
					}

				case "ProductImages":
					switch errors[i].Tag() {
					case "min":
						errorMessages["product_image"] = "不正な値です。"
					case "max":
						errorMessages["product_image"] = "不正な値です。"
					}
				}

			}
		}
		return false, errorMessages
	}

	return true, nil
}

func CreateProductMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			switch errors[i].StructField() {
			case "ProductCategoryID":
				switch errors[i].Tag() {
				case "required":
					errorMessages["product_category_id"] = "必須項目です。"
				}

			case "Title":
				switch errors[i].Tag() {
				case "required":
					errorMessages["title"] = "必須項目です。"
				}

			case "Thumbnail":
				switch errors[i].Tag() {
				case "required":
					errorMessages["thumbnail"] = "必須項目です。"
				}

			case "Price":
				switch errors[i].Tag() {
				case "required":
					errorMessages["price"] = "必須項目です。"
				case "min":
					errorMessages["price"] = "1円以上で入力してください。"
				case "max":
					errorMessages["price"] = "999999999円以内で入力してください。"
				}

			case "Description":
				switch errors[i].Tag() {
				case "required":
					errorMessages["description"] = "必須項目です。"
				case "max":
					errorMessages["description"] = "2000文字以内で入力してください。"
				}

			case "VerticalLength":
				switch errors[i].Tag() {
				case "required":
					errorMessages["vertical_length"] = "必須項目です。"
				case "min":
					errorMessages["vertical_length"] = "999999以下で入力してください。"
				case "max":
					errorMessages["vertical_length"] = "1~999999以下で入力してください。"
				}

			case "HorizontalLength":
				switch errors[i].Tag() {
				case "required":
					errorMessages["horizontal_length"] = "必須項目です。"
				case "min":
					errorMessages["horizontal_length"] = "999999以下で入力してください。"
				case "max":
					errorMessages["horizontal_length"] = "1~999999以下で入力してください。"
				}

			case "IsPublished":
				switch errors[i].Tag() {
				case "min":
					errorMessages["is_published"] = "不正な値です。"
				case "max":
					errorMessages["is_published"] = "不正な値です。"
				}

			case "ProductImages":
				switch errors[i].Tag() {
				case "min":
					errorMessages["product_image"] = "不正な値です。"
				case "max":
					errorMessages["product_image"] = "不正な値です。"
				}
			}

		}
		return errorMessages
	}

	return nil
}

type EditProductRequest struct {
	ProductCategoryID uint     `json:"product_category_id" form:"product_category_id" query:"product_category_id" validate:"required"`
	Title             string   `json:"title" form:"title" query:"title" validate:"required,max=255"`
	Thumbnail         string   `json:"thumbnail" form:"thumbnail" query:"thumbnail" validate:"required"`
	Price             uint     `json:"price" form:"price" query:"price" validate:"required,min=1,max=999999999"`
	Description       string   `json:"description" form:"description" query:"description" validate:"required,max=2000"`
	VerticalLength    uint32   `json:"vertical_length" form:"vertical_length" query:"vertical_length" validate:"required,min=1,max=999999"`
	HorizontalLength  uint32   `json:"horizontal_length" form:"horizontal_length" query:"horizontal_length" validate:"required,min=1,max=999999"`
	IsPublished       uint8    `json:"is_published" form:"is_published" query:"is_published" validate:"min=0,max=1"`
	ProductImages     []string `json:"product_image" form:"product_image" query:"product_image"`
}

func (form *EditProductRequest) Validate() (bool, map[string]string) {
	var errorMessages = make(map[string]string)
	err := validator.New().Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "ProductCategoryID":
					switch errors[i].Tag() {
					case "required":
						errorMessages["product_category_id"] = "必須項目です。"
					}

				case "Title":
					switch errors[i].Tag() {
					case "required":
						errorMessages["title"] = "必須項目です。"
					}

				case "Thumbnail":
					switch errors[i].Tag() {
					case "required":
						errorMessages["thumbnail"] = "必須項目です。"
					}

				case "Price":
					switch errors[i].Tag() {
					case "required":
						errorMessages["price"] = "必須項目です。"
					case "min":
						errorMessages["price"] = "1円以上で入力してください。"
					case "max":
						errorMessages["price"] = "999999999円以内で入力してください。"
					}

				case "Description":
					switch errors[i].Tag() {
					case "required":
						errorMessages["description"] = "必須項目です。"
					case "max":
						errorMessages["description"] = "2000文字以内で入力してください。"
					}

				case "VerticalLength":
					switch errors[i].Tag() {
					case "required":
						errorMessages["vertical_length"] = "必須項目です。"
					case "min":
						errorMessages["vertical_length"] = "999999以下で入力してください。"
					case "max":
						errorMessages["vertical_length"] = "1~999999以下で入力してください。"
					}

				case "HorizontalLength":
					switch errors[i].Tag() {
					case "required":
						errorMessages["horizontal_length"] = "必須項目です。"
					case "min":
						errorMessages["horizontal_length"] = "999999以下で入力してください。"
					case "max":
						errorMessages["horizontal_length"] = "1~999999以下で入力してください。"
					}

				case "IsPublished":
					switch errors[i].Tag() {
					case "min":
						errorMessages["is_published"] = "不正な値です。"
					case "max":
						errorMessages["is_published"] = "不正な値です。"
					}

				case "ProductImages":
					switch errors[i].Tag() {
					case "min":
						errorMessages["product_image"] = "不正な値です。"
					case "max":
						errorMessages["product_image"] = "不正な値です。"
					}
				}

			}
		}
		return false, errorMessages
	}

	return true, nil
}

func EditProductMessage(err error) map[string]string {
	var errorMessages = make(map[string]string)
	errors := err.(validator.ValidationErrors)
	if len(errors) != 0 {
		for i := range errors {
			switch errors[i].StructField() {
			case "ProductCategoryID":
				switch errors[i].Tag() {
				case "required":
					errorMessages["product_category_id"] = "必須項目です。"
				}

			case "Title":
				switch errors[i].Tag() {
				case "required":
					errorMessages["title"] = "必須項目です。"
				}

			case "Thumbnail":
				switch errors[i].Tag() {
				case "required":
					errorMessages["thumbnail"] = "必須項目です。"
				}

			case "Price":
				switch errors[i].Tag() {
				case "required":
					errorMessages["price"] = "必須項目です。"
				case "min":
					errorMessages["price"] = "1円以上で入力してください。"
				case "max":
					errorMessages["price"] = "999999999円以内で入力してください。"
				}

			case "Description":
				switch errors[i].Tag() {
				case "required":
					errorMessages["description"] = "必須項目です。"
				case "max":
					errorMessages["description"] = "2000文字以内で入力してください。"
				}

			case "VerticalLength":
				switch errors[i].Tag() {
				case "required":
					errorMessages["vertical_length"] = "必須項目です。"
				case "min":
					errorMessages["vertical_length"] = "999999以下で入力してください。"
				case "max":
					errorMessages["vertical_length"] = "1~999999以下で入力してください。"
				}

			case "HorizontalLength":
				switch errors[i].Tag() {
				case "required":
					errorMessages["horizontal_length"] = "必須項目です。"
				case "min":
					errorMessages["horizontal_length"] = "999999以下で入力してください。"
				case "max":
					errorMessages["horizontal_length"] = "1~999999以下で入力してください。"
				}

			case "IsPublished":
				switch errors[i].Tag() {
				case "min":
					errorMessages["is_published"] = "不正な値です。"
				case "max":
					errorMessages["is_published"] = "不正な値です。"
				}

			case "ProductImages":
				switch errors[i].Tag() {
				case "min":
					errorMessages["product_image"] = "不正な値です。"
				case "max":
					errorMessages["product_image"] = "不正な値です。"
				}
			}

		}
		return errorMessages
	}

	return nil
}
