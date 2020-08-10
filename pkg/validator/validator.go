// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package validator

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translate "github.com/go-playground/validator/v10/translations/zh"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	c.lazyInit()

	c.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})
	_ = c.validate.RegisterValidation("mobile", mobile)

	c.validate.RegisterTagNameFunc(func(filed reflect.StructField) string {
		name := filed.Tag.Get("label")
		return name
	})

	// i18n
	e := en.New()
	uniTrans := ut.New(e, e, zh.New(), zh_Hant_TW.New())
	translator, _ := uniTrans.GetTranslator("zh")
	zh_translate.RegisterDefaultTranslations(c.validate, translator)

	c.validate.RegisterTranslation("mobile", translator, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0}长度不等于11位或{1}格式错误", true)
	}, func(ut ut.Translator, ve validator.FieldError) string {
		t, _ := ut.T("mobile", ve.Field(), ve.Field())
		return t
	})

	var sb strings.Builder

	err := c.validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			sb.WriteString(err.Translate(translator))
			sb.WriteString(" ")
		}

		return errors.New(sb.String())
	}
	return nil
}

func (c *CustomValidator) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}

func mobile(vf validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, vf.Field().String())
	return ok
}
