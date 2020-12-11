/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/10/26
   Description :
-------------------------------------------------
*/

package zutils

import (
	"errors"
	"regexp"

	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var Validate = newVerify()

type validateUtil struct {
	validateTrans ut.Translator
	validate      *validator.Validate
}

func newVerify() *validateUtil {
	zh := zhongwen.New()
	validateTrans, _ := ut.New(zh, zh).GetTranslator("zh")

	v := &validateUtil{
		validateTrans: validateTrans,
		validate:      validator.New(),
	}
	_ = zh_translations.RegisterDefaultTranslations(v.validate, validateTrans)

	_ = v.validate.RegisterValidation("regex", v.validateRegex)
	_ = v.validate.RegisterValidation("time", v.validateTime)
	return v
}

// 正则匹配
func (*validateUtil) validateRegex(f validator.FieldLevel) bool {
	compile := f.Param()
	text := f.Field().String()
	return regexp.MustCompile(compile).MatchString(text)
}

// 时间匹配
func (*validateUtil) validateTime(f validator.FieldLevel) bool {
	layout := f.Param()
	if layout == "" {
		layout = Time.Layout
	}
	text := f.Field().String()

	_, err := Time.TextToTimeOfLayout(text, layout)
	return err == nil
}

// 将错误翻译为中文
func (u *validateUtil) translateValidateErr(err error) error {
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		for _, e := range errs {
			return errors.New(e.Translate(u.validateTrans))
		}
	}
	return nil
}

// 校验结构体
func (u *validateUtil) ValidateStruct(a interface{}) error {
	err := u.validate.Struct(a)
	return u.translateValidateErr(err)
}

// 验证字段
func (u *validateUtil) ValidateField(a interface{}, tag string) error {
	err := u.validate.Var(a, tag)
	return u.translateValidateErr(err)
}
