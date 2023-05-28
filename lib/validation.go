package lib

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewVaditaion() *Validation {
	translator := en.New()
	UniversalTranslator := ut.New(translator, translator)

	trans, _ := UniversalTranslator.GetTranslator("en")

	validasi := validator.New()

	en_translation.RegisterDefaultTranslations(validasi, trans)
	
	// register tag label
	validasi.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})

	// membuat custom error
	validasi.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} harus diisi", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	return &Validation{validate: validasi, trans: trans}
}

func (v *Validation) Struct(s interface{}) interface{} {
	errors := make(map[string]string)

	err := v.validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.StructField()] = e.Translate(v.trans)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
