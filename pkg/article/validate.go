package article

import (
	"log"

	"github.com/aglide100/db-datasvc-example/pkg/model"
	"github.com/go-playground/validator"
)

func ArticleValidator(article Article) error {
	validate := validator.New()
	validate.RegisterValidation("script", model.ValidateScript)

	err := validate.Struct(article)
	if err != nil {
		log.Printf("Can't validate article, %v", err)
		return err
	}

	return nil
}